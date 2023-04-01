package models

//modelo que usamos para grabar una relacion entre un user y otro
type Relation struct {
	UserID         string `bson:"userId" json:"userId"`
	UserRelationID string `bson:"userRelationId" json:"userRelationId"`
}
