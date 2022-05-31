package services

import (
	"context"
	"gateway/module/domain/model"
	"gateway/module/domain/repositories"
	"log"
)

type UserService struct {
	l              *log.Logger
	userRepository repositories.UserRepository
}

func NewUserService(l *log.Logger, repository repositories.UserRepository) *UserService {
	return &UserService{l, repository}
}
func (u UserService) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := u.userRepository.GetByUsername(ctx, username)

	if err != nil {
		u.l.Println("Invalid username")
		return nil, err
	}
	return user, nil
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

func (u UserService) GetUserSalt(username string) (string, error) {

	salt, err := u.userRepository.GetUserSalt(username)

	if err != nil {
		return "", err
	}
	return salt, nil
}
