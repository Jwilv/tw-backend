package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jwilv/tw-backend/db"

)

func UsersList(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")
	typeUser := r.URL.Query().Get("type")
	search := r.URL.Query().Get("search")

	pageNum, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "page debe ser un entero y tiene que ser mayor que 0", http.StatusBadRequest)
		return
	}

	pag := int64(pageNum)

	result, status := db.GetUsers(IDUser, pag, search, typeUser)

	if !status {
		http.Error(w, "ocurio un error al obtener el listado de usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
