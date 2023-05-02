package db

import (
	"context"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertNote graba la nota en la base de datos
func InsertNote(note models.RegisterNote) (string, bool, error) {

	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("notes")

	register := bson.M{
		"userId":  note.UserId,
		"message": note.Message,
		"date":    note.Date,
		"name":    note.Name,
	}

	result, err := collection.InsertOne(context, register)

	if err != nil {
		return string(""), false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil
}
