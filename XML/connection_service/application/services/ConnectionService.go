package services

import (
	"common/module/logger"
	"connection/module/domain/dto"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
	"connection/module/infrastructure/orchestrators"
	"context"
	tracer "monitoring/module"
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

func (s ConnectionService) CreateConnection(connection *model.Connection, sender string, receiver string, ctx context.Context) (*dto.ConnectionResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "createConnectionService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	response, err := s.connectionRepo.CreateConnection(connection, ctx)
	status, _ := s.connectionRepo.ConnectionStatusForUsers(connection.UserOneUID, connection.UserTwoUID, ctx)
	s.orchestrator.Connect(sender, receiver, status.ConnectionStatus)
	return response, err
}

func (s ConnectionService) AcceptConnection(connection *model.Connection, sender string, receiver string, ctx context.Context) (*dto.ConnectionResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "acceptConnectionService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	s.orchestrator.AcceptConnection(sender, receiver)
	return s.connectionRepo.AcceptConnection(connection, ctx)
}

func (s ConnectionService) GetAllConnectionForUser(userUid string, ctx context.Context) (userNodes []*model.User, error1 error) {
	span := tracer.StartSpanFromContext(ctx, "getAllConnectionForUserService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	response, err := s.connectionRepo.GetAllConnectionForUser(userUid, ctx)
	return response, err
}

func (s ConnectionService) GetAllConnectionRequestsForUser(userUid string, ctx context.Context) (userNodes []*model.User, error1 error) {
	span := tracer.StartSpanFromContext(ctx, "getAllConnectionRequestsForUserService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return s.connectionRepo.GetAllConnectionRequestsForUser(userUid, ctx)
}

func (s ConnectionService) ConnectionStatusForUsers(senderId string, receiverId string, ctx context.Context) (*dto.ConnectionResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "connectionStatusForUserService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return s.connectionRepo.ConnectionStatusForUsers(senderId, receiverId, ctx)
}

func (s ConnectionService) BlockUser(con *model.Connection, ctx context.Context) (*dto.ConnectionResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "blockUserService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	response, err := s.connectionRepo.BlockUser(con, ctx)
	return response, err
}

func (s ConnectionService) GetRecommendedNewConnections(userId string, ctx context.Context) (userNodes []*model.User, error1 error) {
	span := tracer.StartSpanFromContext(ctx, "getRecommendedNewConnectionsService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return s.connectionRepo.GetRecommendedNewConnections(userId, ctx)
}

func (s ConnectionService) GetRecommendedJobOffers(userId string, ctx context.Context) (jobNodes []*model.JobOffer, error1 error) {
	span := tracer.StartSpanFromContext(ctx, "getRecommendedJobOffersService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	return s.connectionRepo.GetRecommendedJobOffers(userId, ctx)
}
