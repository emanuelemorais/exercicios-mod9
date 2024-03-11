package rxwlib900

import (
	Common "ponderada-04/internal/sensors/common"
	"time"
)

type RadiationValues struct {
	SensorID        string  `json:"sensor_id"`
	TimeStamp       time.Time  `json:"timestamp"`
	Radiation float64 `json:"radiation"`
}

var radiationRange = map[string]Common.MaxMin{
	"radiation": {1, 1280},
}

func CreateGasesValues(id string) RadiationValues {
	radiationData := RadiationValues{
		SensorID:        id,
		TimeStamp:       time.Now(),
		Radiation: Common.RandomValues(radiationRange, "radiation"),
	}
	
	return radiationData
}
