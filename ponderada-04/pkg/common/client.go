package common

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	godotenv "github.com/joho/godotenv"
)

const IdPublisher = "go-mqtt-publisher"
const IdSubscriber = "go-mqtt-subscriber"

var Handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received: %s on topic %s\n", msg.Payload(), msg.Topic())
	return
}

func CreateClient(id string, callback_handler mqtt.MessageHandler) mqtt.Client {

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID(id)
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.SetClientID(id)
	opts.SetDefaultPublishHandler(callback_handler)

	return mqtt.NewClient(opts)
}
