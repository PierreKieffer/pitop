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

func ExtractDiskUsage() *DiskInfo {

	var diskInfo DiskInfo

	cmd := "df -h"
	run := exec.Command("bash", "-c", cmd)
	stdout, err := run.Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputLines := strings.Split(string(stdout), "\n")

	for _, outputLine := range outputLines {
		diskInfoSlice := utils.FormatStatSlice(strings.Split(outputLine, " "))
		if len(diskInfoSlice) > 0 {
			if diskInfoSlice[len(diskInfoSlice)-1] == "/" {
				diskInfo.MountingPoint = "/"
				diskInfo.Size = diskInfoSlice[1]
				diskInfo.Used = diskInfoSlice[2]
				diskInfo.Avail = diskInfoSlice[3]
				diskInfo.PercentUse = diskInfoSlice[4]

				break
			}
		}
	}

	return &diskInfo

}
