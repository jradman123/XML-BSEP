package handlers

import (
	common "common/module"
	"common/module/logger"
	pb "common/module/proto/connection_service"
	"connection/module/application/services"
	"connection/module/domain/model"
	"context"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"strings"
)

type ConnectionHandler struct {
	connectionService *services.ConnectionService
	userService       *services.UserService
	logInfo           *logger.Logger
	logError          *logger.Logger
}

const (
	xssError        = "ERR:XSS"
	validationError = "ERR:BAD VALIDATION: POSIBLE INJECTION"
	getUsersError   = "ERR:GET USERS"
)

func (c ConnectionHandler) MustEmbedUnimplementedConnectionServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewConnectionHandler(connectionService *services.ConnectionService, userSer *services.UserService, logInfo *logger.Logger, logError *logger.Logger) *ConnectionHandler {
	return &ConnectionHandler{connectionService, userSer, logInfo, logError}
}

func (c ConnectionHandler) GetAll(ctx context.Context, request *pb.EmptyRequest) (*pb.EmptyRequest, error) {

	return nil, nil
}

func (c ConnectionHandler) GetSomething(ctx context.Context, request *pb.EmptyRequest) (*pb.EmptyRequest, error) {
	fmt.Println("usao u get something")
	user := model.User{
		UserUID:   "1",
		Username:  "username1",
		FirstName: "name1",
		LastName:  "lastname1",
		Status:    model.Public,
	}
	user1 := model.User{
		UserUID:   "2",
		Username:  "username2",
		FirstName: "name2",
		LastName:  "lastname2",
		Status:    model.Public,
	}
	user2 := model.User{
		UserUID:   "3",
		Username:  "username3",
		FirstName: "name3",
		LastName:  "lastname3",
		Status:    model.Private,
	}
	c.userService.CreateUser(user)
	c.userService.CreateUser(user1)
	c.userService.CreateUser(user2)
	//c.connectionService.CreateConnection()
	return &pb.EmptyRequest{}, nil
}

func (c ConnectionHandler) GetConnections(ctx context.Context, request *pb.GetRequest) (*pb.Users, error) {
	fmt.Println("GetConnections handler")
	policy := bluemonday.UGCPolicy()
	request.Uid = strings.TrimSpace(policy.Sanitize(request.Uid))

	p1 := common.BadId(request.Uid)
	//userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	if request.Uid == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Uid,
		}).Errorf(xssError)
	} else if p1 {
		c.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Uid,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"userId": request.Uid,
		}).Infof("INFO:Handling Get connections for user")
	}
	users, err := c.connectionService.GetAllConnectionForUser(request.Uid)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Uid,
		}).Errorf(getUsersError)
		return nil, err
	}
	response := &pb.Users{
		Users: []*pb.UserNode{},
	}

	for _, user := range users {
		current := pb.UserNode{UserUID: user.UserUID, Status: string(user.Status), Username: user.Username, FirstName: user.FirstName, LastName: user.LastName}
		response.Users = append(response.Users, &current)
	}

	return response, nil
}

func (c ConnectionHandler) GetConnectionRequests(ctx context.Context, request *pb.GetRequest) (*pb.Users, error) {
	fmt.Println("GetConnections handler")
	policy := bluemonday.UGCPolicy()
	request.Uid = strings.TrimSpace(policy.Sanitize(request.Uid))

	p1 := common.BadId(request.Uid)
	//userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	if request.Uid == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Uid,
		}).Errorf(xssError)
	} else if p1 {
		c.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Uid,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"userId": request.Uid,
		}).Infof("INFO:Handling Get connections for user")
	}
	users, err := c.connectionService.GetAllConnectionRequestsForUser(request.Uid)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"userId": request.Uid,
		}).Errorf(getUsersError)
		return nil, err
	}
	response := &pb.Users{
		Users: []*pb.UserNode{},
	}

	for _, user := range users {
		current := pb.UserNode{UserUID: user.UserUID, Status: string(user.Status), Username: user.Username, FirstName: user.FirstName, LastName: user.LastName}
		response.Users = append(response.Users, &current)
	}

	return response, nil
}

func (c ConnectionHandler) CreateConnection(ctx context.Context, connection *pb.NewConnection) (*pb.ConnectionResponse, error) {
	fmt.Println("Create connection handler")
	policy := bluemonday.UGCPolicy()
	connection.Connection.UserSender = strings.TrimSpace(policy.Sanitize(connection.Connection.UserSender))
	connection.Connection.UserReceiver = strings.TrimSpace(policy.Sanitize(connection.Connection.UserReceiver))

	p1 := common.BadId(connection.Connection.UserSender)
	p2 := common.BadId(connection.Connection.UserReceiver)
	//userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	if connection.Connection.UserSender == "" || connection.Connection.UserReceiver == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderId": connection.Connection.UserSender,
		}).Errorf(xssError)
	} else if p1 || p2 {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderId": connection.Connection.UserSender,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"userSenderId": connection.Connection.UserSender,
		}).Infof("INFO:Handling Create connection")
	}
	con := &model.Connection{
		UserOneUID: connection.Connection.UserSender,
		UserTwoUID: connection.Connection.UserReceiver,
	}
	conResult, err := c.connectionService.CreateConnection(con)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderId": connection.Connection.UserSender,
		}).Errorf("ERR:CREATE CONNECTION")
		return nil, err
	}
	return &pb.ConnectionResponse{UserReceiver: conResult.UserOneUID, UserSender: conResult.UserTwoUID, ConnectionStatus: conResult.ConnectionStatus}, nil
}

func (c ConnectionHandler) AcceptConnection(ctx context.Context, connection *pb.NewConnection) (*pb.ConnectionResponse, error) {
	fmt.Println("Accept connection handler")
	policy := bluemonday.UGCPolicy()
	connection.Connection.UserSender = strings.TrimSpace(policy.Sanitize(connection.Connection.UserSender))
	connection.Connection.UserReceiver = strings.TrimSpace(policy.Sanitize(connection.Connection.UserReceiver))

	p1 := common.BadId(connection.Connection.UserSender)
	p2 := common.BadId(connection.Connection.UserReceiver)
	//userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	if connection.Connection.UserSender == "" || connection.Connection.UserReceiver == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderId": connection.Connection.UserSender,
		}).Errorf(xssError)
	} else if p1 || p2 {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderId": connection.Connection.UserSender,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"userSenderId": connection.Connection.UserSender,
		}).Infof("INFO:Handling Create connection")
	}
	con := &model.Connection{
		UserOneUID: connection.Connection.UserSender,
		UserTwoUID: connection.Connection.UserReceiver,
	}
	conResult, err := c.connectionService.AcceptConnection(con)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderId": connection.Connection.UserSender,
		}).Errorf("ERR:CREATE CONNECTION")
		return nil, err
	}
	return &pb.ConnectionResponse{UserReceiver: conResult.UserOneUID, UserSender: conResult.UserTwoUID, ConnectionStatus: conResult.ConnectionStatus}, nil

}
