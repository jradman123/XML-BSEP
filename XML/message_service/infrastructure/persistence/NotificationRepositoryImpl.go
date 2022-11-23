package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"message/module/domain/model"
	"message/module/domain/repositories"
	tracer "monitoring/module"
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

func (repo NotificationRepositoryImpl) Create(notification *model.Notification, ctx context.Context) (*model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "createNotificationRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	result, err := repo.db.InsertOne(ctx, notification)
	if err != nil {
		return nil, err
	}
	notification.Id = result.InsertedID.(primitive.ObjectID)
	createdNotification := repo.db.FindOne(ctx, bson.M{"_id": notification.Id})
	var retVal model.Notification
	createdNotification.Decode(&retVal)
	return &retVal, nil
}

func (repo NotificationRepositoryImpl) GetAllForUser(username string, ctx context.Context) ([]*model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "getAllForUserRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.M{"notification_to": username}
	return repo.filter(filter, ctx)
}

func (repo NotificationRepositoryImpl) filter(filter interface{}, ctx context.Context) ([]*model.Notification, error) {
	span := tracer.StartSpanFromContext(ctx, "filterNotification")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	cursor, err := repo.db.Find(ctx, filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	if err != nil {
		return nil, err
	}

	return decodeNoti(cursor, ctx)
}

func decodeNoti(cursor *mongo.Cursor, ctx context.Context) (notifications []*model.Notification, err error) {
	span := tracer.StartSpanFromContext(ctx, "decodeNoti")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	for cursor.Next(ctx) {
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

func (repo NotificationRepositoryImpl) MarkAsRead(id primitive.ObjectID, ctx context.Context) {
	span := tracer.StartSpanFromContext(ctx, "markAsReadRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	repo.db.UpdateOne(ctx, bson.M{"_id": id}, bson.D{
		{"$set", bson.D{{"read", true}}},
	})
}
