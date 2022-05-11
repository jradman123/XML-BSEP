package repository

import (
	"context"
	"user/module/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
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

func (r UserRepository) CreateUser(ctx context.Context, username string, password string, email string, phone string, firstName string, lastName string, gender model.Gender, role string) string {

	user := model.User{
		ID:          uuid.New(),
		Username:    username,
		Password:    password,
		Email:       email,
		PhoneNumber: phone,
		FirstName:   firstName,
		LastName:    lastName,
		Gender:      gender,
		Role:        role,
	}
	r.db.Create(&user)

	return string(user.Email)
}
