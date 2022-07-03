package application

import (
	"common/module/logger"
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
