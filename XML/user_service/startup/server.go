package startup

import (
	"common/module/interceptor"
	"common/module/logger"
	userProto "common/module/proto/user_service"
	saga "common/module/saga/messaging"
	"common/module/saga/messaging/nats"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	hibp "github.com/mattevans/pwned-passwords"
	"google.golang.org/grpc"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"user/module/application/helpers"
	"user/module/application/services"
	"user/module/domain/model"
	"user/module/domain/repositories"
	"user/module/infrastructure/handlers"
	"user/module/infrastructure/orchestrators"
	"user/module/infrastructure/persistance"
	"user/module/startup/config"
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
	QueueGroup = "user_service"
)

func (server *Server) Start() {
	logInfo := logger.InitializeLogger("user-service", context.Background(), "Info")
	logError := logger.InitializeLogger("user-service", context.Background(), "Error")
	pwnedClient := hibp.NewClient()
	db = server.SetupDatabase()
	userRepo := server.InitUserRepo(db)
	emailVerRepo := server.InitEmailVerRepo(db)
	recoveryRepo := server.InitRecoveryRepo(db)

	commandPublisher := server.InitPublisher(server.config.UserCommandSubject)
	replySubscriber := server.InitSubscriber(server.config.UserReplySubject, QueueGroup)
	orchestrator := server.InitOrchestrator(commandPublisher, replySubscriber)

	userService := server.InitUserService(logInfo, logError, userRepo, emailVerRepo, recoveryRepo, orchestrator)
	apiTokenService := server.InitApiTokenService(logInfo, logError, userService)

	validator := validator.New()
	jsonConverters := helpers.NewJsonConverters(logInfo)
	utils := helpers.PasswordUtil{}

	userHandler := server.InitUserHandler(logInfo, logError, userService, validator, jsonConverters, &utils, pwnedClient, apiTokenService)

	server.StartGrpcServer(userHandler, logError)

}

func (server *Server) StartGrpcServer(handler *handlers.UserHandler, logError *logger.Logger) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
	if err != nil {
		log.Fatalf("failed to parse public key: %v", err)
	}
	interceptor := interceptor.NewAuthInterceptor(config.AccessibleRoles(), publicKey, logError)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))
	userProto.RegisterUserServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) InitUserHandler(logInfo *logger.Logger, logError *logger.Logger, userService *services.UserService, validator *validator.Validate,
	jsonConverters *helpers.JsonConverters, passwordUtil *helpers.PasswordUtil, pwnedClient *hibp.Client, tokenService *services.ApiTokenService) *handlers.UserHandler {
	return handlers.NewUserHandler(logInfo, logError, userService, jsonConverters, validator, passwordUtil, pwnedClient, tokenService)
}

func (server *Server) InitUserService(logInfo *logger.Logger, logError *logger.Logger, repo repositories.UserRepository,
	emailRepo repositories.EmailVerificationRepository, recoveryRepo repositories.PasswordRecoveryRequestRepository, orchestrator *orchestrators.UserOrchestrator) *services.UserService {
	return services.NewUserService(logInfo, logError, repo, emailRepo, recoveryRepo, orchestrator)
}

func (server *Server) InitApiTokenService(logInfo *logger.Logger, logError *logger.Logger, userService *services.UserService) *services.ApiTokenService {
	return services.NewApiTokenService(logInfo, logError, userService)
}

func (server *Server) InitUserRepo(db *gorm.DB) repositories.UserRepository {
	return persistance.NewUserRepositoryImpl(db)
}

func (server *Server) InitEmailVerRepo(db *gorm.DB) repositories.EmailVerificationRepository {
	return persistance.NewEmailVerificationRepositoryImpl(db)
}

func (server *Server) InitRecoveryRepo(d *gorm.DB) repositories.PasswordRecoveryRequestRepository {
	return persistance.NewPasswordRecoveryRequestRepositoryImpl(d)
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

func (server *Server) InitSubscriber(subject string, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) InitOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *orchestrators.UserOrchestrator {
	orchestrator, err := orchestrators.NewUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

var db *gorm.DB

func (server *Server) SetupDatabase() *gorm.DB {

	host := os.Getenv("HOST")
	port := os.Getenv("PG_DBPORT")
	user := os.Getenv("PG_USER")
	dbname := os.Getenv("XML_DB_NAME")
	password := os.Getenv("PG_PASSWORD")

	//host := os.Getenv("USER_DB_HOST")
	//port := os.Getenv("USER_DB_PORT")
	//user := os.Getenv("USER_DB_USER")
	//dbname := os.Getenv("USER_DB_NAME")
	//password := os.Getenv("USER_DB_PASS")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}

	db.AutoMigrate(&model.User{}) //This will not remove columns
	db.AutoMigrate(&model.EmailVerification{})
	db.AutoMigrate(&model.PasswordRecoveryRequest{})
	db.AutoMigrate(&model.Skill{})
	db.AutoMigrate(&model.Experience{})
	db.AutoMigrate(&model.Education{})
	db.AutoMigrate(&model.Interest{})
	//db.Create(users) // Use this only once to populate db with data

	return db
}
