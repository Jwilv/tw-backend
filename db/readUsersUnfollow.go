package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Jwilv/tw-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func ReadUsersUnfollow(ID string, page int64, search string, typee string) ([]*models.User, bool) {
	contextDB, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("tw")
	collection := db.Collection("users")

	var result []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 4)
	findOptions.SetLimit(4)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := collection.Find(contextDB, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	var include bool

	for cursor.Next(contextDB) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println(err.Error())
			return result, false
		}

		var relationCheck models.Relation

		relationCheck.UserID = ID
		relationCheck.UserRelationID = user.ID.Hex()

		include = false

		found,_ := FindRelation(relationCheck)

		if typee == "new" && !found {
			include = true
		}

		if include {
			user.Password = ""
			user.Biography = ""
			user.Website = ""
			user.Location = ""
			user.Email = ""
			user.Banner = ""

			result = append(result, &user)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}
	cursor.Close(contextDB)
	return result, true
}
