package worker

import (
	"github.com/PierreKieffer/pitop/pkg/cpu"
	"github.com/PierreKieffer/pitop/pkg/disk"
	"github.com/PierreKieffer/pitop/pkg/mem"
	"github.com/PierreKieffer/pitop/pkg/temp"
	"sync"
)

type Status struct {
	CPULoad *cpu.CPULoad
	CPUFreq *cpu.CPUFreq
	Mem     *mem.MemStat
	Temp    *temp.Temp
	Disk    *disk.DiskInfo
}

func Worker() *Status {

	var status Status

	var wg sync.WaitGroup
	wg.Add(5)

	//CPU Load
	go func() {
		defer wg.Done()
		status.CPULoad = cpu.ComputeCPULoad()
	}()

	// CPU Frequency
	go func() {
		defer wg.Done()
		status.CPUFreq = cpu.ExtractCPUFrequency()
	}()

	//Mem Stats
	go func() {
		defer wg.Done()
		status.Mem = mem.GetMemStats()
	}()

	// Temp
	go func() {
		defer wg.Done()
		status.Temp = temp.ExtractTemp()
	}()

	// Disk
	go func() {
		defer wg.Done()
		status.Disk = disk.ExtractDiskUsage()
	}()
	wg.Wait()

	return &status
}
