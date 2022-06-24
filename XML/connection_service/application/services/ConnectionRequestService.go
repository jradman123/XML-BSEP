package services

import (
	"common/module/logger"
	"connection/module/domain/repositories"
)

type ConnectionRequestService struct {
	connectionRequestRepo repositories.ConnectionRequestRepository
	logInfo               *logger.Logger
	logError              *logger.Logger
}

func NewConnectionRequestService(connectionRepo repositories.ConnectionRequestRepository, logInfo *logger.Logger, logError *logger.Logger) *ConnectionRequestService {
	return &ConnectionRequestService{connectionRepo, logInfo, logError}
}
