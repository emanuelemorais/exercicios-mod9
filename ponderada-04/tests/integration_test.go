package integration

import (
	"fmt"
	//"os"
	"testing"
	"time"
	DefaultClient "ponderada-03/pkg/common"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	//godotenv "github.com/joho/godotenv"
)

func TestIntegration(t *testing.T) {

	client := DefaultClient.CreateClient("client-integration", DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Logf("Error subscribing: %s", token.Error())
		panic(token.Error())
	}

	var data []string
	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		msgData := string(msg.Payload())
		data = append(data, msgData)

		fmt.Printf("Recebido: %s do tópico: %s com QoS: %d\n", msg.Payload(), msg.Topic(), msg.Qos())
	}

	if token := client.Subscribe("sensors", 1, messagePubHandler); token.Wait() && token.Error() != nil {
		t.Logf("Error subscribing: %s", token.Error())
		return
	}

	time.Sleep(5 * time.Second)
	client.Disconnect(250)

	if len(data) == 0 {
		t.Errorf("No message was received")
	}
}
