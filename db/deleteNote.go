package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

// DeleteNote se encarga de buscar la nota en la base de datos mediante los parametros y eliminarla una vez la encuentre
func DeleteNote(ID string, UserID string) error {
	contextDB, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	Db := MongoCN.Database("tw")
	collection := Db.Collection("notes")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"userId": UserID,
	}

	_, err := collection.DeleteOne(contextDB, condition)

	return err
}
