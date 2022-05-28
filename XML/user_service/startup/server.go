package startup

import (
	"common/module/interceptor"
	userProto "common/module/proto/user_service"
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
func (server *Server) Start() {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	pwnedClient := hibp.NewClient()
	db = server.SetupDatabase()
	userRepo := server.InitUserRepo(db)
	emailVerRepo := server.InitEmailVerRepo(db)
	recoveryRepo := server.InitRecoveryRepo(db)
	userService := server.InitUserService(l, userRepo, emailVerRepo, recoveryRepo)

	validator := validator.New()
	jsonConverters := helpers.NewJsonConverters(l)
	utils := helpers.PasswordUtil{}
	userHandler := server.InitUserHandler(l, userService, validator, jsonConverters, &utils, pwnedClient)

	server.StartGrpcServer(userHandler)

}

func (server *Server) StartGrpcServer(handler *handlers.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
	if err != nil {
		log.Fatalf("failed to parse public key: %v", err)
	}
	interceptor := interceptor.NewAuthInterceptor(config.AccessibleRoles(), publicKey)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))

	userProto.RegisterUserServiceServer(grpcServer, handler) //handler implementira metode koje smo definisali
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) InitUserHandler(l *log.Logger, userService *services.UserService, validator *validator.Validate,
	jsonConverters *helpers.JsonConverters, passwordUtil *helpers.PasswordUtil, pwnedClient *hibp.Client) *handlers.UserHandler {
	return handlers.NewUserHandler(l, userService, jsonConverters, validator, passwordUtil, pwnedClient)
}

func (server *Server) InitUserService(l *log.Logger, repo repositories.UserRepository, emailRepo repositories.EmailVerificationRepository, recoveryRepo repositories.PasswordRecoveryRequestRepository) *services.UserService {
	return services.NewUserService(l, repo, emailRepo, recoveryRepo)
}

func (server *Server) InitUserRepo(d *gorm.DB) repositories.UserRepository {
	return persistance.NewUserRepositoryImpl(db)
}

func (server *Server) InitEmailVerRepo(d *gorm.DB) repositories.EmailVerificationRepository {
	return persistance.NewEmailVerificationRepositoryImpl(db)
}

func (server *Server) InitRecoveryRepo(d *gorm.DB) repositories.PasswordRecoveryRequestRepository {
	return persistance.NewPasswordRecoveryRequestRepositoryImpl(d)
}

var db *gorm.DB

func (server *Server) SetupDatabase() *gorm.DB {

	host := os.Getenv("HOST")
	port := os.Getenv("PG_DBPORT")
	user := os.Getenv("PG_USER")
	dbname := os.Getenv("XML_DB_NAME")
	password := os.Getenv("PG_PASSWORD")

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
	//db.Create(users) // Use this only once to populate db with data

	return db
}
