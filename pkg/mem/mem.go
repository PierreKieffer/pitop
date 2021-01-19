package mem

import (
	"fmt"
	"github.com/PierreKieffer/pitop/pkg/utils"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type MemStat struct {
	MemTotal  uint64
	MemFree   uint64
	MemUsage  float32
	SwapTotal uint64
	SwapFree  uint64
	SwapUsage float32

	Buffers uint64
	Cached  uint64
}

type MemUsage struct {
}

func GetMemStats() *MemStat {

	var memStat MemStat

	memStatBytes, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(memStatBytes), "\n")

	for _, statLine := range dataSlice {
		statSlice := utils.FormatStatSlice(strings.Split(statLine, " "))
		if len(statSlice) > 2 {
			if statSlice[0] != "" {
				switch statSlice[0][:len(statSlice[0])-1] {
				case "MemTotal":
					memStat.MemTotal, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "MemFree":
					memStat.MemFree, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "Buffers":
					memStat.Buffers, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "Cached":
					memStat.Cached, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "SwapTotal":
					memStat.SwapTotal, _ = strconv.ParseUint(statSlice[1], 10, 64)
				case "SwapFree":
					memStat.SwapFree, _ = strconv.ParseUint(statSlice[1], 10, 64)
				}
			}
		}
	}

	memStat.MemUsage = (float32(memStat.MemTotal) - float32(memStat.MemFree)) / float32(memStat.MemTotal) * 100
	memStat.SwapUsage = (float32(memStat.SwapTotal) - float32(memStat.SwapFree)) / float32(memStat.SwapTotal) * 100

	return &memStat
}
