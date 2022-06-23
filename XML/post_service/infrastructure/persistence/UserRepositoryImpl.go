package persistence

import (
	"go.mongodb.org/mongo-driver/mongo"
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
