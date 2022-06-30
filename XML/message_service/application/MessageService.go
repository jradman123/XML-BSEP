package application

import (
	"common/module/logger"
	"message/module/domain/model"
	"message/module/domain/repositories"
)

type MessageService struct {
	repository repositories.MessageRepository
	logInfo    *logger.Logger
	logError   *logger.Logger
}

func NewMessageService(repository repositories.MessageRepository, logInfo *logger.Logger, logError *logger.Logger) *MessageService {
	return &MessageService{repository: repository, logInfo: logInfo, logError: logError}
}
func (service *MessageService) GetAllSent(SenderId string) ([]*model.Message, error) {
	return service.repository.GetAllSent(SenderId)
}

func (service *MessageService) GetAllReceived(ReceiverId string) ([]*model.Message, error) {
	return service.repository.GetAllReceived(ReceiverId)
}

func (service *MessageService) SendMessage(message *model.Message) error {
	return service.repository.SendMessage(message)
}

func (service *MessageService) UpdateUserMessages(user *model.User) error {
	return nil
}
