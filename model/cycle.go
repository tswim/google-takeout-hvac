package model

import (
	"time"
);

type Cycle struct {
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
}