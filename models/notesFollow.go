package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotesFollow struct {
	ID             primitive.ObjectID `bson:"_ id" json:"id,omitempty"`
	UserId         string             `bson:"userId" json:"userId,omitempty"`
	UserRelationId string             `bson:"userRelationId" json:"userRelationId,omitempty"`
	Note           struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_ id" json:"id,omitempty"`
	}
}
