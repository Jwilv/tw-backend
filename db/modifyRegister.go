package db

import (
	"context"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

func ModifyRegister(user models.User, ID string) (bool, error) {
	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("users")

	registerDoc := make(map[string]interface{})

	if len(user.Name) > 0 {
		registerDoc["name"] = user.Name
	}
	if len(user.Surname) > 0 {
		registerDoc["surname"] = user.Surname
	}
	if len(user.Email) > 0 {
		registerDoc["email"] = user.Email
	}
	if len(user.Avatar) > 0 {
		registerDoc["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		registerDoc["banner"] = user.Banner
	}
	if len(user.Biography) > 0 {
		registerDoc["biography"] = user.Biography
	}
	if len(user.Location) > 0 {
		registerDoc["location"] = user.Location
	}
	if len(user.Website) > 0 {
		registerDoc["website"] = user.Website
	}

	registerDoc["birthDate"] = user.BirthDate

	updateDoc := bson.M{
		"$set":registerDoc,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": bson.M{"$eq":objID},
	}

	_, err := collection.UpdateOne(context, filter, updateDoc)

	if err != nil{
		return false, err
	}

	return true, nil

}
