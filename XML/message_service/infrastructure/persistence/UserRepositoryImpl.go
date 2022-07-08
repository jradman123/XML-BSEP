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

func (u UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	result, err := u.users.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (u UserRepositoryImpl) UpdateUser(requestUser *model.User) (user *model.User, err error) {
	_, err = u.users.UpdateOne(context.TODO(), bson.M{"user_id": requestUser.UserId}, bson.D{
		{"$set", bson.D{{"email", requestUser.Email}}},
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepositoryImpl) DeleteUser(userId uuid.UUID) (err error) {
	_, err = u.users.DeleteOne(context.TODO(),
		bson.M{"user_id": userId})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (u UserRepositoryImpl) GetByUsername(username string) (user []*model.User, err error) {
	filter := bson.M{"username": username}
	return u.filter(filter)
}

func (repo UserRepositoryImpl) GetSettingsForUser(username string) (*model.NotificationSettings, error) {
	filter := bson.M{"username": username}
	var user model.User
	found := repo.users.FindOne(context.TODO(), filter)
	found.Decode(&user)
	return &user.Settings, nil
}

func (u UserRepositoryImpl) ChangeSettingsForUser(username string, newSettings *model.NotificationSettings) (*model.NotificationSettings, error) {

	_, err := u.users.UpdateOne(context.TODO(), bson.M{"username": username}, bson.D{
		{"$set", bson.D{{"settings", newSettings}}},
	})
	if err != nil {
		return nil, err
	}
	return newSettings, nil
}

func (u UserRepositoryImpl) GetById(userId uuid.UUID) ([]*model.User, error) {
	filter := bson.M{"user_id": userId}
	return u.filter(filter)
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

func (u UserRepositoryImpl) filter(filter interface{}) ([]*model.User, error) {
	cursor, err := u.users.Find(context.TODO(), filter)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.TODO())

	if err != nil {
		return nil, err
	}

	return decodeUser(cursor)
}

func decodeUser(cursor *mongo.Cursor) (users []*model.User, err error) {
	for cursor.Next(context.TODO()) {
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
