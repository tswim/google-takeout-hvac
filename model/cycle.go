package model


type Cycle struct {
	Caption struct {
		Locale         string `json:"locale"`
		Text           string `json:"text"`
		PlainText      string `json:"plainText"`
		AccessibleText string `json:"accessibleText"`
		Parameters     struct {
			StartTime CustomTime `json:"startTime"`
			EndTime   CustomTime `json:"endTime"`
		} `json:"parameters"`
	} `json:"caption"`
	StartTs         CustomTime `json:"startTs"`
	Duration        Duration    `json:"duration"`
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