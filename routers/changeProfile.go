package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/models"

)

// el endpoint que se encarga de modificar el perfil
func ChangeProfile(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil{
		http.Error(w,"error al procesar la informacion" + err.Error(), 400)
		return
	}

	status, errUpdate := db.ModifyRegister(user, IDUser) 

	if errUpdate != nil{
		http.Error(w,"ocurrio un error  al intentar modificar el registro. intentelo nuevamente. error : " + errUpdate.Error(), 400)
		return 
	}

	if !status{
		http.Error(w,"ocurrio un error al intentar modificar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	
}
