package services

import (
	"common/module/logger"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
	"context"
	"fmt"
	tracer "monitoring/module"
)

type UserService struct {
	userRepo repositories.UserRepository
	logInfo  *logger.Logger
	logError *logger.Logger
}

func NewUserService(userRepo repositories.UserRepository, logInfo *logger.Logger, logError *logger.Logger) *UserService {
	return &UserService{userRepo, logInfo, logError}
}

func (s UserService) CreateUser(user model.User) error {

	_, err := s.userRepo.Register(&user)
	if err != nil {
		return err
	}
	return nil
}

func (s UserService) UpdateUser(user model.User) error {
	err := s.userRepo.UpdateUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func (s UserService) DeleteUser(user model.User) error {
	fmt.Println("TODO:delete user from database")
	return nil
}

func (s UserService) GetUserId(username string, ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "getUserIdService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	userId, err := s.userRepo.GetUserId(username, ctx)
	fmt.Println("dobila sam ovaj user id za username " + userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (s UserService) ChangeProfileStatus(user model.User) error {
	fmt.Println("saga : connection service : change profile status : unimplemented")
	err := s.userRepo.ChangeProfileStatus(&user)
	if err != nil {
		return err
	}
	return nil
	return nil
}

func (s UserService) UpdateUserProfessionalDetails(user model.User, details model.UserDetails) error {
	err := s.userRepo.UpdateUserProfessionalDetails(&user, &details)
	if err != nil {
		return err
	}
	return nil
}
