package repositories

import (
	"context"
	"github.com/google/uuid"
	"message/module/domain/model"
)

type MessageRepository interface {
	GetAllSent(SenderId uuid.UUID, ctx context.Context) ([]*model.Message, error)
	GetAllReceived(ReceiverId uuid.UUID, ctx context.Context) ([]*model.Message, error)
	SendMessage(message *model.Message, ctx context.Context) (*model.Message, error)
}
