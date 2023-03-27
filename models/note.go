package models

//modelo de nota
type Note struct {
	Message string `bson:"message" json:"message"`
}
