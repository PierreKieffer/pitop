package core

import (
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PierreKieffer/pitop/pkg/utils"
)

type CPU struct {
	/*
		CPU Object
	*/
	cpuMu sync.Mutex

	Freq int // CPU total average frequency

	// CPU Load on each core
	CPU0 float32
	CPU1 float32
	CPU2 float32
	CPU3 float32
}

type CPUInfo struct {
	/*
		Complete description of the CPU (all cores) at a specific time
	*/

	cpu0 *CoreInfo
	cpu1 *CoreInfo
	cpu2 *CoreInfo
	cpu3 *CoreInfo
}

type CoreInfo struct {
	/*
		Core description at a specific time
	*/

	CoreId    string
	User      uint64
	Nice      uint64
	System    uint64
	Idle      uint64
	IOWait    uint64
	IRQ       uint64
	SoftIRQ   uint64
	Steal     uint64
	Guest     uint64
	GuestNice uint64
}

func (cpu *CPU) Load() {

	// Extract stats
	prevExtract := ExtractCPUInfo() // Return a CPUInfo (with CoreInfo)
	time.Sleep(time.Second)
	extract := ExtractCPUInfo() // Return a CPUInfo (with CoreInfo)

	// NOTE TODO : Here could be interesting to use a sync.Pool for CoreInfo struct ?
	// As it is used only to add data to *CPU

	var wg sync.WaitGroup
	wg.Add(4)

	cpu.cpuMu.Lock()
	defer cpu.cpuMu.Unlock()

	// cpu0
	go func() {
		defer wg.Done()
		cpu.CPU0 = ComputeCoreLoad(extract.cpu0, prevExtract.cpu0) // return float32
	}()

	// cpu1
	go func() {
		defer wg.Done()
		cpu.CPU1 = ComputeCoreLoad(extract.cpu1, prevExtract.cpu1)
	}()

	// cpu2
	go func() {
		defer wg.Done()
		cpu.CPU2 = ComputeCoreLoad(extract.cpu2, prevExtract.cpu2)
	}()

	// cpu3
	go func() {
		defer wg.Done()
		cpu.CPU3 = ComputeCoreLoad(extract.cpu3, prevExtract.cpu3)
	}()

	wg.Wait()

}

func (cpu *CPU) Frequency() {
	/*
		Compute average frequency in MHz
	*/

	cpuInfoBytes, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		panic(err)
	}

	cpuFrequencies := []string{} // TODO : Check the len, to init with prealocated cap
	dataSlice := strings.Split(string(cpuInfoBytes), "\n")
	for i := range dataSlice {
		cpuData := strings.Split(dataSlice[i], " ")

		if len(cpuData) > 1 && cpuData[0] != "" && cpuData[0] == "cpu" && cpuData[1][:3] == "MHz" {
			cpuFrequencies = append(cpuFrequencies, cpuData[2])
		}
	}

	var freqSum int = 0

	for i := range cpuFrequencies {
		freq, _ := strconv.ParseFloat(cpuFrequencies[i], 32)
		freqSum += int(freq)
	}
	avgFreq := freqSum / len(cpuFrequencies)

	cpu.cpuMu.Lock()
	defer cpu.cpuMu.Unlock()

	cpu.Freq = avgFreq
}

func ExtractCPUInfo() *CPUInfo {
	/*
		Method to parse /proc/stat file and extract each stats for each core
	*/

	var cpuInfo CPUInfo

	procStatBytes, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		panic(err)
	}

	dataSlice := strings.Split(string(procStatBytes), "\n")

	for _, statLine := range dataSlice {
		statSlice := utils.FormatStatSlice(strings.Split(statLine, " "))

		if len(statSlice) > 0 {
			if statSlice[0] != "" && statSlice[0][:3] == "cpu" {
				var coreInfo CoreInfo
				coreInfo.CoreId = statSlice[0]
				coreInfo.User, _ = strconv.ParseUint(statSlice[1], 10, 64)
				coreInfo.Nice, _ = strconv.ParseUint(statSlice[2], 10, 64)
				coreInfo.System, _ = strconv.ParseUint(statSlice[3], 10, 64)
				coreInfo.Idle, _ = strconv.ParseUint(statSlice[4], 10, 64)
				coreInfo.IOWait, _ = strconv.ParseUint(statSlice[5], 10, 64)
				coreInfo.IRQ, _ = strconv.ParseUint(statSlice[6], 10, 64)
				coreInfo.SoftIRQ, _ = strconv.ParseUint(statSlice[7], 10, 64)
				coreInfo.Steal, _ = strconv.ParseUint(statSlice[8], 10, 64)
				coreInfo.Guest, _ = strconv.ParseUint(statSlice[9], 10, 64)
				coreInfo.GuestNice, _ = strconv.ParseUint(statSlice[10], 10, 64)

				switch statSlice[0] {
				case "cpu0":
					cpuInfo.cpu0 = &coreInfo
				case "cpu1":
					cpuInfo.cpu1 = &coreInfo
				case "cpu2":
					cpuInfo.cpu2 = &coreInfo
				case "cpu3":
					cpuInfo.cpu3 = &coreInfo
				}

			}
		}
	}

	return &cpuInfo
}

func ComputeCoreLoad(currentStat, previousStat *CoreInfo) float32 {

	/*
	   user    nice   system  idle      iowait irq   softirq  steal  guest  guest_nice
	   cpu  74608   2520   24433   1117073   6176   4054  0        0      0      0
	   *
	   *    Idle = idle + iowait
	   *    Load = user + nice + system + irq + softirq + steal
	   *    Total = Idle + Load
	   *
	   *    DiffTotal = Total_t - Total_t-1
	   *    DiffIdle = Idle_t - Idle_t-1
	   *    percentage = ( DiffTotal - DiffIdle ) / DiffTotal
	   *
	   *
	*/

	if currentStat != nil && previousStat != nil {
		PreviousIdle := previousStat.Idle + previousStat.IOWait
		PreviousLoad := previousStat.User + previousStat.Nice + previousStat.System + previousStat.IRQ + previousStat.SoftIRQ + previousStat.Steal
		PreviousTotal := PreviousIdle + PreviousLoad

		Idle := currentStat.Idle + currentStat.IOWait
		Load := currentStat.User + currentStat.Nice + currentStat.System + currentStat.IRQ + currentStat.SoftIRQ + currentStat.Steal
		Total := Idle + Load

		DiffTotal := Total - PreviousTotal
		DiffIdle := Idle - PreviousIdle

		CPULoadPercentage := (float32(DiffTotal) - float32(DiffIdle)) / float32(DiffTotal) * 100

		return CPULoadPercentage

	} else {
		return 0
	}
}
