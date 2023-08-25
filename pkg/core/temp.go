package core

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Temperature struct {
	T float32
}

func (temperature *Temperature) ExtractTemp() {

	tempBytes, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		panic(err)
	}

	tempInfo, _ := strconv.ParseUint(strings.Split(string(tempBytes), "\n")[0], 10, 64)

	temperature.T = float32(tempInfo) / 1000
}
