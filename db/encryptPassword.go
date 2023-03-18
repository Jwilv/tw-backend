package db

import (
	"golang.org/x/crypto/bcrypt"

)

func EncryptPassword(password string) (string, error) {
	value := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), value)
	return string(bytes), err
}
