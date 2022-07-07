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

func (c ConnectionHandler) mustEmbedUnimplementedConnectionServiceServer() {
	//TODO implement me
	panic("implement me")
}

const (
	xssError        = "ERR:XSS"
	validationError = "ERR:BAD VALIDATION: POSSIBLE INJECTION"
	getUsersError   = "ERR:GET USERS"
	emptyUsers      = "NO USERS"
)

func (c ConnectionHandler) MustEmbedUnimplementedConnectionServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewConnectionHandler(connectionService *services.ConnectionService, userSer *services.UserService, logInfo *logger.Logger, logError *logger.Logger) *ConnectionHandler {
	return &ConnectionHandler{connectionService, userSer, logInfo, logError}
}

func (c ConnectionHandler) GetConnections(ctx context.Context, request *pb.GetRequest) (*pb.Users, error) {
	fmt.Println("GetConnections handler")
	policy := bluemonday.UGCPolicy()
	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))

	p1 := common.BadUsername(request.Username)
	if request.Username == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(xssError)
	} else if p1 {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Infof("INFO:Handling Get connections for user")
	}

	userId, err := c.userService.GetUserId(request.Username)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(getUsersError)
		return nil, err
	}

	users, err := c.connectionService.GetAllConnectionForUser(userId)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(getUsersError)
		return nil, err
	}
	response := &pb.Users{
		Users: []*pb.UserNode{},
	}
	if users == nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(getUsersError)
		return response, nil
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
	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))

	p1 := common.BadUsername(request.Username)
	if request.Username == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(xssError)
	} else if p1 {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Infof("INFO:Handling Get connections for user")
	}

	userId, err := c.userService.GetUserId(request.Username)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(getUsersError)
		return nil, err
	}

	users, err := c.connectionService.GetAllConnectionRequestsForUser(userId)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(getUsersError)
		return nil, err
	}
	response := &pb.Users{
		Users: []*pb.UserNode{},
	}
	if users == nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(getUsersError)
		return response, nil
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

	p1 := common.BadUsername(connection.Connection.UserSender)
	p2 := common.BadUsername(connection.Connection.UserReceiver)
	if connection.Connection.UserSender == "" || connection.Connection.UserReceiver == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Errorf(xssError)
	} else if p1 || p2 {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderId": connection.Connection.UserSender,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Infof("INFO:Handling Create connection")
	}
	userSenderId, err := c.userService.GetUserId(connection.Connection.UserSender)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": connection.Connection.UserSender,
		}).Errorf(getUsersError)
		return nil, err
	}
	userReceiverId, err := c.userService.GetUserId(connection.Connection.UserReceiver)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": connection.Connection.UserSender,
		}).Errorf(getUsersError)
		return nil, err
	}

	con := &model.Connection{
		UserOneUID: userSenderId,
		UserTwoUID: userReceiverId,
	}
	conResult, err := c.connectionService.CreateConnection(con)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
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

	p1 := common.BadUsername(connection.Connection.UserSender)
	p2 := common.BadUsername(connection.Connection.UserReceiver)
	if connection.Connection.UserSender == "" || connection.Connection.UserReceiver == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Errorf(xssError)
	} else if p1 || p2 {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Infof("INFO:Handling Create connection")
	}

	userSenderId, err := c.userService.GetUserId(connection.Connection.UserSender)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": connection.Connection.UserSender,
		}).Errorf(getUsersError)
		return nil, err
	}
	userReceiverId, err := c.userService.GetUserId(connection.Connection.UserReceiver)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": connection.Connection.UserSender,
		}).Errorf(getUsersError)
		return nil, err
	}

	con := &model.Connection{
		UserOneUID: userSenderId,
		UserTwoUID: userReceiverId,
	}
	conResult, err := c.connectionService.AcceptConnection(con)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Errorf("ERR:CREATE CONNECTION")
		return nil, err
	}
	return &pb.ConnectionResponse{UserReceiver: conResult.UserOneUID, UserSender: conResult.UserTwoUID, ConnectionStatus: conResult.ConnectionStatus}, nil

}

