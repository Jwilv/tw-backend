package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var MongoCN = connectionDb()

var clientOptions = options.Client().ApplyURI("")


