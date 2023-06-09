package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnNotes struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId  string             `bson:"userId" json:"userID,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
	Name    string             `bson:"name" json:"name,omitempty"`
}
