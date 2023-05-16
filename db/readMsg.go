package db

import (
	"context"
	"log"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func ReadMsg(ID string, page int64) ([]*models.ReturnNotes, bool) {

	contextDb, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("msgAdm")

	var result []*models.ReturnNotes

	condition := bson.M{
		"userId": ID,
	}

	options := options.Find()
	options.SetSkip((page - 1) * 1)
	options.SetLimit(1)
	options.SetSort(bson.D{{Key: "date", Value: -1}})


	cursor, err := collection.Find(contextDb, condition, options)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var document models.ReturnNotes
		err := cursor.Decode(&document)

		if err != nil{
			return result, false 
		}

		result = append(result, &document )
	}

	return result, true 

}
