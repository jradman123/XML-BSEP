package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
	"time"
	"user/module/domain/model"
	"user/module/domain/repositories"
)

type UserService struct {
	l              *log.Logger
	userRepository repositories.UserRepository
}

func NewUserService(l *log.Logger, repository repositories.UserRepository) *UserService {
	return &UserService{l, repository}
}
func (u UserService) GetUsers() ([]model.User, error) {

	users, err := u.userRepository.GetUsers()
	if err != nil {
		return nil, errors.New("Cant get users")
	}

	return users, err

}
func (u UserService) GetByUsername(ctx context.Context, username string) (*model.User, error) {

	user, err := u.userRepository.GetByUsername(ctx, username)

	if err != nil {
		u.l.Println("Invalid username")
		return nil, err
	}

	return user, nil
}

//GetUserSalt

func (u UserService) GetUserSalt(username string) (string, error) {

	salt, err := u.userRepository.GetUserSalt(username)

	if err != nil {
		return "", err
	}
	return salt, nil
}
func (u UserService) UserExists(username string) error {

	err := u.userRepository.UserExists(username)

	if err != nil {
		return err
	}
	return nil
}

func (u UserService) GetUserRole(username string) (string, error) {

	role, err := u.userRepository.GetUserRole(username)

	if err != nil {
		return "", err
	}
	return role, nil
}
func (u UserService) CreateRegisteredUser(username string, password string, email string, phone string, firstName string, lastName string, gender model.Gender, role model.UserType, salt string, dateOfBirth time.Time) (string, error) {
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
	mail, err := u.userRepository.CreateRegisteredUser(&user)
	if err != nil {
		return mail, err
	}
	return mail, nil
}
