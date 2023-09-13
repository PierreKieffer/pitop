package core

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

// Platform represents information about the host system's platform.
type Platform struct {
	DistributorID string
	Release       int8
	Codename      string
}

// GetPlatform retrieves information about the host system's platform by running the "lsb_release -a" command.
// It populates and returns a Platform struct with the gathered information.
func GetPlatform() Platform {

	var p Platform

	var cmd *exec.Cmd = exec.Command("/bin/bash", "-c", "lsb_release -a")

	var out bytes.Buffer
	var errb bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	var outputLines []string = strings.Split(out.String(), "\n")
	for _, v := range outputLines {
		if strings.Contains(v, "Distributor") {
			var d string = strings.TrimSpace(strings.Split(v, ":")[1])
			p.DistributorID = d
		}
		if strings.Contains(v, "Release") {
			var rs string = strings.TrimSpace(strings.Split(v, ":")[1])
			r, _ := strconv.ParseInt(rs, 10, 8)
			p.Release = int8(r)
		}
		if strings.Contains(v, "Codename") {
			var c string = strings.TrimSpace(strings.Split(v, ":")[1])
			p.Codename = c
		}
	}
	return p
}
