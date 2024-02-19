package controller

import (
	"encoding/json"
	"fmt"
	"time"
	"math/rand"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/emanuelemorais/exercicios-mod9/ponderada-01/internal/mics"
	
)

type SensorConfig struct {
	Sensor           string  `json:"sensor"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	QoS              byte    `json:"qos"`
	Unit             string  `json:"unit"`
	TransmissionRate int     `json:"transmission_rate"`
}

type SendData struct {
	Sensor           string      `json:"sensor"`
	Latitude         float64     `json:"latitude"`
	Longitude        float64     `json:"longitude"`
	Unit             string      `json:"unit"`
	TransmissionRate int         `json:"transmission_rate"`
	CurrentTime      time.Time   `json:"current_time"`
	Values           mics.GasesValues `json:"values"`
}

func ConnectBroker() (MQTT.Client, error) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client, nil
}

func RandomValues() float64 {
	rand.Seed(time.Now().UnixNano()) 
	return rand.Float64() * 100 
}

func Controller() {

	client, err := ConnectBroker()
	if err != nil {
		fmt.Println("Erro ao conectar ao broker MQTT:", err)
	}	

	latitude := RandomValues()
	longitude := RandomValues()

	for {
		senddata := SendData{
			Sensor:           "MiCS-6814",
			Latitude:         latitude,
			Longitude:        longitude,
			Unit:             "ppm",
			CurrentTime:      time.Now(),
			Values:           mics.CreateGasesValues(),
		}

		jsonData, err := json.MarshalIndent(senddata, "", "    ")
		if err != nil {
			fmt.Println("Erro ao converter para JSON:", err)
			return
		}

		token := client.Publish("mics6814", 1, false, string(jsonData)) 
		token.Wait()
		fmt.Println("Publicado:", string(jsonData))
		time.Sleep(1 * time.Second) 
	}
}
