package handlers

import (
	"common/module/logger"
	notificationProto "common/module/proto/notification_service"
	"context"
	"github.com/pusher/pusher-http-go"
	"message/module/application"
)

type NotificationHandler struct {
	logInfo             *logger.Logger
	logError            *logger.Logger
	notificationPusher  *pusher.Client
	notificationService *application.NotificationService
}

func (n NotificationHandler) GetAllSent(ctx context.Context, request *notificationProto.GetRequest) (*notificationProto.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (n NotificationHandler) MustEmbedUnimplementedNotificationServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewNotificationHandler(logInfo *logger.Logger, logError *logger.Logger, notificationPusher *pusher.Client, notificationService *application.NotificationService) *NotificationHandler {
	return &NotificationHandler{logInfo, logError, notificationPusher, notificationService}
}
