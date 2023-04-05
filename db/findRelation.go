package db

import (
	"context"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func FindRelation(relation models.Relation)(bool,error) {

	contextDB, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("relation")

	filter := bson.M{
		"userId" : relation.UserID,
		"userRelationId" : relation.UserRelationID,
	}

	var result models.Relation

	err := collection.FindOne(contextDB,filter).Decode(&result)

	if err != nil{
		return false, err
	}

	return true, nil
}