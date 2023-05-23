package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jwilv/tw-backend/db"

)

func GetUsersUnfollow(w http.ResponseWriter, r *http.Request) {


	result, status := db.GetUsers(IDUser, 1, "", "new")

	if !status {
		http.Error(w, "ocurio un error al obtener el listado de usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
