package repository

import (
	"errors"
	"fmt"
	"user/module/model"

	"gorm.io/gorm"
)

type RegisteredUserRepository struct {
	DB *gorm.DB
}

func (r *RegisteredUserRepository) ChangePassword(user *model.User, newHashedPass string) error {
	fmt.Println("molim te")
	result := r.DB.Model(&user).Update("password", newHashedPass)
	fmt.Print(result)
	return result.Error
}

func (r *RegisteredUserRepository) ActivateUserAccount(user *model.User) {
	result := r.DB.Model(&user).Update("is_confirmed", true)
	fmt.Print(result)
}

func (r *RegisteredUserRepository) CreateRegisteredUser(user *model.User) (string, error) {
	result := r.DB.Create(&user)
	fmt.Print(result)
	return string(user.Email), nil
}

func (r *RegisteredUserRepository) GetByUsername(username string) (*model.User, error) {
	user := &model.User{}
	if r.DB.First(&user, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return user, nil
}
func (r *RegisteredUserRepository) UsernameExists(username string) bool {
	err := r.DB.First(&model.User{}, "username = ?", username)
	return err.Error == nil
	// if r.DB.First(&model.User{}, "username = ?", username).RowsAffected == 0 {
	// 	return errors.New("user does not exist")
	// }
}
