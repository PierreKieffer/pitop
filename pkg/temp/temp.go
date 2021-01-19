package temp

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Temp struct {
	T float32
}

func ExtractTemp() *Temp {
	var temp Temp

	tempBytes, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tempInfo, _ := strconv.ParseUint(strings.Split(string(tempBytes), "\n")[0], 10, 64)

	temp.T = float32(tempInfo) / 1000

	return &temp
}
