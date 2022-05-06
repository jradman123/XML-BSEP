package repository

import (
	"context"
	"user/module/model"

	"github.com/google/uuid"
)

// type UserRepository interface {
// 	Create(ctx context.Context, user *model.User) //(*mongo.InsertOneResult, error)
// 	Update(ctx context.Context, user *model.User) //(*mongo.UpdateResult, error)
// 	GetByID(ctx context.Context, id string)       //(*model.User, error)
// 	GetByUsername(ctx context.Context, username string) (*model.User, error)
// 	GetAllRolesByUserId(ctx context.Context, userId string) //([]model.Role, error)
// 	PhysicalDelete(ctx context.Context, userId string)      //(*mongo.DeleteResult, error)
// }

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (r UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	return &model.User{
		ID:          uuid.New(),
		Username:    "Jack",
		Password:    "abc123",
		Email:       "jack@gmail.com",
		PhoneNumber: "123123",
		FirstName:   "Jack",
		LastName:    "Sparrow",
		Gender:      model.MALE,
	}, nil
}
