package net

import (
	"fmt"
	"github.com/PierreKieffer/pitop/pkg/utils"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type NetStat struct {
	BytesRecv      uint64
	BytesSent      uint64
	TotalBytesRecv uint64
	TotalBytesSent uint64
}

func GetNetStats() *NetStat {

	var netStats []NetStat

	netStatBytes, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataSlice := strings.Split(string(netStatBytes), "\n")

	for _, statLine := range dataSlice[2:] {
		statSlice := utils.FormatStatSlice(strings.Split(statLine, " "))
		ExtractNetStats(&netStats, statSlice)
	}

	var totalNetStat NetStat

	for _, netStat := range netStats {
		totalNetStat.TotalBytesRecv += netStat.TotalBytesRecv
		totalNetStat.TotalBytesSent += netStat.TotalBytesSent

	}

	return &totalNetStat
}

func ExtractNetStats(netStats *[]NetStat, statSlice []string) {
	if len(statSlice) > 1 && statSlice[0] != "" {
		var netStat NetStat
		netStat.TotalBytesRecv, _ = strconv.ParseUint(statSlice[1], 10, 64)
		netStat.TotalBytesSent, _ = strconv.ParseUint(statSlice[9], 10, 64)

		*netStats = append(*netStats, netStat)
	}
}

func ComputeNetStats() *NetStat {
	prevNetStats := GetNetStats()
	time.Sleep(time.Second)
	netStats := GetNetStats()

	if prevNetStats.TotalBytesRecv == 0 && prevNetStats.TotalBytesSent == 0 {
		netStats.BytesRecv = 0
		netStats.BytesSent = 0

		return netStats
	}

	netStats.BytesRecv = netStats.TotalBytesRecv - prevNetStats.TotalBytesRecv
	netStats.BytesSent = netStats.TotalBytesSent - prevNetStats.TotalBytesSent

	return netStats
}
