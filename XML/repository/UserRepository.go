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

var userList = []*model.User{
	{
		ID:          uuid.MustParse("cfa067e6-614c-4ca7-9d0b-012fcb01f9fa"),
		Username:    "Jack",
		Password:    "abc123",
		Email:       "jack@gmail.com",
		PhoneNumber: "123123",
		FirstName:   "Jack",
		LastName:    "Sparrow",
		Gender:      model.MALE,
	},
	{
		ID:          uuid.MustParse("8d4f1e1a-9897-4226-b8f4-1b2aef73457c"),
		Username:    "Tim",
		Password:    "abc123",
		Email:       "mina@gmail.com",
		PhoneNumber: "123123",
		FirstName:   "Tim",
		LastName:    "Burton",
		Gender:      model.MALE,
	},
}

func (r UserRepository) GetUsers() ([]*model.User, error) {

	return userList, nil

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

func (r UserRepository) UserExists(username string) error {
	return nil
}
func (r UserRepository) GetUserRole(username string) (string, error) {
	return "admin", nil
}
