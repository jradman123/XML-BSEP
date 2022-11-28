package persistence

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"message/module/domain/model"
	"message/module/domain/repositories"
	tracer "monitoring/module"
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

func (m MessageRepositoryImpl) SendMessage(message *model.Message, ctx context.Context) (*model.Message, error) {
	span := tracer.StartSpanFromContext(ctx, "SendMessageRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	result, err := m.messages.InsertOne(ctx, message)
	if err != nil {
		return nil, err
	}
	message.Id = result.InsertedID.(primitive.ObjectID)

	return message, nil
}
func (m MessageRepositoryImpl) GetAllSent(SenderId uuid.UUID, ctx context.Context) (messages []*model.Message, err error) {
	span := tracer.StartSpanFromContext(ctx, "GetAllSentRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.M{"sender_id": SenderId}
	return m.filter(filter, ctx)
}

func (m MessageRepositoryImpl) GetAllReceived(ReceiverId uuid.UUID, ctx context.Context) (messages []*model.Message, err error) {
	span := tracer.StartSpanFromContext(ctx, "GetAllReceivedRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.M{"receiver_id": ReceiverId}
	return m.filter(filter, ctx)
}

func (m MessageRepositoryImpl) filter(filter interface{}, ctx context.Context) ([]*model.Message, error) {
	span := tracer.StartSpanFromContext(ctx, "filterMessage")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	cursor, err := m.messages.Find(ctx, filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	if err != nil {
		return nil, err
	}

	return decode(cursor, ctx)
}

func decode(cursor *mongo.Cursor, ctx context.Context) (messages []*model.Message, err error) {
	span := tracer.StartSpanFromContext(ctx, "decode")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	for cursor.Next(ctx) {
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
