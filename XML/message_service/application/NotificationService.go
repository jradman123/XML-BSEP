package application

import (
	"common/module/logger"
	"context"
	"fmt"
	"github.com/pusher/pusher-http-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"message/module/domain/model"
	"message/module/domain/repositories"
	tracer "monitoring/module"
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

func (service *NotificationService) Create(notification *model.Notification, ctx context.Context) (*model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateNotificationService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	result, _ := service.userService.AllowedNotificationForUser(notification.NotificationTo, notification.Type, ctx)
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
