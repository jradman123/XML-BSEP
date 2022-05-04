package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"user/module/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

func SetupDatabase() {
	//Loading env variables

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// DB COnnection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	//Opening connection to DB
	db, err = gorm.Open(dialect, dbURI)

	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("Successfully connected to database!")

	}

	//Close connection when the main name finishes

	defer db.Close()

	// Make database migrations to databaseif
	//db.DropTable(&model.User{})
	db.AutoMigrate(&model.User{}) //This will not remove columns
	//db.Create(users) // Use this only once to populate db with data

}
func main() {

	SetupDatabase()

	router := mux.NewRouter()
	router.HandleFunc("/", getUsers).Methods("GET")

}
func getUsers(rw http.ResponseWriter, r *http.Request) {

}
