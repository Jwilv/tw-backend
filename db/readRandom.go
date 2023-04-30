package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReadRandomNotes() []models.ReturnNotes {
	context, cancel := context.WithTimeout( context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("notes")

	pipeline := mongo.Pipeline{
		// Obtén una muestra aleatoria de 4 documentos
		{{Key : "$sample", Value : bson.M{"size": 4}}},
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

	// Imprime las notas aleatorias obtenidas
	for _, note := range notes {
		fmt.Println(note.ID, note.Message, note.Date)
	}

	return notes
}