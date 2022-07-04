package application

import (
	"common/module/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (service *NotificationService) Create(notification *model.Notification) (*model.Notification, error) {
	return service.notificationRepo.Create(notification)
}

func (service *NotificationService) GetAllForUser(username string) ([]*model.Notification, error) {
	return service.notificationRepo.GetAllForUser(username)
}

func (service *NotificationService) MarkAsRead(id primitive.ObjectID) {
	service.notificationRepo.MarkAsRead(id)
}
