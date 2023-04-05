package db

import (
	"context"
	"time"

	"github.com/Jwilv/tw-backend/models"

)

func RemoveRelation(relation models.Relation) (bool, error) {
	contextDB, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("relation")

	_, err := collection.DeleteOne(contextDB, relation)

	if err != nil{
		return false, err
	}

	return true, nil
}
