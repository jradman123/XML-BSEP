package services

import (
	"common/module/logger"
	"context"
	"errors"
	"gateway/module/domain/model"
	"gateway/module/domain/repositories"
	"log"
	tracer "monitoring/module"
)

type UserService struct {
	l              *log.Logger
	logInfo        *logger.Logger
	logError       *logger.Logger
	userRepository repositories.UserRepository
}

func NewUserService(l *log.Logger, logInfo *logger.Logger, logError *logger.Logger, repository repositories.UserRepository) *UserService {
	return &UserService{l, logInfo, logError, repository}
}
func (u UserService) GetByUsername(username string, ctx context.Context) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "GetByUsername-Service")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	span1 := tracer.StartSpanFromContext(ctx, "ReadUserByUsername")
	user, err := u.userRepository.GetByUsername(username)
	span1.Finish()

	if err != nil {
		tracer.LogError(span1, errors.New(err.Error()))
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

func (u UserService) GetUserRole(username string, ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "GetUserRole-Service")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	span1 := tracer.StartSpanFromContext(ctx, "ReadUserRoleForUser")
	role, err := u.userRepository.GetUserRole(username)
	span1.Finish()

	if err != nil {
		tracer.LogError(span1, errors.New(err.Error()))
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
