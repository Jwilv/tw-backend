package routers

import (
	"errors"
	"strings"

	"github.com/Jwilv/tw-backend/db"
	"github.com/Jwilv/tw-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string

var IDUser string

// procesa si el token recivido es valido
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	key := []byte("oqwepFJO03-49RFSAFN0123@$WEWQAsfoajf")

	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tokenVerify, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return claims, false, string(""), err
	}

	if !tokenVerify.Valid {
		return claims, false, string(""), errors.New("token invaldio")
	}

	_, exist, id := db.CheckExisEmail(claims.Email)

	if !exist {
		return claims, false, string(""), errors.New("no se encontro un user con el token")
	}

	Email = claims.Email
	IDUser = id

	return claims, exist, IDUser, nil

}
