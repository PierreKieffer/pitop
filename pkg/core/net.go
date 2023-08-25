package core

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/PierreKieffer/pitop/pkg/utils"
)

type NetworkStatus struct {
	BytesRecv      uint64
	BytesSent      uint64
	TotalBytesRecv uint64
	TotalBytesSent uint64
}

func (networkStatus *NetworkStatus) ComputeNetworkStatus() {
	prevStat := GetNetworkStatus()
	time.Sleep(time.Second)
	currentStat := GetNetworkStatus()

	if prevStat.TotalBytesRecv == 0 && prevStat.TotalBytesSent == 0 {
		currentStat.BytesRecv = 0
		currentStat.BytesSent = 0

		*networkStatus = *currentStat
		return
	}

	currentStat.BytesRecv = currentStat.TotalBytesRecv - prevStat.TotalBytesRecv
	currentStat.BytesSent = currentStat.TotalBytesSent - prevStat.TotalBytesSent
	*networkStatus = *currentStat
	return
}

func GetNetworkStatus() *NetworkStatus {

	var netStats []NetworkStatus

	netStatBytes, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		panic(err)
	}

	dataSlice := strings.Split(string(netStatBytes), "\n")

	for _, statLine := range dataSlice[2:] {
		statSlice := utils.FormatStatSlice(strings.Split(statLine, " "))
		if len(statSlice) > 1 && statSlice[0] != "" {
			var networkStatus NetworkStatus
			networkStatus.TotalBytesRecv, _ = strconv.ParseUint(statSlice[1], 10, 64)
			networkStatus.TotalBytesSent, _ = strconv.ParseUint(statSlice[9], 10, 64)

			netStats = append(netStats, networkStatus)
		}
	}

	var totalNetStat NetworkStatus

	for _, networkStatus := range netStats {
		totalNetStat.TotalBytesRecv += networkStatus.TotalBytesRecv
		totalNetStat.TotalBytesSent += networkStatus.TotalBytesSent

	}

	return &totalNetStat
}
