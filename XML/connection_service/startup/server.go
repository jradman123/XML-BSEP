package startup

import (
	"common/module/interceptor"
	connectionProto "common/module/proto/connection_service"
	"connection/module/application/services"
	"connection/module/infrastructure/handlers"
	"connection/module/startup/config"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
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
	l := log.New(os.Stdout, "connection-service-server ", log.LstdFlags)
	l.Println("starting connection service server")
	neoClient := server.SetupDatabase()
	//vidjecu dal cu imati repo ili nesto malo drugacije
	//validator := validator.New()
	connectionService := server.InitConnectionService(l, neoClient)
	connectionHandler := server.InitConnectionHandler(l, connectionService)
	server.StartGrpcServer(connectionHandler)

}

func (server *Server) SetupDatabase() *neo4j.Driver {
	client, err := GetClient(server.config.Neo4jUri, server.config.Neo4jUsername, server.config.Neo4jPassword)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) StartGrpcServer(handler *handlers.ConnectionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
	if err != nil {
		log.Fatalf("failed to parse public key: %v", err)
	}
	//interceptor := interceptor.NewAuthInterceptor(config.AccessibleRoles(), publicKey)
	interceptor := interceptor.NewAuthInterceptor(config.AccessibleRoles(), server.config.PublicKey)

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

func (server *Server) InitConnectionHandler(l *log.Logger, conSer *services.ConnectionService) *handlers.ConnectionHandler {
	return handlers.NewConnectionHandler(l, conSer)
}
func (server *Server) InitConnectionService(l *log.Logger, neoClient *neo4j.Driver) *services.ConnectionService {
	return services.NewConnectionService(l, neoClient)
}
