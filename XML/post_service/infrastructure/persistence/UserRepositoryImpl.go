package persistence

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	tracer "monitoring/module"
	"post/module/domain/model"
	"post/module/domain/repositories"
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
func (u UserRepositoryImpl) Get(id primitive.ObjectID) (user *model.User, err error) {
	filter := bson.M{"_id": id}
	return u.filterOne(filter)
}

func (u UserRepositoryImpl) GetByUserId(id uuid.UUID) (user []*model.User, err error) {
	filter := bson.M{"user_id": id}
	return u.filter(filter, context.TODO())
}

func (u UserRepositoryImpl) GetByUsername(username string, ctx context.Context) (user []*model.User, err error) {
	span := tracer.StartSpanFromContext(ctx, "getByUsernameRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	filter := bson.M{"username": username}
	return u.filter(filter, ctx)
}

func (u UserRepositoryImpl) CreateUser(user *model.User) (*model.User, error) {
	result, err := u.users.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (u UserRepositoryImpl) UpdateUser(user *model.User) (*model.User, error) {

	_, err := u.users.UpdateOne(context.TODO(), bson.M{"user_id": user.UserId}, bson.D{
		{"$set", bson.D{{"username", user.Username}}},
		{"$set", bson.D{{"name", user.FirstName}}},
		{"$set", bson.D{{"last_name", user.LastName}}},
		{"$set", bson.D{{"email", user.Email}}},
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

func (u UserRepositoryImpl) ActivateUserAccount(userId uuid.UUID) (err error) {
	_, err = u.users.UpdateOne(context.TODO(), bson.M{"user_id": userId}, bson.D{
		{"$set",
			bson.D{
				{"activated", true},
			}},
	})
	return err
}

func (u UserRepositoryImpl) filterOne(filter bson.M) (user *model.User, err error) {
	result := u.users.FindOne(context.TODO(), filter)
	err = result.Decode(&user)
	return
}

func (u UserRepositoryImpl) filter(filter interface{}, ctx context.Context) ([]*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "filter")
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
