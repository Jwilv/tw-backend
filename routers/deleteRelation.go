package routers

import (
	"net/http"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/models"

)

func DeleteRelation(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w,"Id invalido", http.StatusBadRequest)
		return
	} 

	var relation models.Relation

	relation.UserID = IDUser
	relation.UserRelationID = ID

	status, err := db.RemoveRelation(relation)

	if err != nil || !status{
		http.Error(w,"se produjo un error no esperado al intentar remover la relacion de la base de datos:" + err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}