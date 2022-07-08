package application

import (
	"common/module/logger"
	"fmt"
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
	userService        *UserService
}

func NewNotificationService(logInfo *logger.Logger, logError *logger.Logger, notificationRepo repositories.NotificationRepository, pusher *pusher.Client, userService *UserService) *NotificationService {
	return &NotificationService{logInfo: logInfo, logError: logError, notificationRepo: notificationRepo, notificationPusher: pusher, userService: userService}
}

func (service *NotificationService) Create(notification *model.Notification) (*model.Notification, error) {
	result, _ := service.userService.AllowedNotificationForUser(notification.NotificationTo, notification.Type)
	fmt.Println(result)
	if result {
		noti, err := service.notificationRepo.Create(notification)
		service.notificationPusher.Trigger("notifications", "notification", noti)
		return noti, err
	}
	return nil, nil
}

func (service *NotificationService) GetAllForUser(username string) ([]*model.Notification, error) {
	return service.notificationRepo.GetAllForUser(username)
}

func (service *NotificationService) MarkAsRead(id primitive.ObjectID) {
	service.notificationRepo.MarkAsRead(id)
}
