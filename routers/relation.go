package routers

import (
	"net/http"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/models"

)

func Relation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Id invalido", http.StatusBadRequest)
		return
	}

	var relation models.Relation

	relation.UserID = IDUser
	relation.UserRelationID = ID

	status, err := db.InsertRelation(relation)

	if err != nil || !status{
		http.Error(w,"error al insertar la relacion", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	

}
