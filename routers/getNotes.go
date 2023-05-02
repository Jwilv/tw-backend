package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jwilv/tw-backend/db"
)

// obtengo las notas de la base de datos, segun el id proporcionado
func GetNotes(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "se requiere un id valido", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "se requiere una page valida", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "se necesita un numero entero en page, con un numero mayor a cero", http.StatusBadRequest)
		return
	}

	pageVerify := int64(page)

	result, status := db.ReadNotes(ID, pageVerify)

	if !status {
		http.Error(w, "no se pudo completar la accion, error al obtener las notas", 400)
		return
	}

	w.Header().Set("Content-type", "apllication/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
