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

	//variable para el perfil, le asignamos el modelo user
	var profile models.User

	//pasamos el string del id a objetId
	objectId, _ := primitive.ObjectIDFromHex(ID)

	//es la concion a cumplir, que la _id tenga el mismo valor que la id que llego por parametro
	condition := bson.M{
		"_id": objectId,
	}

	//buscamosun registro en la collection, le damos elcontext, la condition
	//despues decode en el espacio de memoria de profile
	// y si sale mal nos da err
	err := collection.FindOne(context, condition).Decode(&profile)

	//le asignamos unas password vacia aunque sea omitida
	profile.Password = ""

	//evaluamos sinos dio err y devolvemos el profile y el err
	if err != nil {
		fmt.Println("registro no  encontrado" + err.Error())
		return profile, err
	}

	//retornamos el profile y err en nil
	return profile, nil

}
