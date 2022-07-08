package application

import (
	"common/module/logger"
	"github.com/google/uuid"
	"message/module/domain/model"
	"message/module/domain/repositories"
)

type UserService struct {
	repository repositories.UserRepository
	logInfo    *logger.Logger
	logError   *logger.Logger
}

func NewUserService(repository repositories.UserRepository, logInfo *logger.Logger, logError *logger.Logger) *UserService {
	return &UserService{repository: repository, logInfo: logInfo, logError: logError}
}

func (s UserService) CreateUser(requestUser *model.User) (user *model.User, err error) {
	user, err = s.repository.CreateUser(requestUser)
	return user, err
}

func (s UserService) UpdateUser(requestUser *model.User) (user *model.User, err error) {
	user, err = s.repository.UpdateUser(requestUser)
	return user, err
}

func (s UserService) DeleteUser(userId uuid.UUID) (err error) {
	err = s.repository.DeleteUser(userId)
	return err
}
func (s UserService) GetByUsername(username string) (user []*model.User, err error) {
	user, err = s.repository.GetByUsername(username)
	return user, err
}

func (s UserService) GetById(userId uuid.UUID) (user []*model.User, err error) {
	user, err = s.repository.GetById(userId)
	return user, err
}
