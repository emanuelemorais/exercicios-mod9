package controller

import (
	"reflect"
	"testing"
	"time"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)


func TestRandomValues(t *testing.T) {
	resultConfig := RandomValues()
	if reflect.TypeOf(resultConfig).Kind() != reflect.Float64 {
		t.Errorf("Random value is not float64")
	}
}

func TestReceivingMessage(t *testing.T) {

	var data []string
	var expectedQoS = byte(1)

	var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		msgData := string(msg.Payload())
		data = append(data, msgData)
		
		if msg.Qos() != expectedQoS { //Confere QOS das mensagens recebidas
			t.Errorf("Expected QoS %d, received %d", expectedQoS, msg.Qos())
			t.FailNow()
		}
	}

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("sensors", expectedQoS, nil); token.Wait() && token.Error() != nil {
		t.Logf("Error subscribing: %s", token.Error())
		return
	}

	time.Sleep(5 * time.Second)

	// Verifica o recebimento das mensagens pelo tópico "sensors"
	if len(data) == 0 {
		t.Errorf("No messages received")
	} else {
		t.Log("Message received successfully")

	}

	// Verifica se o tempo de execução recebe a quatidade de mensagens esperadas de acordo com o perfil de QoS
	// O disparo de mensagens é feito a cada 1 segundo, logo, espera-se que 5 mensagens sejam recebidas
	if len(data) != 5 {
		t.Errorf("Expected 5 messages, received %d", len(data))
	} else {
		t.Log("Received 5 messages")
	}



}
