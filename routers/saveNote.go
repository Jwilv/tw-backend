package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/models"

)

// SaveNote se encarga de guardar la nota en la base de datos
func SaveNote(w http.ResponseWriter, r *http.Request) {

	var message models.Note

	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, "error al extraer el mensaje "+err.Error(), 400)
		return
	}

	register := models.RegisterNote{
		UserId:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, errInsert := db.InsertNote(register)

	if errInsert != nil {
		http.Error(w, "ocurrio un error al ingresar el registro. intente nuevamente"+errInsert.Error(), 400)
		return
	}

	if !status {
		http.Error(w,"no se logro grabar la nota ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
