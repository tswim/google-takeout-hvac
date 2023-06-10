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
	SystemCapabilities       struct {
		HasStage1Cool          bool `json:"hasStage1Cool"`
		HasStage2Cool          bool `json:"hasStage2Cool"`
		HasStage3Cool          bool `json:"hasStage3Cool"`
		HasStage1Heat          bool `json:"hasStage1Heat"`
		HasStage2Heat          bool `json:"hasStage2Heat"`
		HasStage3Heat          bool `json:"hasStage3Heat"`
		HasStage1AlternateHeat bool `json:"hasStage1AlternateHeat"`
		HasStage2AlternateHeat bool `json:"hasStage2AlternateHeat"`
		HasHumidifier          bool `json:"hasHumidifier"`
		HasDehumidifier        bool `json:"hasDehumidifier"`
		HasDualFuel            bool `json:"hasDualFuel"`
		HasAuxHeat             bool `json:"hasAuxHeat"`
		HasEmergencyHeat       bool `json:"hasEmergencyHeat"`
		HasAirFilter           bool `json:"hasAirFilter"`
		HasFossilFuel          bool `json:"hasFossilFuel"`
		HasHotWaterControl     bool `json:"hasHotWaterControl"`
		HasHeatPump            bool `json:"hasHeatPump"`
		HasFan                 bool `json:"hasFan"`
	} `json:"systemCapabilities"`
	Cycles []struct {
		Caption struct {
			Locale         string `json:"locale"`
			Text           string `json:"text"`
			PlainText      string `json:"plainText"`
			AccessibleText string `json:"accessibleText"`
			Parameters     struct {
				StartTime time.Time `json:"startTime"`
				EndTime   time.Time `json:"endTime"`
			} `json:"parameters"`
		} `json:"caption"`
		StartTs         time.Time `json:"startTs"`
		Duration        string    `json:"duration"`
		IsComplete      bool      `json:"isComplete"`
		Heat1           bool      `json:"heat1"`
		Heat2           bool      `json:"heat2"`
		Heat3           bool      `json:"heat3"`
		HeatAux         bool      `json:"heatAux"`
		AltHeat         bool      `json:"altHeat"`
		AltHeat2        bool      `json:"altHeat2"`
		EmergencyHeat   bool      `json:"emergencyHeat"`
		Cool1           bool      `json:"cool1"`
		Cool2           bool      `json:"cool2"`
		Cool3           bool      `json:"cool3"`
		Fan             bool      `json:"fan"`
		FanCooling      bool      `json:"fanCooling"`
		Humidifier      bool      `json:"humidifier"`
		Dehumidifier    bool      `json:"dehumidifier"`
		AutoDehumdifier bool      `json:"autoDehumdifier"`
		WaterHeater     bool      `json:"waterHeater"`
	} `json:"cycles"`
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