package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBName = "local"
const URI = "mongodb://127.0.0.1:27017"

func InitDB(ctx context.Context) *mongo.Database {

	clientOpts := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		fmt.Println("Error connecting to db", err)
	}

	db := client.Database(DBName)
	fmt.Println(db.Name())
	return db

}
