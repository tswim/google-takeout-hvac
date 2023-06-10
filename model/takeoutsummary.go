package model

import (
	"time"
);

type TakeOutSummary struct {
	StartTs                  time.Time `json:"startTs"`
	EndTs                    time.Time `json:"endTs"`
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