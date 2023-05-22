package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jwilv/tw-backend/db"
)

func GetRandomUsers(w http.ResponseWriter, r *http.Request){

	users := db.RandomUsers()

	w.Header().Set("Content-type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&users)
}