package model

type SystemCapabilities       struct {
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
}