package takeout

import (
	"google-takeout/hvacparser/takeout/model"
)

type YM struct {
	Name       string
	Stage1Heat float64
	Stage2Heat float64
	Stage3Heat float64
	AuxHeat    float64
	Stage1Cool float64
	Stage2Cool float64
	Stage3Cool float64
	Humidifier float64
}

func newYM(name string) YM {
	
	return YM{name, 0, 0, 0, 0, 0, 0, 0, 0}
}

type Thermostat struct {
	Name         string
	Capabilities model.SystemCapabilities
	Runtimes     []YM
	Starts       []YM
}

func newStat(name string) Thermostat {
	t := model.SystemCapabilities{}
	return Thermostat{name, t, []YM{}, []YM{}}
}
