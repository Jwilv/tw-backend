package middlewares

import (
	"net/http"
	"github.com/Jwilv/tw-backend/routers"

)

//validacion del token

func ValidateJwtAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessAdmin(r.Header.Get("x-token"))
		if err != nil {
			http.Error(w, "Error en el toekn ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
