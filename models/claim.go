package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	jwt "github.com/dgrijalva/jwt-go"

)

type Claim struct {
	Email string `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}