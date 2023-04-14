package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/jwt"
	"github.com/Jwilv/tw-backend/models"
)

// Register es la funcion que nos permite registrar un nuevo usuario en la base de datos
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error en los datos recibidos " + err.Error(), 400)
		return 
	}

	if len(user.Email) == 0 {
		http.Error(w, "Se requiere un email", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "la contraseña debe de ser de 6 caracteres minimo", 400)
		return
	}

	_,exist,_ := db.CheckExisEmail(user.Email)

	if exist{
		http.Error(w,"Este email ya esta en uso ", 400)
		return
	}

	_,status, err := db.RegisterUser(user)

	if err != nil {
		http.Error(w, "Ocurrio un error al registar el usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se pudo registrar el usuario ", 400)
		return
	}

	documet, exist := db.IntentLogin(user.Email, user.Password)
	if !exist {
		http.Error(w, "Usuario invalido", 400)
		return
	}
	jwtKey, err := jwt.GenerateJwt(documet)
	if err != nil {
		http.Error(w, "Error al intentar generar el token "+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}

