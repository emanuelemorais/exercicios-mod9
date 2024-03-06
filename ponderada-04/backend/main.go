package main

import (
	"fmt"
	"log"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"os"
	godotenv "github.com/joho/godotenv"
)

type KafkaConsumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewKafkaConsumer(configMap *ckafka.ConfigMap, topics []string) *KafkaConsumer {
	return &KafkaConsumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (c *KafkaConsumer) Consume(msgChan chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		log.Printf("Error creating kafka consumer: %v", err)
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		log.Printf("Error subscribing to topics: %v", err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		log.Printf("Message on %s: %s", msg.TopicPartition, string(msg.Value))
		if err == nil {
			msgChan <- msg
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	msgChan := make(chan *ckafka.Message)

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"sasl.mechanisms":    "PLAIN",
		"security.protocol":  "SASL_SSL",
		"sasl.username":      os.Getenv("CONFLUENT_API_KEY"),
		"sasl.password":      os.Getenv("CONFLUENT_API_SECRET"),
		"session.timeout.ms": 6000,
		"group.id":           "manu",
		"auto.offset.reset":  "latest",
	}

	kafkaRepository := NewKafkaConsumer(configMap, []string{"sensors_queue"}) 

	go func() {
		if err := kafkaRepository.Consume(msgChan); err != nil {
			log.Printf("Error consuming kafka queue: %v", err)
		}
	}()


	for msg := range msgChan {
		log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		log.Printf("teste")
	}
}