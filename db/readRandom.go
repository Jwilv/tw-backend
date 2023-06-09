package db

import (
	"context"
	"log"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadRandomNotes(num int64) []models.ReturnNotes {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("notes")

	pipeline := mongo.Pipeline{
		// Obtén una muestra aleatoria de 4 documentos
		{{Key: "$sample", Value: bson.M{"size": num}}},
	}

	cursor, err := collection.Aggregate(context, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context)

	var notes []models.ReturnNotes
	if err = cursor.All(context, &notes); err != nil {
		log.Fatal(err)
	}

	return notes
}
