package takeout

type YM struct {
	Name string
	Heat1 float64 
	Heat2 float64
	HeatAux float64
	Cool float64
	Humidifier float64
}

func newYM(name string) (YM) {
	return YM{name,0,0,0,0,0};
}

type Thermostat struct {
	Name string
	Capabilities Capabilities
	Runtimes [] YM
	Starts [] YM
}

func newStat(name string) (Thermostat) {
	return Thermostat{name, Capabilities{}, [] YM{}, [] YM{}};
}

type Capabilities struct {
	Heat1 bool 
	Heat2 bool
	HeatAux bool
	Cool bool
	Humidifier bool
}