func (c ConnectionHandler) ConnectionStatusForUsers(ctx context.Context, connection *pb.NewConnection) (*pb.ConnectionResponse, error) {
	fmt.Println("ConnectionStatusForUsers handler")
	policy := bluemonday.UGCPolicy()
	connection.Connection.UserSender = strings.TrimSpace(policy.Sanitize(connection.Connection.UserSender))
	connection.Connection.UserReceiver = strings.TrimSpace(policy.Sanitize(connection.Connection.UserReceiver))

	p1 := common.BadUsername(connection.Connection.UserSender)
	p2 := common.BadUsername(connection.Connection.UserReceiver)
	if connection.Connection.UserSender == "" || connection.Connection.UserReceiver == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Errorf(xssError)
	} else if p1 || p2 {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Infof("INFO:Handling ConnectionStatusForUsers")
	}

	userSenderId, err := c.userService.GetUserId(connection.Connection.UserSender)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": connection.Connection.UserSender,
		}).Errorf(getUsersError)
		return nil, err
	}
	userReceiverId, err := c.userService.GetUserId(connection.Connection.UserReceiver)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": connection.Connection.UserSender,
		}).Errorf(getUsersError)
		return nil, err
	}

	conResult, err := c.connectionService.ConnectionStatusForUsers(userSenderId, userReceiverId)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Errorf("ERR:GET CONNECTION")
		return nil, err
	}
	return &pb.ConnectionResponse{UserReceiver: conResult.UserOneUID, UserSender: conResult.UserTwoUID, ConnectionStatus: conResult.ConnectionStatus}, nil

}
func (c ConnectionHandler) BlockUser(ctx context.Context, connection *pb.NewConnection) (*pb.ConnectionResponse, error) {
	fmt.Println("Block user handler")
	policy := bluemonday.UGCPolicy()
	connection.Connection.UserSender = strings.TrimSpace(policy.Sanitize(connection.Connection.UserSender))
	connection.Connection.UserReceiver = strings.TrimSpace(policy.Sanitize(connection.Connection.UserReceiver))

	p1 := common.BadUsername(connection.Connection.UserSender)
	p2 := common.BadUsername(connection.Connection.UserReceiver)
	if connection.Connection.UserSender == "" || connection.Connection.UserReceiver == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Errorf(xssError)
	} else if p1 || p2 {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Infof("INFO:Handling Create connection")
	}

	userSenderId, err := c.userService.GetUserId(connection.Connection.UserSender)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": connection.Connection.UserSender,
		}).Errorf(getUsersError)
		return nil, err
	}
	userReceiverId, err := c.userService.GetUserId(connection.Connection.UserReceiver)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": connection.Connection.UserSender,
		}).Errorf(getUsersError)
		return nil, err
	}

	con := &model.Connection{
		UserOneUID: userSenderId,
		UserTwoUID: userReceiverId,
	}
	conResult, err := c.connectionService.BlockUser(con)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"userSenderUsername": connection.Connection.UserSender,
		}).Errorf("ERR:CREATE CONNECTION")
		return nil, err
	}
	return &pb.ConnectionResponse{UserReceiver: conResult.UserOneUID, UserSender: conResult.UserTwoUID, ConnectionStatus: conResult.ConnectionStatus}, nil

}

func (c ConnectionHandler) GetRecommendedNewConnections(ctx context.Context, request *pb.GetRequest) (*pb.Users, error) {
	fmt.Println("GetRecommendedNewConnections handler")
	policy := bluemonday.UGCPolicy()
	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))

	p1 := common.BadUsername(request.Username)
	if request.Username == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(xssError)
	} else if p1 {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Infof("INFO:Handling GetRecommendedNewConnections for user")
	}

	userId, err := c.userService.GetUserId(request.Username)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(getUsersError)
		return nil, err
	}

	users, err := c.connectionService.GetRecommendedNewConnections(userId)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(getUsersError)
		return nil, err
	}
	response := &pb.Users{
		Users: []*pb.UserNode{},
	}
	if users == nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(emptyUsers)
		return response, nil
	}

	for _, user := range users {
		current := pb.UserNode{UserUID: user.UserUID, Status: string(user.Status), Username: user.Username, FirstName: user.FirstName, LastName: user.LastName}
		response.Users = append(response.Users, &current)
	}

	return response, nil
}

func (c ConnectionHandler) GetRecommendedJobOffers(ctx context.Context, request *pb.GetRequest) (*pb.Offers, error) {
	fmt.Println("GetRecommendedJobOffers handler")
	policy := bluemonday.UGCPolicy()
	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))

	p1 := common.BadUsername(request.Username)
	if request.Username == "" {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(xssError)
	} else if p1 {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(validationError)
	} else {
		c.logInfo.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Infof("INFO:Handling GetRecommendedJobOffers for user")
	}

	userId, err := c.userService.GetUserId(request.Username)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(getUsersError)
		return nil, err
	}

	offers, err := c.connectionService.GetRecommendedJobOffers(userId)
	if err != nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(getUsersError)
		return nil, err
	}
	response := &pb.Offers{
		Offers: []*pb.OfferNode{},
	}
	if offers == nil {
		c.logError.Logger.WithFields(logrus.Fields{
			"username": request.Username,
		}).Errorf(emptyUsers)
		return response, nil
	}

	for _, offer := range offers {
		current := pb.OfferNode{Id: offer.JobId.Hex(), JobDescription: offer.JobDescription, Position: offer.Position, Duration: offer.Duration, DatePosted: offer.DatePosted, Publisher: offer.Publisher, Requirements: offer.Requirements}
		response.Offers = append(response.Offers, &current)
	}

	return response, nil
}
