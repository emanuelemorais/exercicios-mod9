package mics6814

import (
	Common "ponderada-04/internal/sensors/common"
	"time"
)

type GasesValues struct {
	SensorID        string  `json:"sensor_id"`
	TimeStamp       time.Time  `json:"timestamp"`
	CarbonMonoxide  float64 `json:"carbon_monoxide"`
	NitrogenDioxide float64 `json:"nitrogen_dioxide"`
	Ethanol         float64 `json:"ethanol"`
	Hydrogen        float64 `json:"hydrogen"`
	Ammonia         float64 `json:"ammonia"`
	Methane         float64 `json:"methane"`
	Propane         float64 `json:"propane"`
	IsoButane       float64 `json:"iso_butane"`
}

var gasesRange = map[string]Common.MaxMin{
	"carbon_monoxide":  {1, 1000},
	"nitrogen_dioxide": {0.05, 10},
	"ethanol":          {10, 500},
	"hydrogen":         {1, 1000},
	"ammonia":          {1, 500},
	"methane":          {1001, 9999}, // ">1000 ppm"
	"propane":          {1001, 9999}, // ">1000 ppm"
	"iso_butane":       {1001, 9999}, // ">1000 ppm"
}

func CreateGasesValues(id string) GasesValues {
	gasesData := GasesValues{
		SensorID:        id,
		TimeStamp:       time.Now(),
		CarbonMonoxide:  Common.RandomValues(gasesRange, "carbon_monoxide"),
		NitrogenDioxide: Common.RandomValues(gasesRange, "nitrogen_dioxide"),
		Ethanol:         Common.RandomValues(gasesRange, "ethanol"),
		Hydrogen:        Common.RandomValues(gasesRange, "hydrogen"),
		Ammonia:         Common.RandomValues(gasesRange, "ammonia"),
		Methane:         Common.RandomValues(gasesRange, "methane"),
		Propane:         Common.RandomValues(gasesRange, "propane"),
		IsoButane:       Common.RandomValues(gasesRange, "iso_butane"),
	}
		
	return gasesData
}
