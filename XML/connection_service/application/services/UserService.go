package services

import (
	"common/module/logger"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
	logInfo  *logger.Logger
	logError *logger.Logger
}

func NewUserService(userRepo repositories.UserRepository, logInfo *logger.Logger, logError *logger.Logger) *UserService {
	return &UserService{userRepo, logInfo, logError}
}

func (s UserService) CreateUser() {
	user := &model.User{UserUID: "2", Status: model.Public}
	user1 := &model.User{UserUID: "3", Status: model.Public}
	s.userRepo.Register(user)
	s.userRepo.Register(user1)
}
