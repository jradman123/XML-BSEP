package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"message/module/domain/model"
)

type NotificationRepository interface {
	Create(notification *model.Notification) error
	GetAllForUser(username string) ([]*model.Notification, error)
	MarkAsRead(id primitive.ObjectID)
}
