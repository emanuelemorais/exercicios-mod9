package controller

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"reflect"
	"testing"
	"time"
)

func TestRandomValues(t *testing.T) {
	resultConfig := RandomValues()
	if reflect.TypeOf(resultConfig).Kind() != reflect.Float64 {
		t.Errorf("Random value is not float64")
	}
}


func TestReceivingMessage(t *testing.T) {

	var data []string
	var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
		data = append(data, fmt.Sprintf("Recebido: %s do tópico: %s\n", msg.Payload(), msg.Topic()))
	}

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go func() {
		Controller(1)
	}()

	go func() {
		if token := client.Subscribe("sensors", 1, nil); token.Wait() && token.Error() != nil {
			t.Logf("Error subscribing: %s", token.Error())
			return
		}
	}()

	time.Sleep(5 * time.Second)

	// Verifica o recebimento das mensagens pelo tópico "sensors"
	if len(data) == 0 {
		t.Errorf("No messages received")
	} else {
		for _, receipt := range data {
			t.Log(receipt)
			t.Log("Message received successfully")
		}
	}

	// Verifica se o tempo de execução recebe a quatidade de mensagens esperadas de acordo com o perfil de QoS
	// O disparo de mensagens é feito a cada 1 segundo, logo, espera-se que 5 mensagens sejam recebidas
	if len(data) != 5 {
		t.Errorf("Expected 5 messages, received %d", len(data))
	} else {
		t.Log("Received 5 messages")
	}
	
}

