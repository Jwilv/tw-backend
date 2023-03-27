package db

import (
	"context"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"

)

// InsertNote graba la nota en la base de datos
func InsertNote(note models.RegisterNote) (string, bool, error) {

	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("notes")

	register := bson.M{
		"userId": note.userId,
		"message": note.message,
		"date":note.date,

	}
}
