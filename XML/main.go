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
	mymiddleware "user/module/middleware"
	"user/module/model"
	"user/module/repository"
	"user/module/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" //na ovom importu se crveni ali bez njega nece da radi
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	users = &model.User{
		ID:          uuid.New(),
		Username:    "Jack",
		Password:    "abc123",
		Email:       "jack@gmail.com",
		PhoneNumber: "123123",
		FirstName:   "Jack",
		LastName:    "Sparrow",
		Gender:      model.MALE,
	}
)

var db *gorm.DB
var err error

func SetupDatabase() *gorm.DB {
	// //Loading env variables

	// //dialect := os.Getenv("DIALECT")
	// host := os.Getenv("HOST")
	// port := os.Getenv("DBPORT")
	// user := os.Getenv("USER")
	// dbname := os.Getenv("NAME")
	// password := os.Getenv("PASSWORD")

	// //dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// //Opening connection to DB
	// //db, err = sql.Open(dialect, dbURI)
	// db, err := sql.Open("postgres", psqlInfo)

	// if err != nil {
	// 	log.Fatal(err)

	// } else {
	// 	fmt.Println("Successfully connected to database!")

	// }

	// //Close connection when the main name finishes

	// defer db.Close()

	// //Make database migrations to databaseif
	// //db.DropTable(&model.User{})
	// //db.AutoMigrate(&model.User{}) //This will not remove columns
	// //db.Create(users) // Use this only once to populate db with data
	host := os.Getenv("HOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	//psqlInfo := fmt.Sprintf("host=localhost  user=postgres password=fakultet dbname=xws_project	sslmode=disable")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//Opening connection to DB
	//db, err = sql.Open(dialect, dbURI)
	//j db, err := sql.Open("postgres", psqlInfo)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("Successfully connected to database!")

	}

	//Close connection when the main name finishes

	//j defer db.Close()

	//Make database migrations to databaseif
	//db.DropTable(&model.User{})
	db.AutoMigrate(&model.User{}) //This will not remove columns
	//db.Create(users) // Use this only once to populate db with data

	return db
}
func main() {

	//ovo postaviti kao promjenljive sistema
	os.Setenv("HOST", "localhost")
	os.Setenv("DBPORT", "5432")
	os.Setenv("USER", "postgres")
	os.Setenv("PASSWORD", "fakultet")
	os.Setenv("NAME", "xws_project")

	db = SetupDatabase()

	l := log.New(os.Stdout, "products-api ", log.LstdFlags) // Logger koji dajemo handlerima
	jsonConverters := helpers.NewJsonConverters(l)
	repository := repository.NewUserRepository(db)
	service := service.NewUserService(l, repository)
	userHandler := handlers.NewUserHandler(l, *service, *jsonConverters, repository)

	router := mux.NewRouter()
	UnauthorizedPostRouter := router.Methods(http.MethodPost).Subrouter()
	UnauthorizedPostRouter.HandleFunc("/login", userHandler.LoginUser)
	UnauthorizedPostRouter.HandleFunc("/register", userHandler.AddUsers)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", userHandler.GetUsers)
	getRouter.Use(mymiddleware.ValidateToken)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", userHandler.UpdateUsers)
	putRouter.Use(mymiddleware.ValidateToken)

	//postRouter := router.Methods(http.MethodPost).Subrouter()
	//postRouter.HandleFunc("/", userHandler.AddUsers)
	putRouter.Use(mymiddleware.ValidateToken)

	// create a new server
	s := http.Server{
		Addr:         ":8081",           // configure the bind address
		Handler:      router,            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 8081")

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
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
