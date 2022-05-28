package persistance

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"user/module/domain/model"
	"user/module/domain/repositories"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r UserRepositoryImpl) ActivateUserAccount(user *model.User) (bool, error) {
	result := r.db.Model(&user).Update("is_confirmed", true)
	fmt.Print(result)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func NewUserRepositoryImpl(db *gorm.DB) repositories.UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r UserRepositoryImpl) GetUsers() ([]model.User, error) {
	var users []model.User
	r.db.Select("*").Find(&users)
	return users, nil
}

func (r UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{}
	if r.db.First(&user, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("User not found")

	}
	return user, nil
}

func (r UserRepositoryImpl) CreateRegisteredUser(user *model.User) (*model.User, error) {
	result := r.db.Create(&user)
	fmt.Print(result)
	regUser := &model.User{}
	r.db.First(&regUser, user.ID)
	return regUser, nil
}

func (r UserRepositoryImpl) UserExists(username string) error {
	if r.db.First(&model.User{}, "username = ?", username).RowsAffected == 0 {
		return errors.New("user does not exist")

	}
	return nil
}

func (r UserRepositoryImpl) GetUserSalt(username string) (string, error) {
	var result string = ""
	r.db.Table("users").Select("salt").Where("username = ?", username).Scan(&result)
	if result == "" {
		return "", errors.New("User salt not found!")
	}
	return result, nil
}

func (r UserRepositoryImpl) GetUserRole(username string) (string, error) {
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
