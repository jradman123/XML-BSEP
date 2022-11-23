package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"message/module/domain/model"
)

type NotificationRepository interface {
	Create(notification *model.Notification, ctx context.Context) (*model.Notification, error)
	GetAllForUser(username string, ctx context.Context) ([]*model.Notification, error)
	MarkAsRead(id primitive.ObjectID, ctx context.Context)
}
