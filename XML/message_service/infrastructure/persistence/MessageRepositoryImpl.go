package persistence

import (
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

func (m MessageRepositoryImpl) GetAllSent(SenderId string) ([]*model.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepositoryImpl) GetAllReceived(ReceiverId string) ([]*model.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m MessageRepositoryImpl) SendMessage(offer *model.Message) error {
	//TODO implement me
	panic("implement me")
}
