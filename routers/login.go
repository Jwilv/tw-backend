package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/models"

)

// Login es la funcion que nos permite logear al user
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var user models.User



}
