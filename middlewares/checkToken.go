package middlewares

import (
	"net/http"

	"github.com/Jwilv/tw-backend/routers"

)

//validacion del token

func validateJwt(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("x-token"))

	}
}
