package core

import (
	"sync"
)

type Status struct {
	statusMu sync.Mutex

	CPU         *CPU
	Memory      *Memory
	Temperature *Temperature
	Disk        *Disk
	Network     *NetworkStatus
}

func (status *Status) Worker() {

	var wg sync.WaitGroup
	wg.Add(6)

	status.statusMu.Lock()
	defer status.statusMu.Unlock()

	//CPU Load
	go func() {
		defer wg.Done()
		status.CPU.Load()
	}()

	// CPU Frequency
	go func() {
		defer wg.Done()
		status.CPU.Frequency()
	}()

	//Mem Stats
	go func() {
		defer wg.Done()
		status.Memory.Usage()
	}()

	// Temp
	go func() {
		defer wg.Done()
		status.Temperature.ExtractTemp()
	}()

	// Disk
	go func() {
		defer wg.Done()
		status.Disk.ExtractDiskUsage()
	}()

	// Net
	go func() {
		defer wg.Done()
		status.Network.ComputeNetworkStatus()
	}()
	wg.Wait()
}
