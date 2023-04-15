package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/jwt"
	"github.com/Jwilv/tw-backend/models"
)

func RenewToken(w http.ResponseWriter, r *http.Request) {

	user, err, _ := db.CheckExisEmail(Email)

	if !err {
		http.Error(w, "Error al intentar encontrar  el user", http.StatusBadRequest)
		return
	}

	jwtKey, errToken := jwt.GenerateJwt(user)
	if errToken != nil {
		http.Error(w, "Error al intentar generar el token "+errToken.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
