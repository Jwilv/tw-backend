package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jwilv/tw-backend/db"
)

func GetRandomNotes(w http.ResponseWriter, r *http.Request){

	notes := db.ReadRandomNotes(3)

	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&notes)
}