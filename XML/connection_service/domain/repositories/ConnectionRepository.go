package repositories

import (
	"connection/module/domain/dto"
	connectionModel "connection/module/domain/model"
	"context"
)

type ConnectionRepository interface {
	CreateConnection(connection *connectionModel.Connection, ctx context.Context) (*dto.ConnectionResponse, error)
	AcceptConnection(connection *connectionModel.Connection, ctx context.Context) (*dto.ConnectionResponse, error)
	GetAllConnectionForUser(userUid string, ctx context.Context) (userNodes []*connectionModel.User, error1 error)
	GetAllConnectionRequestsForUser(userUid string, ctx context.Context) (userNodes []*connectionModel.User, error1 error)
	ConnectionStatusForUsers(senderId string, receiverId string, ctx context.Context) (*dto.ConnectionResponse, error)
	BlockUser(con *connectionModel.Connection, ctx context.Context) (*dto.ConnectionResponse, error)
	GetRecommendedNewConnections(id string, ctx context.Context) (userNodes []*connectionModel.User, error1 error)
	GetRecommendedJobOffers(id string, ctx context.Context) (jobNodes []*connectionModel.JobOffer, error1 error)
}
