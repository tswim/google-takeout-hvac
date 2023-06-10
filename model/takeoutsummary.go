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
	Events           []struct {
		StartTs        time.Time `json:"startTs"`
		TimezoneOffset int       `json:"timezoneOffset"`
		Duration       string    `json:"duration"`
		Continuation   bool      `json:"continuation"`
		EventType      string    `json:"eventType"`
		SetPoint       struct {
			SetPointType    string `json:"setPointType"`
			ScheduleType    string `json:"scheduleType"`
			Preconditioning bool   `json:"preconditioning"`
			Targets         struct {
				HeatingTarget float64 `json:"heatingTarget"`
				CoolingTarget float64 `json:"coolingTarget"`
			} `json:"targets"`
			TouchedBy             string    `json:"touchedBy"`
			TouchedWhen           time.Time `json:"touchedWhen"`
			TouchedTimezoneOffset int       `json:"touchedTimezoneOffset"`
			TouchedWhere          string    `json:"touchedWhere"`
			TouchedUserID         string    `json:"touchedUserId"`
			ScheduledStart        int       `json:"scheduledStart"`
			ScheduledDay          int       `json:"scheduledDay"`
			PreviousType          string    `json:"previousType"`
			EmergencyHeatActive   bool      `json:"emergencyHeatActive"`
		} `json:"setPoint"`
	} `json:"events"`
	EventsIncomplete bool  `json:"eventsIncomplete"`
	Rates            []any `json:"rates"`
}