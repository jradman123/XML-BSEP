package application

import (
	"common/module/logger"
	"context"
	"github.com/google/uuid"
	"message/module/domain/model"
	"message/module/domain/repositories"
	tracer "monitoring/module"
)

type MessageService struct {
	repository repositories.MessageRepository
	logInfo    *logger.Logger
	logError   *logger.Logger
}

func NewMessageService(repository repositories.MessageRepository, logInfo *logger.Logger, logError *logger.Logger) *MessageService {
	return &MessageService{repository: repository, logInfo: logInfo, logError: logError}
}
func (service *MessageService) GetAllSent(SenderId uuid.UUID, ctx context.Context) ([]*model.Message, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAllSentService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.GetAllSent(SenderId, ctx)
}

func (service *MessageService) GetAllReceived(ReceiverId uuid.UUID, ctx context.Context) ([]*model.Message, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAllReceivedService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.GetAllReceived(ReceiverId, ctx)
}

func (service *MessageService) SendMessage(message *model.Message, ctx context.Context) (*model.Message, error) {
	span := tracer.StartSpanFromContext(ctx, "SendMessageService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	return service.repository.SendMessage(message, ctx)
}

func (service *MessageService) UpdateUserMessages(user *model.User, ctx context.Context) error {
	return nil
}
