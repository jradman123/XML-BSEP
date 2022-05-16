package service

import (
	"time"
	"user/module/model"
	"user/module/repository"

	"github.com/google/uuid"
)

type RegisteredUserService struct {
	Repo *repository.RegisteredUserRepository
}

func (service *RegisteredUserService) CreateRegisteredUser(username string, password string, email string, phone string, firstName string, lastName string, gender model.Gender, role model.UserType, salt string, dateOfBirth time.Time) (string, error) {
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
		IsConfirmed: false,
		Salt:        salt,
		DateOfBirth: dateOfBirth,
	}
	mail, err := service.Repo.CreateRegisteredUser(&user)
	if err != nil {
		return mail, err
	}
	return mail, nil
}
