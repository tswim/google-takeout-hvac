package model

type SystemCapabilities struct {
	Stage1Cool      bool `json:"hasStage1Cool"`
	Stage2Cool      bool `json:"hasStage2Cool"`
	Stage3Cool      bool `json:"hasStage3Cool"`
	Stage1Heat      bool `json:"hasStage1Heat"`
	Stage2Heat      bool `json:"hasStage2Heat"`
	Stage3Heat      bool `json:"hasStage3Heat"`
	Stage1AltHeat   bool `json:"hasStage1AlternateHeat"`
	Stage2AltHeat   bool `json:"hasStage2AlternateHeat"`
	Humidifier      bool `json:"hasHumidifier"`
	Dehumidifier    bool `json:"hasDehumidifier"`
	DualFuel        bool `json:"hasDualFuel"`
	AuxHeat         bool `json:"hasAuxHeat"`
	EmergencyHeat   bool `json:"hasEmergencyHeat"`
	AirFilter       bool `json:"hasAirFilter"`
	FossilFuel      bool `json:"hasFossilFuel"`
	HotWaterControl bool `json:"hasHotWaterControl"`
	HeatPump        bool `json:"hasHeatPump"`
	Fan             bool `json:"hasFan"`
}
