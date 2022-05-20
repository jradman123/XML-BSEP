package repository

import (
	"context"
	"errors"
	"fmt"
	"user/module/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (r UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	r.db.Select("*").Find(&users)
	return users, nil

}
func (r UserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	if r.db.First(&user, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("User not found")

	}
	return user, nil
}

func (r UserRepository) CreateRegisteredUser(ctx context.Context, user *model.User) (string, error) {
	result := r.db.Create(&user)
	fmt.Print(result)
	return string(user.Email), nil
}

func (r UserRepository) UserExists(username string) error {
	if r.db.First(&model.User{}, "username = ?", username).RowsAffected == 0 {
		return errors.New("User does not exist")

	}
	return nil
}

func (r UserRepository) GetUserSalt(username string) (string, error) {
	var result string = ""
	r.db.Table("users").Select("salt").Where("username = ?", username).Scan(&result)
	if result == "" {
		return "", errors.New("User salt not found!")
	}
	return result, nil
}

func (r UserRepository) GetUserRole(username string) (string, error) {
	var result int
	r.db.Table("users").Select("role").Where("username = ?", username).Scan(&result)

	if result == 1 {
		return "user", nil
	}
	if result == 2 {
		return "admin", nil
	}
	return "", errors.New("User role not found for username" + username)
}
