package main

import (
	"log"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/handlers"

)

func main() {
	if !db.ChekingConnection() {
		log.Fatal("Sin conexion a la base de datos")
	}
	handlers.Drivers()
}
