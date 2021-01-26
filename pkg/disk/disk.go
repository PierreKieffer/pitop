package disk

import (
	"fmt"
	"github.com/PierreKieffer/pitop/pkg/utils"
	"os"
	"os/exec"
	"strings"
)

type DiskInfo struct {
	MountingPoint string
	Size          string
	Used          string
	Avail         string
	PercentUse    string
}

func ExtractDiskUsage() *[]DiskInfo {

	var disks []DiskInfo

	cmd := "df -h"
	run := exec.Command("sh", "-c", cmd)
	stdout, err := run.Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputLines := strings.Split(string(stdout), "\n")

	for _, outputLine := range outputLines {
		diskInfoSlice := utils.FormatStatSlice(strings.Split(outputLine, " "))
		if len(diskInfoSlice) > 0 {
			if diskInfoSlice[0][:4] == "/dev" && diskInfoSlice[0][:9] != "/dev/loop" {

				var diskInfo DiskInfo

				diskInfo.MountingPoint = diskInfoSlice[len(diskInfoSlice)-1]
				diskInfo.Size = diskInfoSlice[1]
				diskInfo.Used = diskInfoSlice[2]
				diskInfo.Avail = diskInfoSlice[3]
				diskInfo.PercentUse = diskInfoSlice[4]

				disks = append(disks, diskInfo)
			}
		}
	}

	return &disks

}
