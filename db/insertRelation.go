package db

import (
	"context"
	"time"

	"github.com/Jwilv/tw-backend/models"

)

//inserta la relacion en la base de datos
func InsertRelation(relation models.Relation) (bool, error) {
	contextDb, cancel := context.WithTimeout(context.Background(), time.Second * 15)
	defer cancel()

	Db := MongoCN.Database("tw")
	collection := Db.Collection("relation")

	_, err := collection.InsertOne(contextDb, relation)

	if err != nil {
		return false, err
	}

	return true, nil
}
