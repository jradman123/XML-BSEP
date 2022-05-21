package startup

import (
	user "common/module/proto/user_service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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
	l := log.New(os.Stdout, "products-api ", log.LstdFlags) // Logger koji dajemo handlerima
	db = server.SetupDatabase()
	userRepo := server.InitUserRepo(db)
	userService := server.InitUserService(l, userRepo)

	validator := validator.New()
	jsonConverters := helpers.NewJsonConverters(l)
	utils := helpers.PasswordUtil{}
	userHandler := server.InitUserHandler(l, userService, validator, jsonConverters, &utils)

	server.StartGrpcServer(userHandler)

}

func (server *Server) StartGrpcServer(handler *handlers.UserHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, handler) //handler implementira metode koje smo definisali
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) InitUserHandler(l *log.Logger, userService *services.UserService, validator *validator.Validate,
	jsonConverters *helpers.JsonConverters, passwordUtil *helpers.PasswordUtil) *handlers.UserHandler {
	return handlers.NewUserHandler(l, userService, jsonConverters, validator, passwordUtil)
}

func (server *Server) InitUserService(l *log.Logger, repo repositories.UserRepository) *services.UserService {
	return services.NewUserService(l, repo)
}

func (server *Server) InitUserRepo(d *gorm.DB) repositories.UserRepository {
	return persistance.NewUserRepositoryImpl(db)
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
	//db.Create(users) // Use this only once to populate db with data

	return db
}

//
//func initServer() {
//
//	router := mux.NewRouter()
//	s := http.Server{
//		Addr:         ":8082",           // configure the bind address
//		Handler:      router,            // set the default handler
//		ErrorLog:     l,                 // set the logger for the server
//		ReadTimeout:  5 * time.Second,   // max time to read request from the client
//		WriteTimeout: 10 * time.Second,  // max time to write response to the client
//		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
//	}
//	go func() {
//		l.Println("Starting server on port 8082")
//
//		err := s.ListenAndServe()
//		if err != nil {
//			l.Printf("Error starting server: %s\n", err)
//			os.Exit(1)
//		}
//	}()
//
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, os.Interrupt)
//	signal.Notify(c, os.Kill)
//
//	sig := <-c
//	log.Println("Got signal:", sig)
//
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//	s.Shutdown(ctx)
//}
