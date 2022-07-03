package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"message/module/domain/model"
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

func (repo NotificationRepositoryImpl) Create(notification *model.Notification) error {
	result, err := repo.db.InsertOne(context.TODO(), notification)
	if err != nil {
		return err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)

	return nil
}
