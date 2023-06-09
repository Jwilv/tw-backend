package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jwilv/tw-backend/db"

)

func GetNotesFollow(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")

	if len(page) < 1 {
	http.Error(w,"No se recivio la page", http.StatusBadRequest)
	return
	}

	pageNum, errPage := strconv.Atoi(page)

	if errPage != nil {
		http.Error(w, "page debe es invalida debe ser un entero mayor que 0", http.StatusBadRequest)
		return
	}

	pag := int(pageNum)

	results, status := db.ReadNotesFollow(IDUser, pag)

	if !status {
		http.Error(w, "Error al solicitar las notas", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}
