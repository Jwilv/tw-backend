package jwt

import (
	"time"

	"github.com/Jwilv/tw-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJwt(user models.User) (string, error) {
	key := []byte("oqwepFJO03-49RFSAFN0123@$WEWQAsfoajf")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"surname":   user.Surname,
		"biography": user.Biography,
		"location":  user.Location,
		"website":   user.Website,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
