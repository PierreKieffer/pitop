package core

import (
	"os/exec"
	"strings"

	"github.com/PierreKieffer/pitop/pkg/utils"
)

type MountingPoint struct {
	FileSystem string
	Size       string
	Used       string
	Avail      string
	PercentUse string
}

type Disk struct {
	MountingPoints []MountingPoint
}

func (disk *Disk) ExtractDiskUsage() {

	disk.MountingPoints = []MountingPoint{} // TODO : Can we preallocate here ?

	cmd := "df -h"
	run := exec.Command("bash", "-c", cmd)
	stdout, err := run.Output()

	if err != nil {
		panic(err)
	}

	outputLines := strings.Split(string(stdout), "\n")

	for i := range outputLines {
		mountingPointInfoSlice := utils.FormatStatSlice(strings.Split(outputLines[i], " "))
		if len(mountingPointInfoSlice) > 0 {
			if mountingPointInfoSlice[0][:4] == "/dev" && mountingPointInfoSlice[0][:9] != "/dev/loop" {

				var mountingPoint MountingPoint

				mountingPoint.FileSystem = mountingPointInfoSlice[len(mountingPointInfoSlice)-1]
				mountingPoint.Size = mountingPointInfoSlice[1]
				mountingPoint.Used = mountingPointInfoSlice[2]
				mountingPoint.Avail = mountingPointInfoSlice[3]
				mountingPoint.PercentUse = mountingPointInfoSlice[4]

				disk.MountingPoints = append(disk.MountingPoints, mountingPoint)
			}
		}
	}
}
