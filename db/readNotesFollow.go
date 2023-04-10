package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"

)

// retorna una lista de las notas de las perosnas que el usuario sigue y un stado
func ReadNotesFollow(ID string, page int) ([]models.NotesFollow, bool) {

	contextDB, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("relation")

	skip := (page - 1) * 20



	conditions := make([]bson.M,0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userId":ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":"notes",
			"localField":"userRelationId",
			"foreignField" : "userId",
			"as":"note",
		}})

	conditions = append(conditions, bson.M{ "$unwind" : "$note"})
	conditions = append(conditions, bson.M{ "$sort" : bson.M{"note.date":-1}})
	conditions = append(conditions, bson.M{"$skip" : skip})
	conditions = append(conditions, bson.M{"$limit" : 20})

cursor, err := collection.Aggregate(contextDB, conditions)

var result []models.NotesFollow

if err != nil{
	fmt.Println(err.Error())
	return result, false 
}

errCursor := cursor.All(contextDB, &result)

if errCursor != nil{
	return result, false 
}

return result, true 

}
