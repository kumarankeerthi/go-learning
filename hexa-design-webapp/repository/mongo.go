package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/kumarankeerthi/go-learning/hexa-design-webapp/chat"
)

type mongoRepositoy struct {
	DB *mongo.Database
}

func NewConnection(dbAddress string, dbName string) chat.ChatRepository {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbAddress))
	if err != nil {
		fmt.Println("Error connecting to mongo")
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error connecting to mongo")
	}
	db := client.Database(dbName)
	mongoRepo := &mongoRepositoy{
		DB: db,
	}
	return mongoRepo
}

func (m *mongoRepositoy) ViewChat(sender string, recipient string) ([]chat.Chat, error) {
	fmt.Println("at the mongo end")
	return nil, nil
}

func (m *mongoRepositoy) SendMessage(sender string, recipient string, message string) error {
	fmt.Println("at the mongo end")
	return nil
}
