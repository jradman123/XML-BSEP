package service

import (
	"context"
	"errors"
	"log"
	"user/module/model"
	"user/module/repository"
)

type UserService struct {
	l              *log.Logger
	userRepository repository.UserRepository
}

//sth like a constructor
func NewUserService(l *log.Logger, repository repository.UserRepository) *UserService {
	return &UserService{l, repository}
}

func (u UserService) GetByUsername(ctx context.Context, username string) (*model.User, error) {

	user, err := u.userRepository.GetByUsername(ctx, username)

	if err != nil {
		u.l.Println("Invalid username")
		return nil, errors.New("Invalid username")
	}

	return user, err
}
