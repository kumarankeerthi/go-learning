package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/kumarankeerthi/go-learning/hexa-design-webapp/chat"
)

type mongoRepositoy struct {
	DB      *mongo.Database
	context context.Context
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
		DB:      db,
		context: ctx,
	}
	return mongoRepo
}

func (m *mongoRepositoy) ViewChat(sender string, recipient string) ([]chat.Messages, error) {
	ctx := context.TODO()
	fmt.Println("fecthing message sent from ", sender+" to", recipient)
	chatColl := m.DB.Collection("Chats")
	//msgColl := m.DB.Collection("Message")

	var result []bson.M
	//filter := bson.M{}
	filter := bson.M{"Sender": sender, "Recipient": recipient}

	f := bson.M{"_id": 1234}
	fmt.Println(filter)
	fmt.Println(f)

	cursor, err := chatColl.Find(ctx, filter)
	if err != nil {
		fmt.Println("error inserting record", err)
	}
	err = cursor.All(ctx, &result)
	var msgs []chat.Messages
	var msg chat.Messages
	for _, res := range result {
		for k, v := range res {
			switch k {
			case "MsgID":
				{
					msg = m.getMessageByID(v)
				}
			}

		}
		msgs = append(msgs, msg)
	}

	fmt.Println(msgs)
	return msgs, nil
}

func (m *mongoRepositoy) getMessageByID(id interface{}) chat.Messages {
	ctx := context.TODO()
	msgColl := m.DB.Collection("Message")
	fmt.Println(msgColl.Name())
	filter := bson.M{"_id": id}
	fmt.Println(`bson.M{"_id": docID}:`, filter)

	singlResult := msgColl.FindOne(ctx, filter)

	var result bson.M
	singlResult.Decode(&result)
	var msg chat.Messages
	msg.Message = fmt.Sprintf("%s", result["Message"])
	ts := cast.ToTime(result["Timestamp"])
	msg.Timestamp = ts.String()
	//err = cursor.All(ctx, &result)
	return msg
	//return fmt.Sprintf("%s", result["Message"])
}

func (m *mongoRepositoy) SendMessage(chat chat.Chat) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	fmt.Println("at the mongo end", chat)
	chatColl := m.DB.Collection("Chats")
	msgColl := m.DB.Collection("Message")
	chat.Timestamp = time.Now().String()
	msgId, err := msgColl.InsertOne(ctx, bson.D{
		{Key: "Message", Value: chat.Message},
		{Key: "Timestamp", Value: chat.Timestamp},
	})

	if err != nil {
		fmt.Println("error inserting record", err)
	}
	//id := fmt.Sprintf("%v", msgId.InsertedID)
	//fmt.Println("ID value is  -----", id)

	_, err = chatColl.InsertOne(ctx, bson.D{
		{Key: "Sender", Value: chat.Sender},
		{Key: "Recipient", Value: chat.Recipent},
		{Key: "MsgID", Value: msgId.InsertedID},
	})
	if err != nil {
		fmt.Println("error inserting record", err)
	}

	return nil
}
