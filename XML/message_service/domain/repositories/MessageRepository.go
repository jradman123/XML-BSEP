package repositories

import "message/module/domain/model"

type MessageRepository interface {
	GetAllSent(SenderId string) ([]*model.Message, error)
	GetAllReceived(ReceiverId string) ([]*model.Message, error)
	SendMessage(message *model.Message) error
}
