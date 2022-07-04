package application

import (
	"common/module/logger"
	"message/module/domain/model"
	"message/module/domain/repositories"
)

type NotificationService struct {
	logInfo          *logger.Logger
	logError         *logger.Logger
	notificationRepo repositories.NotificationRepository
}

func NewNotificationService(logInfo *logger.Logger, logError *logger.Logger, notificationRepo repositories.NotificationRepository) *NotificationService {
	return &NotificationService{logInfo: logInfo, logError: logError, notificationRepo: notificationRepo}
}

func (service *NotificationService) Create(notification *model.Notification) error {
	return service.notificationRepo.Create(notification)
}

func (service *NotificationService) GetAllForUser(username string) ([]*model.Notification, error) {
	return service.notificationRepo.GetAllForUser(username)
}
