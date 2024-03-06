package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(""))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("sensors_database")

	sensors := db.Collection("sensors")
	radiation := db.Collection("radiation_log")
	gases := db.Collection("gases_log")


	documento := bson.D{{"nome", "Fulano2"}, {"idade", 30}}
	resultadoInsercao, err := colecao.InsertOne(context.TODO(), documento)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Documento inserido com o ID: %v\n", resultadoInsercao.InsertedID)

	var documentoEncontrado bson.M
	err = colecao.Find(context.TODO(), bson.D{{"nome", "Fulano"}}).Decode(&documentoEncontrado)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Documento encontrado: %v\n", documentoEncontrado)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}