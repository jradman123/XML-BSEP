package services

import (
	"common/module/logger"
	"connection/module/domain/dto"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
	"fmt"
)

type ConnectionService struct {
	connectionRepo repositories.ConnectionRepository
	logInfo        *logger.Logger
	logError       *logger.Logger
}

func NewConnectionService(connectionRepo repositories.ConnectionRepository, logInfo *logger.Logger, logError *logger.Logger) *ConnectionService {
	return &ConnectionService{connectionRepo, logInfo, logError}
}

func (s ConnectionService) CreateConnection(connection *model.Connection) (*dto.ConnectionResponse, error) {
	//con := &model.Connection{
	//	UserOneUID: "1",
	//	UserTwoUID: "2",
	//}
	response, err := s.connectionRepo.CreateConnection(connection)
	return response, err
}

func (s ConnectionService) AcceptConnection(connection *model.Connection) (*dto.ConnectionResponse, error) {
	return s.connectionRepo.AcceptConnection(connection)
}

func (s ConnectionService) GetAllConnectionForUser(userUid string) (userNodes []*model.User, error1 error) {
	response, err := s.connectionRepo.GetAllConnectionForUser(userUid)
	fmt.Println("dobila sam niz konekcija ove duzine i prvi username je  username" + response[0].Username)
	return response, err
}

func (s ConnectionService) GetAllConnectionRequestsForUser(userUid string) (userNodes []*model.User, error1 error) {
	return s.connectionRepo.GetAllConnectionRequestsForUser(userUid)
}
