package middlewares

import (
	"net/http"

	"github.com/Jwilv/tw-backend/db"

)

// me permite conocer el estado de la base de datos 
func CheckDb(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.ChekingConnection() {
			http.Error(w, "Se perdio la conexion con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
