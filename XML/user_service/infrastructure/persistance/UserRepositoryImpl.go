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

func (r UserRepositoryImpl) EditUserDetails(user *model.User) (bool, error) {
	result := r.db.Model(&user).Updates(&user)
	r.db.Model(&user).Association("Skills").Replace(user.Skills)
	r.db.Model(&user).Association("Interests").Replace(user.Interests)
	r.db.Model(&user).Association("Educations").Replace(user.Educations)
	r.db.Model(&user).Association("Experiences").Replace(user.Experiences)
	r.db.Where("user_id is null").Delete(&model.Skill{})
	r.db.Where("user_id is null").Delete(&model.Interest{})
	r.db.Where("user_id is null").Delete(&model.Education{})
	r.db.Where("user_id is null").Delete(&model.Experience{})
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
	if r.db.Preload("Skills").Preload("Interests").Preload("Educations").Preload("Experiences").First(&user, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("user not found")

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

	if result == 0 {
		return "Regular", nil
	}

	if result == 1 {
		return "Admin", nil
	}
	if result == 2 {
		return "Agent", nil
	}
	return "", errors.New("User role not found for username" + username)
}

func (r UserRepositoryImpl) ChangePassword(user *model.User, password string) error {

	result := r.db.Model(&user).Update("password", password)
	fmt.Print(result)
	return result.Error
}
