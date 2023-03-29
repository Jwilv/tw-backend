package routers

import (
	"net/http"

	"github.com/Jwilv/tw-backend/db"
)

// elimina una nota de la base de datos
func RemoveNote(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "id invalido", 400)
		return
	}

	err := db.DeleteNote(ID)

}
