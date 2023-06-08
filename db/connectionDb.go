package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

//MongoCN es el objeto se conexion a nuestra base de datos
var MongoCN = ConnectionDb()

var clientOptions = options.Client().ApplyURI("mongodb+srv://mongo:9wOn3E9enmNYRWpEw1IC@cluster0.sxvlsh2.mongodb.net/tw")

//ConnectionDb me permite conectarme a la data base
func ConnectionDb() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("connection DB")
	return client
}

//chequeo de la conexion a la base de datos
// y me conecto usando MongoCN
func ChekingConnection() bool{
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
