package startup

import (
	"common/module/interceptor"
	"common/module/logger"
	connectionProto "common/module/proto/connection_service"
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

func (server *Server) Start() {
	logInfo := logger.InitializeLogger("connection-service", context.Background(), "Info")
	logError := logger.InitializeLogger("connection-service", context.Background(), "Error")
	fmt.Println("starting connection service server")
	neoClient := server.SetupDatabase()
	//vidjecu dal cu imati repo ili nesto malo drugacije
	//validator := validator.New()
	connectionRepo := server.InitConnectionRepository(neoClient, logInfo, logError)
	connectionService := server.InitConnectionService(connectionRepo, logInfo, logError)
	connectionRequestRepo := server.InitConnectionRequestRepository(neoClient, logInfo, logError)
	connectionRequestService := server.InitConnectionRequestService(connectionRequestRepo, logInfo, logError)
	userRepo := server.InitUserRepository(neoClient, logInfo, logError)
	userService := server.InitUserService(userRepo, logInfo, logError)
	connectionHandler := server.InitConnectionHandler(connectionService, connectionRequestService, userService, logInfo, logError)
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

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &driver, nil //TODO: ref driver ?
}

func (server *Server) InitConnectionHandler(conSer *services.ConnectionService, conReqSer *services.ConnectionRequestService, userSer *services.UserService, logInfo *logger.Logger, logError *logger.Logger) *handlers.ConnectionHandler {
	return handlers.NewConnectionHandler(conSer, conReqSer, userSer, logInfo, logError)
}
func (server *Server) InitConnectionService(repo repositories.ConnectionRepository, logInfo *logger.Logger, logError *logger.Logger) *services.ConnectionService {
	return services.NewConnectionService(repo, logInfo, logError)
}

func (server *Server) InitConnectionRepository(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger) repositories.ConnectionRepository {
	return persistance.NewConnectionRepositoryImpl(client, logInfo, logError)
}

func (server *Server) InitConnectionRequestService(repo repositories.ConnectionRequestRepository, logInfo *logger.Logger, logError *logger.Logger) *services.ConnectionRequestService {
	return services.NewConnectionRequestService(repo, logInfo, logError)
}

func (server *Server) InitConnectionRequestRepository(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger) repositories.ConnectionRequestRepository {
	return persistance.NewConnectionRequestRepositoryImpl(client, logInfo, logError)
}

func (server *Server) InitUserService(repo repositories.UserRepository, logInfo *logger.Logger, logError *logger.Logger) *services.UserService {
	return services.NewUserService(repo, logInfo, logError)
}

func (server *Server) InitUserRepository(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger) repositories.UserRepository {
	return persistance.NewUserRepositoryImpl(client, logInfo, logError)
}
