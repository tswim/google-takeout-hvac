package model

type TakeOutSummary struct {
	StartTs                  CustomTime `json:"startTs"`
	EndTs                    CustomTime `json:"endTs"`
	EnergyWinner             string    `json:"energyWinner"`
	EnergyLeaf               bool      `json:"energyLeaf"`
	TotalHeatingSeconds      int       `json:"totalHeatingSeconds"`
	TotalCoolingSeconds      int       `json:"totalCoolingSeconds"`
	TotalFanCoolingSeconds   int       `json:"totalFanCoolingSeconds"`
	TotalHumidifierSeconds   int       `json:"totalHumidifierSeconds"`
	TotalDehumidifierSeconds int       `json:"totalDehumidifierSeconds"`
	RecentAverageUsedSeconds int       `json:"recentAverageUsedSeconds"`
	SecondsUsageOverAverage  int       `json:"secondsUsageOverAverage"`
	SystemCapabilities       SystemCapabilities `json:"SystemCapabilities"`
	Cycles [] 				 Cycle `json:"cycles"`
	CyclesIncomplete bool `json:"cyclesIncomplete"`
	Events           [] Event `json:"events"`
	EventsIncomplete bool  `json:"eventsIncomplete"`
	Rates            []any `json:"rates"`
}