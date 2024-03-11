package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sensors struct {
	ID        string `json:"_id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Name      string `json:"name"`
}

func ReturnClient() *mongo.Client {

	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_DATABASE")))
	if err != nil {
		log.Fatal(err)
	}
	return client

}

func GetAllSensors() ([]Sensors, error) {

	client := ReturnClient()
	collection := client.Database("sensors_database").Collection("sensors")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var sensors []Sensors
	for cursor.Next(context.TODO()) {

		var doc bson.M
		err := cursor.Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := json.MarshalIndent(doc, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		var sensor Sensors
		err = json.Unmarshal(jsonData, &sensor)
		if err != nil {
			log.Fatal(err)
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func RegisterNewSensor(latitude float64, longitude float64, name string) {
	data := bson.D{{"latitude", latitude}, {"longitude", longitude}, {"name", name}}
	InsertDb("sensors", data)
}

func InsertDb(collectionName string, data interface{}) {
	client := ReturnClient()
	collection := client.Database("sensors_database").Collection(collectionName)
	insertResult, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
