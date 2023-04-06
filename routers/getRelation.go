package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/models"

)

func GetRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "id invalido", http.StatusBadRequest)
		return
	}

	var relation models.Relation
	relation.UserID = IDUser
	relation.UserRelationID = ID

	status, err := db.FindRelation(relation)

	if err != nil || !status {

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode()

}
