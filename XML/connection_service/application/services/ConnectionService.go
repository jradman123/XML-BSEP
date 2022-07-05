package services

import (
	"common/module/logger"
	"connection/module/domain/dto"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
	"connection/module/infrastructure/orchestrators"
)

type ConnectionService struct {
	connectionRepo repositories.ConnectionRepository
	logInfo        *logger.Logger
	logError       *logger.Logger
	orchestrator   *orchestrators.ConnectionOrchestrator
}

func NewConnectionService(connectionRepo repositories.ConnectionRepository, logInfo *logger.Logger, logError *logger.Logger, orchestrator *orchestrators.ConnectionOrchestrator) *ConnectionService {
	return &ConnectionService{connectionRepo, logInfo, logError, orchestrator}
}

func (s ConnectionService) CreateConnection(connection *model.Connection, sender string, receiver string) (*dto.ConnectionResponse, error) {
	response, err := s.connectionRepo.CreateConnection(connection)
	status, _ := s.connectionRepo.ConnectionStatusForUsers(connection.UserOneUID, connection.UserTwoUID)
	s.orchestrator.Connect(sender, receiver, status.ConnectionStatus)
	return response, err
}

func (s ConnectionService) AcceptConnection(connection *model.Connection, sender string, receiver string) (*dto.ConnectionResponse, error) {
	s.orchestrator.AcceptConnection(sender, receiver)
	return s.connectionRepo.AcceptConnection(connection)
}

func (s ConnectionService) GetAllConnectionForUser(userUid string) (userNodes []*model.User, error1 error) {
	response, err := s.connectionRepo.GetAllConnectionForUser(userUid)
	return response, err
}

func (s ConnectionService) GetAllConnectionRequestsForUser(userUid string) (userNodes []*model.User, error1 error) {
	return s.connectionRepo.GetAllConnectionRequestsForUser(userUid)
}

func (s ConnectionService) ConnectionStatusForUsers(senderId string, receiverId string) (*dto.ConnectionResponse, error) {
	return s.connectionRepo.ConnectionStatusForUsers(senderId, receiverId)
}

func (s ConnectionService) BlockUser(con *model.Connection) (*dto.ConnectionResponse, error) {
	response, err := s.connectionRepo.BlockUser(con)
	return response, err
}

func (s ConnectionService) GetRecommendedNewConnections(userId string) (userNodes []*model.User, error1 error) {
	return s.connectionRepo.GetRecommendedNewConnections(userId)
}
