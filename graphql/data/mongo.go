package data

import (
	"fmt"

	"github.com/kumarankeerthi/go-learning/graphql/core"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	DB *mongo.Database
}

func CreateRepository(dbName string) *MongoRepo {

	return &MongoRepo{
		DB: nil,
	}
}

func (mp *MongoRepo) SavePerson(p core.Person) (string, error) {
	return "", nil
}

func (mp *MongoRepo) GetPersonById(pid string) core.Person {
	fmt.Println("in data layer")

	return core.Person{}
}

func (mp *MongoRepo) GetPersons() []core.Person {
	return nil
}
