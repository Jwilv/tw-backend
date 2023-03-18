package db

import (
	"context"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"

)

// CheckExisEmail verificamos si ya existe un usuario con el email que de parametro
func CheckExisEmail(email string) (models.User, bool, string) {
	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("users")

	condition := bson.M{"email":email}

	var result models.User

	err := collection.FindOne(context, condition).Decode(&result)

	ID := result.ID.Hex()


	
}
