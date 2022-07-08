package persistence

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"message/module/domain/model"
	"message/module/domain/repositories"
)

//NewMessageRepositoryImpl

const (
	DATABASE           = "messages_service"
	CollectionMessages = "messagesData"
)

type MessageRepositoryImpl struct {
	messages *mongo.Collection
}

func NewMessageRepositoryImpl(client *mongo.Client) repositories.MessageRepository {
	messages := client.Database(DATABASE).Collection(CollectionMessages)
	return &MessageRepositoryImpl{
		messages: messages,
	}
}

func (m MessageRepositoryImpl) SendMessage(message *model.Message) (*model.Message, error) {
	result, err := m.messages.InsertOne(context.TODO(), message)
	if err != nil {
		return nil, err
	}
	message.Id = result.InsertedID.(primitive.ObjectID)

	return message, nil
}
func (m MessageRepositoryImpl) GetAllSent(SenderId uuid.UUID) (messages []*model.Message, err error) {
	filter := bson.M{"sender_id": SenderId}
	return m.filter(filter)
}

func (m MessageRepositoryImpl) GetAllReceived(ReceiverId uuid.UUID) (messages []*model.Message, err error) {
	filter := bson.M{"receiver_id": ReceiverId}
	return m.filter(filter)
}

func (m MessageRepositoryImpl) filter(filter interface{}) ([]*model.Message, error) {
	cursor, err := m.messages.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}

	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (messages []*model.Message, err error) {
	for cursor.Next(context.TODO()) {
		var message model.Message
		err = cursor.Decode(&message)
		if err != nil {
			return
		}
		messages = append(messages, &message)
	}

	err = cursor.Err()
	return
}
