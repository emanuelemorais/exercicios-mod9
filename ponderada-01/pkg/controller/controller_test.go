package controller

import (
	"testing"
	"fmt"
	"time"
	"reflect"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)	

func TestConnectBroker(t *testing.T) {
	client, _ := ConnectBroker()
	defer client.Disconnect(500)
	if !client.IsConnected() {
		t.Errorf("Unable to connect to MQTT broker\x1b[0m")
	}
}

func TestRandomValues(t *testing.T) {
	resultConfig := RandomValues()
	if reflect.TypeOf(resultConfig).Kind() != reflect.Float64 {
		t.Errorf("Random value is not float64")
	}
}

func TestCotroller(t *testing.T) {

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
		Controller()
		make(chan bool) <- true
	}()


	go func() {
		if token := client.Subscribe("mics6814", 1, nil); token.Wait() && token.Error() != nil {
			t.Logf("Error subscribing: %s", token.Error())
			return
		}
	}()

	time.Sleep(2 * time.Second)

	if len(data) == 0 {
		t.Errorf("No messages received")
	} else {
		for _, receipt := range data {
			t.Log(receipt)
		}
	}
}	