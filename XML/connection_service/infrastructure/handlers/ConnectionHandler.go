package handlers

import (
	"common/module/logger"
	pb "common/module/proto/connection_service"
	"connection/module/application/services"
	"context"
	"fmt"
)

type ConnectionHandler struct {
	connectionService        *services.ConnectionService
	connectionRequestService *services.ConnectionRequestService
	userService              *services.UserService
	logInfo                  *logger.Logger
	logError                 *logger.Logger
}

func (c ConnectionHandler) MustEmbedUnimplementedConnectionServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewConnectionHandler(connectionService *services.ConnectionService, conReqSer *services.ConnectionRequestService, userSer *services.UserService, logInfo *logger.Logger, logError *logger.Logger) *ConnectionHandler {
	return &ConnectionHandler{connectionService, conReqSer, userSer, logInfo, logError}
}

func (c ConnectionHandler) GetAll(ctx context.Context, request *pb.EmptyRequest) (*pb.EmptyRequest, error) {

	return nil, nil
}

func (c ConnectionHandler) GetSomething(ctx context.Context, request *pb.EmptyRequest) (*pb.EmptyRequest, error) {
	fmt.Println("usao u get something")
	//c.userService.CreateUser()
	c.connectionService.CreateConnection()
	return &pb.EmptyRequest{}, nil
}
