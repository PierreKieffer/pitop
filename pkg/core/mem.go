package core

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/PierreKieffer/pitop/pkg/utils"
)

type Memory struct {
	memoryInfoMu sync.Mutex

	MemTotal  uint64
	MemFree   uint64
	MemUsage  float32
	SwapTotal uint64
	SwapFree  uint64
	SwapUsage float32

	Buffers uint64
	Cached  uint64
}

func (memoryInfo *Memory) Usage() {

	memoryInfo.memoryInfoMu.Lock()
	defer memoryInfo.memoryInfoMu.Unlock()

	memInfoBytes, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(memInfoBytes), "\n")

	for _, statLine := range dataSlice {
		statSlice := utils.FormatStatSlice(strings.Split(statLine, " "))
		if len(statSlice) > 2 {
			if statSlice[0] != "" {
				switch statSlice[0][:len(statSlice[0])-1] {
				case "MemTotal":
					memoryInfo.MemTotal, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "MemFree":
					memoryInfo.MemFree, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "Buffers":
					memoryInfo.Buffers, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "Cached":
					memoryInfo.Cached, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "SwapTotal":
					memoryInfo.SwapTotal, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "SwapFree":
					memoryInfo.SwapFree, _ = strconv.ParseUint(statSlice[1], 10, 64)
				}
			}
		}
	}

	memoryInfo.MemUsage = (float32(memoryInfo.MemTotal) - float32(memoryInfo.MemFree)) / float32(memoryInfo.MemTotal) * 100
	memoryInfo.SwapUsage = (float32(memoryInfo.SwapTotal) - float32(memoryInfo.SwapFree)) / float32(memoryInfo.SwapTotal) * 100
}
