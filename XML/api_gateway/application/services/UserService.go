package services

import (
	"common/module/logger"
	"gateway/module/domain/model"
	"gateway/module/domain/repositories"
	"log"
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
func (u UserService) GetByUsername(username string) (*model.User, error) {

	user, err := u.userRepository.GetByUsername(username)

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
