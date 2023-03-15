package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectionDb()

var clientOptions = options.Client().ApplyURI("mongodb+srv://tw:Ofsw6KqbfoZol4k2@cluster0.sxvlsh2.mongodb.net/tw")

func ConnectionDb() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("connection DB")
	return client
}
