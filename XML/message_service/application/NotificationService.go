package application

import (
	"common/module/logger"
	"github.com/pusher/pusher-http-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"message/module/domain/model"
	"message/module/domain/repositories"
)

type NotificationService struct {
	logInfo            *logger.Logger
	logError           *logger.Logger
	notificationRepo   repositories.NotificationRepository
	notificationPusher *pusher.Client
}

func NewNotificationService(logInfo *logger.Logger, logError *logger.Logger, notificationRepo repositories.NotificationRepository, pusher *pusher.Client) *NotificationService {
	return &NotificationService{logInfo: logInfo, logError: logError, notificationRepo: notificationRepo, notificationPusher: pusher}
}

func (service *NotificationService) Create(notification *model.Notification) (*model.Notification, error) {
	noti, err := service.notificationRepo.Create(notification)
	service.notificationPusher.Trigger("notifications", "notification", noti)
	return noti, err
}

func (service *NotificationService) GetAllForUser(username string) ([]*model.Notification, error) {
	return service.notificationRepo.GetAllForUser(username)
}

func (service *NotificationService) MarkAsRead(id primitive.ObjectID) {
	service.notificationRepo.MarkAsRead(id)
}
