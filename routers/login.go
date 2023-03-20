package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/models"

)

// Login es la funcion que nos permite logear al user
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(user.Email) < 6 {
		http.Error(w, "Se requiere el email del usuario", 400)
		return
	}

	documet, exist := db.IntentLogin(user.Email, user.Password)
	if !exist {
		http.Error(w, "Usuario invalido", 400)
		return
	}
	jwtKey, err := jwt.GenerateJwt(documet)
	if err != nil {
		http.Error(w, "Error al intentar generar el token "+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}


}
