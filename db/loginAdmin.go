package db

import (
	"github.com/Jwilv/tw-backend/models"
	"golang.org/x/crypto/bcrypt"
)

// IntentLogin es la funcion que noos permite logearnos
func LoginAdmin(email string, password string) (models.User, bool) {
	user, exist, _ := CheckExisEmail(email)
	if !exist {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordBytesDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordBytesDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
