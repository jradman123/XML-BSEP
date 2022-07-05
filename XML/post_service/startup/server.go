package startup

import (
	"common/module/interceptor"
	"common/module/logger"
	postsProto "common/module/proto/posts_service"
	saga "common/module/saga/messaging"
	"common/module/saga/messaging/nats"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"post/module/application"
	"post/module/domain/repositories"
	"post/module/infrastructure/handlers"
	"post/module/infrastructure/orchestrators"
	"post/module/infrastructure/persistence"
	"post/module/startup/config"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{config: config}
}

const (
	QueueGroupUser = "post_service_user"
	QueueGroupPost = "post_service_post"
)

func (server *Server) Start() {
	logInfo := logger.InitializeLogger("post-service", context.Background(), "Info")
	logError := logger.InitializeLogger("post-service", context.Background(), "Error")

	mongoClient := server.InitMongoClient()

	postRepo := server.InitPostsRepo(mongoClient)

	commandSubscriber := server.InitSubscriber(server.config.UserCommandSubject, QueueGroupUser)
	replyPublisher := server.InitPublisher(server.config.UserReplySubject)

	commandPublisher := server.InitPublisher(server.config.PostNotificationCommandSubject)
	replySubscriber := server.InitSubscriber(server.config.PostNotificationReplySubject, QueueGroupPost)

	orchestrator := server.InitOrchestrator(commandPublisher, replySubscriber)
	postService := server.InitPostService(postRepo, logInfo, logError, orchestrator)

	userRepo := server.InitUserRepo(mongoClient)
	userService := server.InitUserService(userRepo, logInfo, logError)
	postHandler := server.InitPostHandler(postService, userService, logInfo, logError)
	server.InitCreateUserCommandHandler(userService, postService, replyPublisher, commandSubscriber)

	server.StartGrpcServer(postHandler, logError)
}

func (server *Server) InitMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connected to mongo database!")
	}

	return client
}

func (server *Server) InitOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *orchestrators.PostOrchestrator {
	orchestrator, err := orchestrators.NewPostOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) InitPostsRepo(client *mongo.Client) repositories.PostRepository {
	return persistence.NewPostRepositoryImpl(client)
}

func (server *Server) InitPostService(repo repositories.PostRepository, logInfo *logger.Logger, logError *logger.Logger, orchestrator *orchestrators.PostOrchestrator) *application.PostService {
	return application.NewPostService(repo, logInfo, logError, orchestrator)
}

func (server *Server) InitPostHandler(postService *application.PostService, userService *application.UserService, logInfo *logger.Logger, logError *logger.Logger) *handlers.PostHandler {
	return handlers.NewPostHandler(postService, userService, logInfo, logError)
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

func (server *Server) InitUserRepo(client *mongo.Client) repositories.UserRepository {
	return persistence.NewUserRepositoryImpl(client)
}

func (server *Server) InitUserService(repo repositories.UserRepository, logInfo *logger.Logger, logError *logger.Logger) *application.UserService {
	return application.NewUserService(repo, logInfo, logError)
}

func (server *Server) InitCreateUserCommandHandler(userService *application.UserService, postService *application.PostService, publisher saga.Publisher,
	subscriber saga.Subscriber) *handlers.UserCommandHandler {
	handler, err := handlers.NewUserCommandHandler(userService, postService, publisher, subscriber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return handler
}
func (server *Server) StartGrpcServer(postHandler *handlers.PostHandler, logError *logger.Logger) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
	intercept := interceptor.NewAuthInterceptor(config.AccessibleRoles(), publicKey, logError)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(intercept.Unary()))
	postsProto.RegisterPostServiceServer(grpcServer, postHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
