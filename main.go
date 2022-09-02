package main

import (
	"fmt"

	"github.com/PierreKieffer/pitop/pkg/core"
)

func main() {
	// ui.App()

	var status = &core.Status{
		CPU:         &core.CPU{},
		Memory:      &core.Memory{},
		Temperature: &core.Temperature{},
		Disk:        &core.Disk{},
		Network:     &core.NetworkStatus{},
	}
	status.Worker()

	fmt.Println("CPU : ")
	fmt.Println(status.CPU)
	fmt.Println("Memory : ")
	fmt.Println(status.Memory)
	fmt.Println("Temperature : ")
	fmt.Println(status.Temperature)
	fmt.Println("Disk : ")
	fmt.Println(status.Disk)
	fmt.Println("Network : ")
	fmt.Println(status.Network)
}
