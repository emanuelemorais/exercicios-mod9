package rxwlib900

import (
	Common "ponderada-02/internal/common"
)

type RadiationValues struct {
	Radiation float64 `json:"radiation"`
}

type SensorConfig struct {
	Sensor          string          `json:"sensor"`
	Unit            string          `json:"unit"`
	RadiationValues RadiationValues `json:"radiation-values"`
}

var radiationRange = map[string]Common.MaxMin{
	"radiation": {1, 1280},
}

func CreateGasesValues() SensorConfig {
	radiationData := RadiationValues{
		Radiation: Common.RandomValues(radiationRange, "radiation"),
	}
	sensorData := SensorConfig{
		Sensor:          "RXWLIB900",
		Unit:            "W/m2",
		RadiationValues: radiationData,
	}
	return sensorData
}
