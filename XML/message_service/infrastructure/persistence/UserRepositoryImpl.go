package persistence

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"message/module/domain/model"
	"message/module/domain/repositories"
	tracer "monitoring/module"
)

const (
	CollectionUser = "usersData"
)

type UserRepositoryImpl struct {
	users *mongo.Collection
}

func NewUserRepositoryImpl(client *mongo.Client) repositories.UserRepository {
	users := client.Database(DATABASE).Collection(CollectionUser)
	return &UserRepositoryImpl{
		users: users,
	}
}

func (u UserRepositoryImpl) CreateUser(user *model.User, ctx context.Context) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateUserRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	result, err := u.users.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (u UserRepositoryImpl) UpdateUser(requestUser *model.User, ctx context.Context) (user *model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "UpdateUserRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	_, err = u.users.UpdateOne(ctx, bson.M{"user_id": requestUser.UserId}, bson.D{
		{"$set", bson.D{{"email", requestUser.Email}}},
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepositoryImpl) DeleteUser(userId uuid.UUID, ctx context.Context) (err error) {
	span := tracer.StartSpanFromContext(ctx, "DeleteUserRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	_, err = u.users.DeleteOne(ctx,
		bson.M{"user_id": userId})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (u UserRepositoryImpl) GetByUsername(username string, ctx context.Context) (user []*model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "GetByUsernameRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.M{"username": username}
	return u.filter(filter, ctx)
}

func (repo UserRepositoryImpl) GetSettingsForUser(username string, ctx context.Context) (*model.NotificationSettings, error) {
	span := tracer.StartSpanFromContext(ctx, "GetSettingsForUserRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.M{"username": username}
	var user model.User
	found := repo.users.FindOne(ctx, filter)
	found.Decode(&user)
	return &user.Settings, nil
}

func (u UserRepositoryImpl) ChangeSettingsForUser(username string, newSettings *model.NotificationSettings, ctx context.Context) (*model.NotificationSettings, error) {
	span := tracer.StartSpanFromContext(ctx, "ChangeSettingsForUserRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	_, err := u.users.UpdateOne(ctx, bson.M{"username": username}, bson.D{
		{"$set", bson.D{{"settings", newSettings}}},
	})
	if err != nil {
		return nil, err
	}
	return newSettings, nil
}

func (u UserRepositoryImpl) GetById(userId uuid.UUID, ctx context.Context) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "GetByIdRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.M{"user_id": userId}
	return u.filter(filter, ctx)
}

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

func (u UserRepositoryImpl) filterOne(filter bson.M) (user *model.User, err error) {
	result := u.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func (u UserRepositoryImpl) filter(filter interface{}, ctx context.Context) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "filterUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	cursor, err := u.users.Find(ctx, filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	if err != nil {
		return nil, err
	}

	return decodeUser(cursor, ctx)
}

func decodeUser(cursor *mongo.Cursor, ctx context.Context) (users []*model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "decodeUser")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	for cursor.Next(ctx) {
		var user model.User
		err = cursor.Decode(&user)
		if err != nil {
			return
		}
		users = append(users, &user)
	}
	err = cursor.Err()
	return
}
