package controller

import (
	"testing"
	"time"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)


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

	// Verifica se o tempo de execução recebe a quatidade de mensagens esperadas de acordo com o perfil de QoS
	// O disparo de mensagens é feito a cada 1 segundo, logo, espera-se que 5 mensagens sejam recebidas
	if len(data) != 5 {
		t.Errorf("Expected 5 messages, received %d", len(data))
	} else {
		t.Log("Received 5 messages")
	}



}
