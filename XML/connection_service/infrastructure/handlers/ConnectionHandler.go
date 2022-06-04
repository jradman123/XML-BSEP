package handlers

import (
	pb "common/module/proto/connection_service"
	"connection/module/application/services"
	"context"
	"log"
)

type ConnectionHandler struct {
	l                 *log.Logger
	connectionService *services.ConnectionService
}

func (c ConnectionHandler) MustEmbedUnimplementedConnectionServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewConnectionHandler(l *log.Logger, connectionService *services.ConnectionService) *ConnectionHandler {
	return &ConnectionHandler{l, connectionService}
}

func (c ConnectionHandler) GetAll(ctx context.Context, request *pb.EmptyRequest) (*pb.EmptyRequest, error) {

	return nil, nil
}
