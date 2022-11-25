package services

import (
	"common/module/logger"
	"context"
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
func (u UserService) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "GetByUsernameService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err := u.userRepository.GetByUsername(ctx, username)

	if err != nil {
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
	span := tracer.StartSpanFromContext(ctx, "GetUserRoleService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	role, err := u.userRepository.GetUserRole(username, ctx)

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
