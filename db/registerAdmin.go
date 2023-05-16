package db

import (
	"context"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RegisterAdmin es la funcion que obtiene como parametro el user y lo registra en la base de datos
func RegisterAdmin(user models.User) (string, bool, error) {

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("tw")
	collection := db.Collection("admin")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := collection.InsertOne(context, user)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
