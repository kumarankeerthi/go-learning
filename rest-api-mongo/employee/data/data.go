package data

import (
	"context"
	"log"

	"github.com/kumarankeerthi/go-learning/rest-api-mongo/employee/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CollectionHandle(dbname string, collectionName string) *mongo.Collection {
	client := establishConnection()
	collection := client.Database(dbname).Collection(collectionName)
	return collection
}

func establishConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// hostnames := []string{"mongodb://localhost:27017"}
	// clientOptions.SetHosts(hostnames)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("mongo connection failed!!")
		return nil
	}
	return client
}

func FindEmployeeByID(id string, collection *mongo.Collection) model.Employee {
	filter := bson.D{{"id", "3"}}
	var result model.Employee
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
