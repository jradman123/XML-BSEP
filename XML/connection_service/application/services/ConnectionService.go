package services

import (
	"common/module/logger"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
)

type ConnectionService struct {
	connectionRepo repositories.ConnectionRepository
	logInfo        *logger.Logger
	logError       *logger.Logger
}

func NewConnectionService(connectionRepo repositories.ConnectionRepository, logInfo *logger.Logger, logError *logger.Logger) *ConnectionService {
	return &ConnectionService{connectionRepo, logInfo, logError}
}

func (s ConnectionService) CreateConnection() {
	con := &model.Connection{
		UserOneUID: "2",
		UserTwoUID: "3",
	}
	s.connectionRepo.CreateConnection(con)

}
