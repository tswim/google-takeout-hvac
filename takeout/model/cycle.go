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
	Duration        Duration   `json:"duration"`
	IsComplete      bool       `json:"isComplete"`
	Stage1Heat      bool       `json:"heat1"`
	Stage2Heat      bool       `json:"heat2"`
	Stage3Heat      bool       `json:"heat3"`
	AuxHeat         bool       `json:"heatAux"`
	Stage1AltHeat   bool       `json:"altHeat"`
	Stage2AltHeat   bool       `json:"altHeat2"`
	EmergencyHeat   bool       `json:"emergencyHeat"`
	Stage1Cool      bool       `json:"cool1"`
	Stage2Cool      bool       `json:"cool2"`
	Stage3Cool      bool       `json:"cool3"`
	Fan             bool       `json:"fan"`
	FanCooling      bool       `json:"fanCooling"`
	Humidifier      bool       `json:"humidifier"`
	Dehumidifier    bool       `json:"dehumidifier"`
	AutoDehumdifier bool       `json:"autoDehumdifier"`
	WaterHeater     bool       `json:"waterHeater"`
}
