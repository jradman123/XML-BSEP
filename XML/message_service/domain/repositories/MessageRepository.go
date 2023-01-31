package repositories

import (
	"github.com/google/uuid"
	"message/module/domain/model"
)

type MessageRepository interface {
	GetAllSent(SenderId uuid.UUID) ([]*model.Message, error)
	GetAllReceived(ReceiverId uuid.UUID) ([]*model.Message, error)
	SendMessage(message *model.Message) (*model.Message, error)
}
