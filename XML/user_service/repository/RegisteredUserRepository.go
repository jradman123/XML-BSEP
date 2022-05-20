package repository

import (
	"fmt"
	"user/module/model"

	"gorm.io/gorm"
)

type RegisteredUserRepository struct {
	DB *gorm.DB
}

func (r *RegisteredUserRepository) CreateRegisteredUser(user *model.User) (string, error) {
	result := r.DB.Create(&user)
	fmt.Print(result)
	return string(user.Email), nil
}
