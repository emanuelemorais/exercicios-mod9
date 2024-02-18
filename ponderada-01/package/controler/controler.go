package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	MQTT "github.com/eclipse/paho.mqtt.golang"
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
	QoS              byte        `json:"qos"`
	Unit             string      `json:"unit"`
	TransmissionRate int         `json:"transmission_rate"`
	CurrentTime      time.Time   `json:"current_time"`
	Values           string `json:"values"`
}


func ReadConfigs() (SensorConfig, error) {

	sensorConfig := "../../config/sensor-config.json"

	configData, err := os.ReadFile(sensorConfig)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo de configuração:", err)
		return SensorConfig{}, err
	}

	// Decodifica o conteúdo do arquivo JSON de configuração
	var config SensorConfig
	err = json.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println("Erro ao decodificar o arquivo de configuração:", err)
		return SensorConfig{}, err
	}

	return config, nil

}

func main() {

	// Conecta ao broker MQTT
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	config, err := ReadConfigs()
	if err != nil {
		fmt.Println("Erro ao decodificar o arquivo de configuração:", err)
	}

	for {

		// Cria a estrutura de dados para enviar ao broker MQTT
		senddata := SendData{
			Sensor:           config.Sensor,
			Latitude:         config.Latitude,
			Longitude:        config.Longitude,
			QoS:              config.QoS,
			Unit:             config.Unit,
			TransmissionRate: config.TransmissionRate,
			CurrentTime:      time.Now(),
		}

		jsonData, err := json.MarshalIndent(senddata, "", "    ")
		if err != nil {
			fmt.Println("Erro ao converter para JSON:", err)
			return
		}

		token := client.Publish(config.Sensor, config.QoS, false, string(jsonData))
		token.Wait()
		fmt.Println("Publicado:", string(jsonData))
		time.Sleep(time.Duration(config.TransmissionRate) * time.Second)
	}
}
