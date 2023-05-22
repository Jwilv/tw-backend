package db

import (
	"context"
	"log"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RandomUsers() []models.User {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("users")

	pipeline := mongo.Pipeline{
		// Obtén una muestra aleatoria de 4 documentos
		{{Key: "$sample", Value: bson.M{"size": 4}}},
	}

	cursor, err := collection.Aggregate(context, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context)

	var user []models.User
	if err = cursor.All(context, &user); err != nil {
		log.Fatal(err)
	}

	return user
}
