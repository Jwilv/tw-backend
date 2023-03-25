package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jwilv/tw-backend/db"

)

func GetProfile(w http.ResponseWriter, r *http.Request) {

	//obtenemos el id
	ID := r.URL.Query().Get("id")

	//comprobamos que le id sea correcto, caso contrario retornamos error 
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	//si el id existe lo buscamos en la base de datos y retornamos el perfil 
	//caso contrario devolvemos error 
	profile, err := db.SearchProfile(ID)

	//si se da un err retornamos
	if err != nil {
		http.Error(w, "Ocurrio un error al buscar el registro "+err.Error(), 400)
		return
	}

	//si no hay error 
	//devolvemos un json con StatusCreated
	//devolviendo el profile en la response 
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)

}
