package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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

func (repo NotificationRepositoryImpl) Create(notification *model.Notification) (*model.Notification, error) {
	result, err := repo.db.InsertOne(context.TODO(), notification)
	if err != nil {
		return nil, err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)
	createdNotification := repo.db.FindOne(context.TODO(), bson.M{"_id": notification.Id})
	var retVal model.Notification
	createdNotification.Decode(&retVal)
	return &retVal, nil
}

func (repo NotificationRepositoryImpl) GetAllForUser(username string) ([]*model.Notification, error) {
	filter := bson.M{"notification_to": username}
	return repo.filter(filter)
}

func (repo NotificationRepositoryImpl) filter(filter interface{}) ([]*model.Notification, error) {
	cursor, err := repo.db.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeNoti(cursor)
}

func decodeNoti(cursor *mongo.Cursor) (notifications []*model.Notification, err error) {
	for cursor.Next(context.TODO()) {
		var notification model.Notification
		err = cursor.Decode(&notification)
		if err != nil {
			return
		}
		notifications = append(notifications, &notification)
	}
	err = cursor.Err()
	return
}
func decodeNotification(cursor *mongo.Cursor) (notifications *model.Notification, err error) {
	var notification model.Notification
	err = cursor.Decode(&notification)
	if err != nil {
		return nil, err
	}

	err = cursor.Err()
	return &notification, nil
}

func (repo NotificationRepositoryImpl) MarkAsRead(id primitive.ObjectID) {
	repo.db.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		{"$set", bson.D{{"read", true}}},
	})
}
