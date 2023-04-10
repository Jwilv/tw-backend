package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotesFollow struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId         string             `bson:"userId" json:"userId,omitempty"`
	UserRelationId string             `bson:"userRelationId" json:"userRelationId,omitempty"`
	Note           struct {
		Message string    `bson:"message" json:"message"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"id"`
	}
}
