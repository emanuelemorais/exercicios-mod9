package controller

import (
	"encoding/json"
	"fmt"
	MICS6814 "ponderada-04/internal/sensors/mics6814"
	RXWLIB900 "ponderada-04/internal/sensors/rxwlib900"
	DefaultClient "ponderada-04/pkg/common"
	"time"
)

type SendData struct {
	PayloadGases    MICS6814.GasesValues  `json:"gases"`
	PayloadRadiation RXWLIB900.RadiationValues `json:"radiation"`
}

func (s *SendData) ToJSON() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func Controller(id string) {

	client := DefaultClient.CreateClient(DefaultClient.IdPublisher, DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {

		senddata := SendData{
			PayloadGases:     MICS6814.CreateGasesValues(id),
			PayloadRadiation: RXWLIB900.CreateGasesValues(id),
		}
	
		payload, _ := senddata.ToJSON()

		token := client.Publish("sensors", 1, false, payload)
		token.Wait()

		fmt.Printf("Published message: %s\n", payload)

		time.Sleep(30 * time.Second)
	}
}
