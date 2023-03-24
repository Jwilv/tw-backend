package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfile(ID string) (models.User, error) {

	context, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("tw")
	collection := db.Collection("users")

	var profile models.User

	objectId, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objectId,
	}

	err := collection.FindOne(context, condition).Decode(&profile)

	profile.Password = ""

	if err != nil {
		fmt.Println("registro no  encontrado" + err.Error())
		return profile, err
	}

	return profile, nil

}
