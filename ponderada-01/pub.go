package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
	"encoding/json"
	"io/ioutil"
	"os"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type SensorConfig struct {
	Sensor            string  `json:"sensor"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	QoS               byte     `json:"qos"`
	Unit              string  `json:"unit"`
	TransmissionRate  int     `json:"transmission_rate"`
}	

type SendData struct {
	Sensor string `json:"sensor"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	QoS byte `json:"qos"`
	Unit string `json:"unit"`
	TransmissionRate int	`json:"transmission_rate"`
	CurrentTime time.Time `json:"current_time"`
	Values gasesValues `json:"values"`
}

type gasesValues struct {
	CarbonMonoxide float64 `json:"carbon_monoxide"`
	NitrogenDioxide float64 `json:"nitrogen_dioxide"`
	Ethanol float64 `json:"ethanol"`
	Hydrogen float64 `json:"hydrogen"`
	Ammonia float64 `json:"ammonia"`
	Methane float64 `json:"methane"`
	Propane float64 `json:"propane"`
	IsoButane float64 `json:"iso_butane"`
}

type GasValues struct {
	MaxValue float64 `json:"max_value"`
	MinValue float64 `json:"min_value"`
}

var gasesRange = map[string]GasValues{
	"carbon_monoxide": {1, 1000},
	"nitrogen_dioxide": {0.05, 10},
	"ethanol":          {10, 500},
	"hydrogen":         {1, 1000},
	"ammonia":          {1, 500},
	"methane":          {1001, 9999}, // ">1000 ppm"
	"propane":          {1001, 9999}, // ">1000 ppm"
	"iso_butane":       {1001, 9999}, // ">1000 ppm"
}

func randomValues(gas string) (float64) {
	rand.Seed(time.Now().UnixNano()) // Inicializa a semente do gerador de números aleatórios

	maxValue := gasesRange[gas].MaxValue
	minValue := gasesRange[gas].MinValue
	value := rand.Float64()*(maxValue-minValue) + minValue
	return math.Round(value * 100) / 100
}

func createGasesValues() (gasesValues) {
	data := gasesValues{
		CarbonMonoxide: randomValues("carbon_monoxide"),
		NitrogenDioxide: randomValues("nitrogen_dioxide"),
		Ethanol: randomValues("ethanol"),
		Hydrogen: randomValues("hydrogen"),
		Ammonia: randomValues("ammonia"),
		Methane: randomValues("methane"),
		Propane: randomValues("propane"),
		IsoButane: randomValues("iso_butane"),
	}
	return data
}

func readConfigs() (SensorConfig, error){

	sensorConfig := os.Args[1]

	configData, err := ioutil.ReadFile(sensorConfig)
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

	config, err := readConfigs()
	if err != nil {
		fmt.Println("Erro ao decodificar o arquivo de configuração:", err)
	}	


	for {

		// Cria a estrutura de dados para enviar ao broker MQTT
		senddata := SendData{
		Sensor: config.Sensor,
		Latitude: config.Latitude,
		Longitude: config.Longitude,
		QoS: config.QoS,
		Unit: config.Unit,
		TransmissionRate: config.TransmissionRate,
		CurrentTime: time.Now(),
		Values: createGasesValues(),
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