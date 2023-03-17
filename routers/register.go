package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jwilv/tw-backend/models"

)

// Register es la funcion que nos permite registrar un nuevo usuario en la base de datos
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error en los datos recibidos " + err.Error(), 400)
		return 
	}

	if len(user.Email) == 0 {
		http.Error(w, "Se requiere un email", 400)
	}

	if len(user.Password) >= 6 {
		http.Error(w, "la contraseña debe de ser de 6 caracteres minimo", 400)
	}
}

