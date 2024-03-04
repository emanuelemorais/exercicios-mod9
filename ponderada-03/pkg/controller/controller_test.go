package controller

import (
	"testing"
	"time"
	"regexp"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	DefaultClient "ponderada-03/pkg/common"
)

func validateData(data string) bool {
	regexPattern := `^\{"identifier":\d+,"latitude":\d+(\.\d+)?,"longitude":\d+(\.\d+)?,"current_time":"\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.\d{6}-\d{2}:\d{2}",` +
		`"gases-values":\{"sensor":"MiCS-6814","unit":"ppm","gases-values":\{"carbon_monoxide":\d+(\.\d+)?,"nitrogen_dioxide":\d+(\.\d+)?,` +
		`"ethanol":\d+(\.\d+)?,"hydrogen":\d+(\.\d+)?,"ammonia":\d+(\.\d+)?,"methane":\d+(\.\d+)?,"propane":\d+(\.\d+)?,"iso_butane":\d+(\.\d+)?\}\},` +
		`"radiation-values":\{"sensor":"RXWLIB900","unit":"W/m2","radiation-values":\{"radiation":\d+(\.\d+)?\}\}\}$`

	re := regexp.MustCompile(regexPattern)

	return re.MatchString(data)
}

func TestReceivingMessage(t *testing.T) {

	client := DefaultClient.CreateClient("client-test", DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Logf("Error subscribing: %s", token.Error())
		panic(token.Error())
	}

	var data []string
	var expectedQoS = byte(1)
	
	var messageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		msgData := string(msg.Payload())
		data = append(data, msgData)

		if msg.Qos() != expectedQoS { //Confere QOS das mensagens recebidas
			t.Errorf("Expected QoS %d, received %d", expectedQoS, msg.Qos())
			t.FailNow()
		}

		if !validateData(string(msg.Payload())) { //Confere se a mensagem recebida é válida
			t.Errorf("Invalid pattern")
		} 
	}

	if token := client.Subscribe("sensors", expectedQoS, messageHandler); token.Wait() && token.Error() != nil {
		t.Logf("Error subscribing: %s", token.Error())
		return
	}

	time.Sleep(5 * time.Second)
	client.Disconnect(250)

	// Verifica se o tempo de execução recebe a quatidade de mensagens esperadas de acordo com o perfil de QoS
	// O disparo de mensagens é feito a cada 1 segundo, logo, espera-se que 5 mensagens sejam recebidas
	if len(data) != 5 {
		t.Errorf("Expected 5 messages, received %d", len(data))
	} else {
		t.Log("Received 5 messages")
	}



}
