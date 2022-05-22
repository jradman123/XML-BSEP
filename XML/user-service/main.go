package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"user/module/handlers"
	"user/module/helpers"
	my_middleware "user/module/middleware"
	"user/module/model"
	"user/module/repository"
	"user/module/service"

	"github.com/euroteltr/rbac"
	"github.com/euroteltr/rbac/middlewares/echorbac"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	hibp "github.com/mattevans/pwned-passwords"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDatabase() *gorm.DB {

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

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.EmailVerification{}) //This will not remove columns
	db.AutoMigrate(&model.PasswordRecoveryRequest{})
	//db.Create(users) // Use this only once to populate db with data

	return db
}

func initPasswordUtil() *helpers.PasswordUtil {
	return &helpers.PasswordUtil{}
}

func initRegisteredUserRepo(database *gorm.DB) *repository.RegisteredUserRepository {
	return &repository.RegisteredUserRepository{DB: database}
}
func initEmailVerificationRepo(database *gorm.DB) *repository.EmailVerificationRepository {
	return &repository.EmailVerificationRepository{DB: database}
}

func initPasswordRecoveryRepo(database *gorm.DB) *repository.PasswordRecoveryRepository {
	return &repository.PasswordRecoveryRepository{DB: database}
}

func initRegisterUserService(repo *repository.RegisteredUserRepository, verRepo *repository.EmailVerificationRepository, recoveryRepo *repository.PasswordRecoveryRepository) *service.RegisteredUserService {
	return &service.RegisteredUserService{Repo: repo,
		EmailRepo:    verRepo,
		RecoveryRepo: recoveryRepo}
}

var R *rbac.RBAC

func main() {

	db = SetupDatabase()

	R := rbac.New(rbac.NewConsoleLogger())
	adminRole, _ := R.RegisterRole("admin", "Admin role")
	userRole, _ := R.RegisterRole("user", "User role")

	ApproveAction := rbac.Action("approve") //costume action
	GetAllAction := rbac.Action("getAll")
	usersPerm, err := R.RegisterPermission("users", "User resource", rbac.Read, rbac.Create, ApproveAction, GetAllAction)
	if err != nil {
		panic(err)
	}

	// Now load or define roles
	R.Permit(adminRole.ID, usersPerm, usersPerm.Actions()...) // OVDJE DAJEMO PERMISSION
	R.Permit(userRole.ID, usersPerm, rbac.Read)

	// Middleware function shorthand
	isGranted := echorbac.HasRole(R)
	isGranted(usersPerm)

	fmt.Printf("Admin is granted permission over users : %v\n", R.IsGranted(adminRole.ID, usersPerm, usersPerm.Actions()...)) //should be true
	//-------------------------------------------------------------------------------------------------//

	//--------jelena-------
	pwnedClient := hibp.NewClient()
	l := log.New(os.Stdout, "products-api ", log.LstdFlags) // Logger koji dajemo handlerima
	validator := validator.New()
	jsonConverters := helpers.NewJsonConverters(l)
	repository := repository.NewUserRepository(db)
	registerdUserRepo := initRegisteredUserRepo(db)
	emailVerificationRepo := initEmailVerificationRepo(db)
	passwordRecoveryRepo := initPasswordRecoveryRepo(db)
	userService := service.NewUserService(l, repository)
	registerUserService := initRegisterUserService(registerdUserRepo, emailVerificationRepo, passwordRecoveryRepo)
	passwordUtil := initPasswordUtil()
	userHandler := handlers.NewUserHandler(l, userService, registerUserService, jsonConverters, &repository, validator, passwordUtil, pwnedClient)

	//--------jelena-------
	//l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	//repository := repository.NewUserRepository()
	//userService := service.NewUserService(l, repository)
	//userHandler := handlers.NewUserHandler(l, *userService)
	//-----------------------------------------------------------------------------------------------//

	router := mux.NewRouter()
	UnauthorizedPostRouter := router.Methods(http.MethodPost).Subrouter()
	UnauthorizedPostRouter.HandleFunc("/login", userHandler.LoginUser)
	UnauthorizedPostRouter.HandleFunc("/register", userHandler.AddUsers)
	UnauthorizedPostRouter.HandleFunc("/pwnedPassword", userHandler.CheckIfPwned)
	UnauthorizedPostRouter.HandleFunc("/activateAccount", userHandler.ActivateUserAccount)
	UnauthorizedPostRouter.HandleFunc("/recoverPasswordRequest", userHandler.RecoverPasswordRequest)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	authMiddleware := my_middleware.NewAuthorizationHandler(*R, usersPerm, usersPerm.Actions(), userService)
	getRouter.Use(my_middleware.ValidateToken, authMiddleware.PermissionGranted)
	getRouter.HandleFunc("/", userHandler.GetUsers)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.Use(my_middleware.ValidateToken)
	putRouter.HandleFunc("/{id:[0-9]+}", userHandler.UpdateUsers)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	putRouter.Use(my_middleware.ValidateToken)
	postRouter.HandleFunc("/", userHandler.AddUsers)
	postRouter.HandleFunc("/pwnedPassword", userHandler.CheckIfPwned)
	postRouter.HandleFunc("/activateAccount", userHandler.ActivateUserAccount)
	postRouter.HandleFunc("/recoverPasswordRequest", userHandler.RecoverPasswordRequest)

	// create a new server
	s := http.Server{
		Addr:         ":8082",           // configure the bind address
		Handler:      router,            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 8082")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
