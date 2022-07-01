package startup

import (
	"common/module/logger"
	connGw "common/module/proto/connection_service"
	postsGw "common/module/proto/posts_service"
	userGw "common/module/proto/user_service"
	"context"
	"fmt"
	"gateway/module/application/helpers"
	"gateway/module/application/services"
	"gateway/module/domain/model"
	"gateway/module/domain/repositories"
	"gateway/module/infrastructure/handlers"
	"gateway/module/infrastructure/persistance"
	cfg "gateway/module/startup/config"
	gorilla_handlers "github.com/gorilla/handlers"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux // Part of grpcGateway library
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	//Povezuje sa grpc generisanim fajlovima
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(20*1024*1024),
			grpc.MaxCallSendMsgSize(20*1024*1024)),
	}
	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	postsEndpoint := fmt.Sprintf("%s:%s", server.config.PostsHost, server.config.PostsPort)
	connectionsEndpoint := fmt.Sprintf("%s:%s", server.config.ConnectionsHost, server.config.ConnectionsPort)

	err := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}
	err = postsGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postsEndpoint, opts)
	if err != nil {
		panic(err)
	}
	err = connGw.RegisterConnectionServiceHandlerFromEndpoint(context.TODO(), server.mux, connectionsEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

//Gateway ima svoje endpointe
func (server *Server) initCustomHandlers() {

	logInfo := logger.InitializeLogger("api-gateway", context.Background(), "Info")
	logError := logger.InitializeLogger("api-gateway", context.Background(), "Error")
	db = server.SetupDatabase()
	userRepo := server.InitUserRepo(db)
	tfauthRepo := server.InitTFAuthRepo(db)

	l := log.New(os.Stdout, "gateway ", log.LstdFlags) // Logger koji dajemo handlerima
	userService := server.InitUserService(l, logInfo, logError, userRepo)
	tfauthService := server.InitTFAuthService(l, tfauthRepo)
	lVerificationRepo := server.InitLoginVerificationRepo(db)
	passwordlessService := server.InitPasswordlessService(logInfo, logError, lVerificationRepo)

	validator := validator.New()

	passwordUtil := &helpers.PasswordUtil{}

	authHandler := handlers.NewAuthenticationHandler(l, logInfo, logError, userService, tfauthService, validator, passwordUtil, passwordlessService)
	authHandler.Init(server.mux)
	userFeedHandler := handlers.NewUserFeedHandler(logInfo, logError, server.config)
	userFeedHandler.Init(server.mux)
}

func (server *Server) Start() {
	cors := gorilla_handlers.CORS(
		gorilla_handlers.AllowedOrigins([]string{"https://localhost:4200", "https://localhost:4200/**", "http://localhost:4200", "http://localhost:4200/**", "http://localhost:8080/**"}),
		gorilla_handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorilla_handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Authorization", "Access-Control-Allow-*", "Access-Control-Allow-Origin", "*"}),
		gorilla_handlers.AllowCredentials(),
	)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), cors(muxMiddleware(server))))
}
func muxMiddleware(server *Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.mux.ServeHTTP(w, r)
	})
}

func (server *Server) InitUserService(l *log.Logger, logInfo *logger.Logger, logError *logger.Logger, repo repositories.UserRepository) *services.UserService {
	return services.NewUserService(l, logInfo, logError, repo)
}

func (server *Server) InitTFAuthService(l *log.Logger, repo repositories.TFAuthRepository) *services.TFAuthService {
	return services.NewTFAuthService(l, repo)
}
func (server *Server) InitUserRepo(db *gorm.DB) repositories.UserRepository {
	return persistance.NewUserRepositoryImpl(db)
}

func (server *Server) InitTFAuthRepo(db *gorm.DB) repositories.TFAuthRepository {
	return persistance.NewTFAuthRepositoryImpl(db)
}

var db *gorm.DB

func (server *Server) SetupDatabase() *gorm.DB {

	host := server.config.UserDBHost
	port := server.config.UserDBPort
	user := server.config.UserDBUser
	dbname := server.config.UserDBName
	password := server.config.UserDBPass

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database!")
	}

	db.AutoMigrate(&model.User{}) //This will not remove columns
	db.AutoMigrate(&model.QrCode{})
	db.AutoMigrate(&model.LoginVerification{}) //This will not remove columns
	//db.Create(users) // Use this only once to populate db with data

	return db
}

func (server *Server) InitPasswordlessService(logInfo *logger.Logger, logError *logger.Logger, repo repositories.LoginVerificationRepository) *services.PasswordLessService {
	return services.NewPasswordLessService(logInfo, logError, repo)
}

func (server *Server) InitLoginVerificationRepo(db *gorm.DB) repositories.LoginVerificationRepository {
	return persistance.NewLoginVerificationRepositoryImpl(db)
}
