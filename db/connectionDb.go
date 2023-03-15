package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var MongoCN = connectionDb()

var clientOptions = options.Client().ApplyURI("")

func connectionDb() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
}
