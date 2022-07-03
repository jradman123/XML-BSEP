package persistence

import (
	"go.mongodb.org/mongo-driver/mongo"
	"message/module/domain/repositories"
)

const (
	CollectionNotification = "notificationData"
)

type NotificationRepositoryImpl struct {
	db *mongo.Collection
}

func NewNotificationRepositoryImpl(client *mongo.Client) repositories.NotificationRepository {
	db := client.Database(DATABASE).Collection(CollectionNotification)
	return &NotificationRepositoryImpl{
		db: db,
	}
}
