package integration

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	DefaultClient "ponderada-03/pkg/common"
	"regexp"
	"testing"
	"time"
)

func ValidateData(data string) bool {
	regexPattern := `{"gases":\{"sensor_id":"[^"]+","timestamp":"[^"]+","carbon_monoxide":[^,]+,"nitrogen_dioxide":[^,]+,"ethanol":[^,]+,"hydrogen":[^,]+,"ammonia":[^,]+,"methane":[^,]+,"propane":[^,]+,"iso_butane":[^}]+\},"radiation":\{"sensor_id":"[^"]+","timestamp":"[^"]+","radiation":[^}]+\}}`

	re := regexp.MustCompile(regexPattern)

	return re.MatchString(data)
}

func TestIntegration(t *testing.T) {

	client := DefaultClient.CreateClient("client-integration", DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Logf("Error subscribing: %s", token.Error())
		panic(token.Error())
	}

	var data []string
	var expectedQoS = byte(1)

	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		msgData := string(msg.Payload())
		data = append(data, msgData)

		if msg.Qos() != expectedQoS {
			t.Errorf("Expected QoS %d, received %d", expectedQoS, msg.Qos())
			t.FailNow()
		}

		if !ValidateData(string(msg.Payload())) {
			t.Errorf("Invalid pattern")
			t.FailNow()
		}

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

	receiptRate := float64(len(data)) / 5
	if receiptRate > 1.5 || receiptRate < 0.5{
		t.Errorf("Transmission rate is not within the expected range.")
	}


}
