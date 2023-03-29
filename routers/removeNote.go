package routers

import (
	"net/http"

	"github.com/Jwilv/tw-backend/db"
)

// elimina una nota de la base de datos
func RemoveNote(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "id invalido", http.StatusBadRequest)
		return
	}

	err := db.DeleteNote(ID, IDUser)

	if err != nil {
		http.Error(w, "error al intentar borrar la nota"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
