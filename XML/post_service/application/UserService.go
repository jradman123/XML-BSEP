package application

import (
	"common/module/logger"
	"post/module/domain/model"
	"post/module/domain/repositories"
)

type UserService struct {
	repository repositories.UserRepository
	logInfo    *logger.Logger
	logError   *logger.Logger
}

func NewUserService(repository repositories.UserRepository, logInfo *logger.Logger, logError *logger.Logger) *UserService {
	return &UserService{repository: repository, logInfo: logInfo, logError: logError}
}

func (s UserService) CreateUser(user model.User) (err error) {
	return err
}

func (s UserService) UpdateUser(user model.User) (err error) {
	return err
}

func (s UserService) DeleteUser(username string) (err error) {
	return err
}
