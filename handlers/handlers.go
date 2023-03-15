package handlers

import (
	"os"

	"github.com/gorilla/mux"
)

func Drivers() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
}
