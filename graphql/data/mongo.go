package data

import (
	"context"
	"fmt"
	"time"

	"github.com/kumarankeerthi/go-learning/graphql/core"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	DB *mongo.Database
}

func CreateRepository(server string, dbName string) *MongoRepo {
	ctx, _ := context.WithTimeout(context.Background(), 18*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+server))
	if err != nil {
		fmt.Println("error connecting to db")
	}
	db := client.Database(dbName)
	return &MongoRepo{
		DB: db,
	}
}

func (mp *MongoRepo) SavePerson(p core.Person) (string, error) {
	ctx := context.TODO()
	personColl := mp.DB.Collection("Person")
	result, err := personColl.InsertOne(ctx, p)
	if err != nil {
		fmt.Println("Error saving data into db")
		return "Failed", err
	}
	pId := cast.ToString(result.InsertedID)
	return pId, nil
}

func (mp *MongoRepo) GetPersonById(pid string) core.Person {
	id, err := primitive.ObjectIDFromHex(pid)
	if err != nil {
		fmt.Println("error fetching person with id :", pid)
	}
	personColl := mp.DB.Collection("Person")
	person := personColl.FindOne(context.TODO(), bson.M{"_id": id})
	var p core.Person
	person.Decode(&p)
	return p
}

func (mp *MongoRepo) GetPersons() []core.Person {
	return nil
}
