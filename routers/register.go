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
	}

}
