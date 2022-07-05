package startup

import (
	"common/module/interceptor"
	"common/module/logger"
	connectionProto "common/module/proto/connection_service"
	saga "common/module/saga/messaging"
	"common/module/saga/messaging/nats"
	"connection/module/application/services"
	"connection/module/domain/repositories"
	"connection/module/infrastructure/handlers"
	"connection/module/infrastructure/persistance"
	"connection/module/startup/config"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "connection_service"
)

func (server *Server) Start() {
	logInfo := logger.InitializeLogger("connection-service", context.Background(), "Info")
	logError := logger.InitializeLogger("connection-service", context.Background(), "Error")

	fmt.Println("starting connection service server")
	neoClient := server.SetupDatabase()

	connectionRepo := server.InitConnectionRepository(neoClient, logInfo, logError)
	connectionService := server.InitConnectionService(connectionRepo, logInfo, logError)

	userRepo := server.InitUserRepository(neoClient, logInfo, logError, connectionRepo)
	userService := server.InitUserService(userRepo, logInfo, logError)

	connectionHandler := server.InitConnectionHandler(connectionService, userService, logInfo, logError)

	commandSubscriber := server.InitSubscriber(server.config.UserCommandSubject, QueueGroup)
	replyPublisher := server.InitPublisher(server.config.UserReplySubject)
	server.InitCreateUserCommandHandler(userService, replyPublisher, commandSubscriber)

	server.StartGrpcServer(connectionHandler, logError)

}

func (server *Server) SetupDatabase() *neo4j.Driver {
	client, err := GetClient(server.config.Neo4jUri, server.config.Neo4jUsername, server.config.Neo4jPassword)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) StartGrpcServer(handler *handlers.ConnectionHandler, logError *logger.Logger) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
	if err != nil {
		log.Fatalf("failed to parse public key: %v", err)
	}
	//interceptor := interceptor.NewAuthInterceptor(config.AccessibleRoles(), publicKey)
	interceptor := interceptor.NewAuthInterceptor(config.AccessibleRoles(), publicKey, logError)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))
	connectionProto.RegisterConnectionServiceServer(grpcServer, handler) //handler implementira metode koje smo definisali
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func GetClient(uri, username, password string) (*neo4j.Driver, error) {

	auth := neo4j.BasicAuth(username, password, "")
	driver, err := neo4j.NewDriver(uri, auth)
	if err != nil {
		fmt.Println("nije se naprvaio neo4j klijent")
		fmt.Println(err)
		return nil, err
	}

	return &driver, nil //TODO: ref driver ?
}

func (server *Server) InitConnectionHandler(conSer *services.ConnectionService, userSer *services.UserService, logInfo *logger.Logger, logError *logger.Logger) *handlers.ConnectionHandler {
	return handlers.NewConnectionHandler(conSer, userSer, logInfo, logError)
}
func (server *Server) InitConnectionService(repo repositories.ConnectionRepository, logInfo *logger.Logger, logError *logger.Logger) *services.ConnectionService {
	return services.NewConnectionService(repo, logInfo, logError)
}

func (server *Server) InitConnectionRepository(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger) repositories.ConnectionRepository {
	return persistance.NewConnectionRepositoryImpl(client, logInfo, logError)
}

func (server *Server) InitUserService(repo repositories.UserRepository, logInfo *logger.Logger, logError *logger.Logger) *services.UserService {
	return services.NewUserService(repo, logInfo, logError)
}

func (server *Server) InitUserRepository(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger, connRepo repositories.ConnectionRepository) repositories.UserRepository {
	return persistance.NewUserRepositoryImpl(client, logInfo, logError, connRepo)
}

func (server *Server) InitSubscriber(subject string, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) InitPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) InitCreateUserCommandHandler(service *services.UserService, publisher saga.Publisher,
	subscriber saga.Subscriber) *handlers.UserCommandHandler {
	handler, err := handlers.NewUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return handler
}
