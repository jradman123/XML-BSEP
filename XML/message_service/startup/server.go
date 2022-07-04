package startup

import (
	"common/module/interceptor"
	"common/module/logger"
	messagesProto "common/module/proto/message_service"
	notificationProto "common/module/proto/notification_service"
	saga "common/module/saga/messaging"
	"common/module/saga/messaging/nats"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pusher/pusher-http-go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"message/module/application"
	"message/module/domain/repositories"
	"message/module/infrastructure/handlers"
	"message/module/infrastructure/persistence"
	"message/module/startup/config"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{config: config}
}

const (
	QueueGroup = "message_service"
)

func (server *Server) Start() {
	logInfo := logger.InitializeLogger("post-service", context.Background(), "Info")
	logError := logger.InitializeLogger("post-service", context.Background(), "Error")

	mongoClient := server.InitMongoClient()
	notificationPusher := server.InitNotificationPusher()

	messageRepo := server.InitMessageRepo(mongoClient)
	messageService := server.InitMessageService(messageRepo, logInfo, logError)

	notificationRepo := server.InitNotificationRepo(mongoClient)
	notificationService := server.InitNotificationService(logInfo, logError, notificationRepo)

	commandSubscriber := server.InitSubscriber(server.config.UserCommandSubject, QueueGroup)
	replyPublisher := server.InitPublisher(server.config.UserReplySubject)
	userRepo := server.InitUserRepo(mongoClient)
	userService := server.InitUserService(userRepo, logInfo, logError)

	messageHandler := server.InitMessageHandler(messageService, userService, logInfo, logError)
	notificationHandler := server.InitNotificationHandler(logInfo, logError, &notificationPusher, notificationService, userService)
	server.InitCreateUserCommandHandler(userService, messageService, replyPublisher, commandSubscriber)

	server.StartGrpcServer(messageHandler, notificationHandler, logError)
}

func (server *Server) InitMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.MessageDBHost, server.config.MessageDBPort)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connected to mongo database!")
	}
	return client
}

func (server *Server) InitMessageRepo(client *mongo.Client) repositories.MessageRepository {
	return persistence.NewMessageRepositoryImpl(client)
}

func (server *Server) InitMessageService(repo repositories.MessageRepository, logInfo *logger.Logger, logError *logger.Logger) *application.MessageService {
	return application.NewMessageService(repo, logInfo, logError)
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

func (server *Server) InitMessageHandler(messageService *application.MessageService, userService *application.UserService, logInfo *logger.Logger, logError *logger.Logger) *handlers.MessageHandler {
	return handlers.NewMessageHandler(messageService, userService, logInfo, logError)
}

func (server *Server) InitCreateUserCommandHandler(userService *application.UserService, postService *application.MessageService, publisher saga.Publisher,
	subscriber saga.Subscriber) *handlers.UserCommandHandler {
	handler, err := handlers.NewUserCommandHandler(userService, postService, publisher, subscriber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return handler
}

func (server *Server) StartGrpcServer(messageHandler *handlers.MessageHandler, notificationHandler *handlers.NotificationHandler, logError *logger.Logger) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
	intercept := interceptor.NewAuthInterceptor(config.AccessibleRoles(), publicKey, logError)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(intercept.Unary()))
	messagesProto.RegisterMessageServiceServer(grpcServer, messageHandler)
	notificationProto.RegisterNotificationServiceServer(grpcServer, notificationHandler)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) InitNotificationPusher() pusher.Client {
	pusherClient := pusher.Client{
		AppID:   server.config.NotificationAppID,
		Key:     server.config.NotificationKey,
		Secret:  server.config.NotificationSecret,
		Cluster: server.config.NotificationCluster,
		Secure:  server.config.NotificationSecure,
	}
	return pusherClient
}

func (server *Server) InitNotificationRepo(client *mongo.Client) repositories.NotificationRepository {
	return persistence.NewNotificationRepositoryImpl(client)
}

func (server *Server) InitNotificationService(info *logger.Logger, logError *logger.Logger, repo repositories.NotificationRepository) *application.NotificationService {
	return application.NewNotificationService(info, logError, repo)
}

func (server *Server) InitNotificationHandler(info *logger.Logger, logError *logger.Logger, notificationPusher *pusher.Client, service *application.NotificationService, userService *application.UserService) *handlers.NotificationHandler {
	return handlers.NewNotificationHandler(info, logError, notificationPusher, service, userService)
}
