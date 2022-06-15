package services

import (
	"common/module/logger"
	"context"
	"gateway/module/domain/model"
	"gateway/module/domain/repositories"
)

type UserService struct {
	logInfo        *logger.Logger
	logError       *logger.Logger
	userRepository repositories.UserRepository
}

func NewUserService(logInfo *logger.Logger, logError *logger.Logger, repository repositories.UserRepository) *UserService {
	return &UserService{logInfo, logError, repository}
}
func (u UserService) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := u.userRepository.GetByUsername(ctx, username)

	if err != nil {
		u.logError.Logger.Errorf("Invalid username")
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
