package ui

import (
	"time"

	"fmt"

	"github.com/PierreKieffer/pitop/pkg/core"
	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func App() {

	if err := termui.Init(); err != nil {
		panic(err)
	}
	defer termui.Close()

	//Init
	println("\n    Loading ...    ")
	var status = &core.Status{
		CPU:         &core.CPU{},
		Memory:      &core.Memory{},
		Temperature: &core.Temperature{},
		Disk:        &core.Disk{},
		Network:     &core.NetworkStatus{},
	}
	status.Worker()

	//CPU Load
	g1 := widgets.NewGauge()
	g1.Title = " CPU0 "
	g1.Percent = int(status.CPU.CPU0)
	g1.SetRect(0, 0, 50, 3)
	g1.BarColor = GetColorPercent(status.CPU.CPU0)
	g1.BorderStyle.Fg = termui.ColorWhite
	g1.LabelStyle = termui.NewStyle(termui.ColorWhite)

	g2 := widgets.NewGauge()
	g2.Title = " CPU1 "
	g2.Percent = int(status.CPU.CPU1)
	g2.SetRect(0, 3, 50, 6)
	g2.BarColor = GetColorPercent(status.CPU.CPU1)
	g2.BorderStyle.Fg = termui.ColorWhite
	g2.LabelStyle = termui.NewStyle(termui.ColorWhite)

	g3 := widgets.NewGauge()
	g3.Title = " CPU2 "
	g3.Percent = int(status.CPU.CPU2)
	g3.SetRect(0, 6, 50, 9)
	g3.BarColor = GetColorPercent(status.CPU.CPU2)
	g3.BorderStyle.Fg = termui.ColorWhite
	g3.LabelStyle = termui.NewStyle(termui.ColorWhite)

	g4 := widgets.NewGauge()
	g4.Title = " CPU3 "
	g4.Percent = int(status.CPU.CPU3)
	g4.SetRect(0, 9, 50, 12)
	g4.BarColor = GetColorPercent(status.CPU.CPU3)
	g4.BorderStyle.Fg = termui.ColorWhite
	g4.LabelStyle = termui.NewStyle(termui.ColorWhite)

	// Memory
	gMemUsage := widgets.NewGauge()
	gMemUsage.Title = " Memory usage "
	gMemUsage.Percent = int(status.Memory.MemUsage)
	gMemUsage.SetRect(50, 0, 100, 3)
	gMemUsage.BarColor = GetColorPercent(status.Memory.MemUsage)
	gMemUsage.BorderStyle.Fg = termui.ColorWhite
	gMemUsage.LabelStyle = termui.NewStyle(termui.ColorWhite)

	gSwapUsage := widgets.NewGauge()
	gSwapUsage.Title = " Swap usage "
	gSwapUsage.Percent = int(status.Memory.SwapUsage)
	gSwapUsage.SetRect(50, 3, 100, 6)
	gSwapUsage.BarColor = GetColorPercent(status.Memory.SwapUsage)
	gSwapUsage.BorderStyle.Fg = termui.ColorWhite
	gSwapUsage.LabelStyle = termui.NewStyle(termui.ColorWhite)

	tableMem := widgets.NewTable()
	tableMem.Title = " Memory values "
	tableMem.RowSeparator = false
	tableMem.TextStyle = termui.Theme.Table.Text
	tableMem.TextAlignment = termui.AlignCenter
	tableMem.Rows = [][]string{
		[]string{"Used", "Free", "Total"},
		[]string{fmt.Sprintf("%.2f Gb", (float32(status.Memory.MemTotal)-float32(status.Memory.MemFree))/1000000), fmt.Sprintf("%.2f Gb", float32(status.Memory.MemFree)/1000000), fmt.Sprintf("%.2f Gb", float32(status.Memory.MemTotal)/1000000)},
	}
	tableMem.TextStyle = termui.NewStyle(termui.ColorWhite)
	tableMem.SetRect(50, 6, 100, 10)

	// CPU Frequency
	var freqBuffer = make([]float64, 40)
	freqBuffer[len(freqBuffer)-1] = float64(status.CPU.Freq) / 1000

	cpuFreqPlot := widgets.NewPlot()
	cpuFreqPlot.Title = " CPU frequency GHz "
	cpuFreqPlot.Data = [][]float64{freqBuffer}
	cpuFreqPlot.SetRect(0, 12, 50, 24)
	cpuFreqPlot.AxesColor = termui.ColorWhite
	cpuFreqPlot.LineColors[0] = termui.ColorCyan

	// Disk
	tableDisk := widgets.NewTable()
	tableDisk.Title = " Disk usage "
	tableDisk.RowSeparator = false
	tableDisk.TextStyle = termui.Theme.Table.Text
	tableDisk.TextAlignment = termui.AlignCenter
	tableDisk.Rows = BuildTableDisk(status.Disk.MountingPoints)
	tableDisk.TextStyle = termui.NewStyle(termui.ColorWhite)
	tableDisk.SetRect(50, 10, 100, 17)

	// Temperature
	var tempBuffer = make([]float64, 40)
	tempBuffer[len(tempBuffer)-1] = float64(status.Temperature.T)

	tempPlot := widgets.NewPlot()
	tempPlot.Title = " Temperature °C "
	tempPlot.Data = [][]float64{tempBuffer}
	tempPlot.SetRect(0, 24, 50, 32)
	tempPlot.AxesColor = termui.ColorWhite
	tempPlot.LineColors[0] = termui.ColorCyan

	// Network
	var netRxBuffer = make([]float64, 45)
	netRxBuffer[len(netRxBuffer)-1] = float64(status.Network.BytesRecv)
	var netTxBuffer = make([]float64, 45)
	netTxBuffer[len(netTxBuffer)-1] = float64(status.Network.BytesSent)

	netRx := widgets.NewSparkline()
	netRx.Title = fmt.Sprintf(" Rx:  %v B/s ", status.Network.BytesRecv)
	netRx.Data = netRxBuffer
	netRx.LineColor = termui.ColorRed

	netTx := widgets.NewSparkline()
	netTx.Title = fmt.Sprintf(" Tx:  %v B/s ", status.Network.BytesSent)
	netTx.Data = netTxBuffer
	netTx.LineColor = termui.ColorRed

	netPlot := widgets.NewSparklineGroup(netRx, netTx)
	netPlot.Title = " Network usage "
	netPlot.SetRect(50, 17, 100, 32)

	render := func() {
		status.Worker()

		g1.Percent = int(status.CPU.CPU0)
		g1.BarColor = GetColorPercent(status.CPU.CPU0)

		g2.Percent = int(status.CPU.CPU1)
		g2.BarColor = GetColorPercent(status.CPU.CPU1)

		g3.Percent = int(status.CPU.CPU2)
		g3.BarColor = GetColorPercent(status.CPU.CPU2)

		g4.Percent = int(status.CPU.CPU3)
		g4.BarColor = GetColorPercent(status.CPU.CPU3)

		gMemUsage.Percent = int(status.Memory.MemUsage)
		gMemUsage.BarColor = GetColorPercent(status.Memory.MemUsage)

		gSwapUsage.Percent = int(status.Memory.SwapUsage)
		gSwapUsage.BarColor = GetColorPercent(status.Memory.SwapUsage)

		tableMem.Rows = [][]string{
			[]string{"Used", "Free", "Total"},
			[]string{fmt.Sprintf("%.2f Gb", (float32(status.Memory.MemTotal)-float32(status.Memory.MemFree))/1000000), fmt.Sprintf("%.2f Gb", float32(status.Memory.MemFree)/1000000), fmt.Sprintf("%.2f Gb", float32(status.Memory.MemTotal)/1000000)},
		}

		// TODO : Can we use a sync pool for []float64 buffers ? as it's used everywhere
		freqBuffer = UpdateBuffer(freqBuffer, float64(status.CPU.Freq)/1000)
		cpuFreqPlot.Data = [][]float64{freqBuffer}

		tableDisk.Rows = BuildTableDisk(status.Disk.MountingPoints)

		tempBuffer = UpdateBuffer(tempBuffer, float64(status.Temperature.T))
		tempPlot.Data = [][]float64{tempBuffer}

		netRxBuffer = UpdateBuffer(netRxBuffer, float64(status.Network.BytesRecv))
		netTxBuffer = UpdateBuffer(netTxBuffer, float64(status.Network.BytesSent))
		netRx.Data = netRxBuffer
		netTx.Data = netTxBuffer
		netRx.Title = fmt.Sprintf(" Rx:  %v B/s ", status.Network.BytesRecv)
		netTx.Title = fmt.Sprintf(" Tx:  %v B/s ", status.Network.BytesSent)

		termui.Render(g1, g2, g3, g4, gMemUsage, gSwapUsage, tableMem, cpuFreqPlot, tableDisk, tempPlot, netPlot)
	}

	uiEvents := termui.PollEvents()
	ticker := time.NewTicker(1010 * time.Millisecond).C

	fmt.Print("\033[H\033[2J")

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-ticker:
			render()
		}
	}
}

// TODO : Improve this (pop) method to update buffer ?
func UpdateBuffer(inputBuffer []float64, inputValue float64) []float64 {
	history := inputBuffer[1:]
	updateBuffer := append(history, inputValue)
	return updateBuffer
}

// TODO : Improve this method ?
func BuildTableDisk(mountingPoints []core.MountingPoint) [][]string {
	var diskRows [][]string // TODO : pre allocation here ?
	header := []string{"Mount", "Size", "Used", "Free", "Usage"}
	diskRows = append(diskRows, header)

	for _, mp := range mountingPoints {
		diskRow := []string{mp.FileSystem, mp.Size, mp.Used, mp.Avail, mp.PercentUse}
		diskRows = append(diskRows, diskRow)
	}

	return diskRows
}

func GetColorPercent(inputValue float32) termui.Color {
	switch {
	case inputValue <= 50:
		return termui.ColorGreen
	case inputValue > 50 && inputValue <= 70:
		return termui.ColorYellow
	case inputValue > 70:
		return termui.ColorRed
	}
	return termui.ColorWhite
}
