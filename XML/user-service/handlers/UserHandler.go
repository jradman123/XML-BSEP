package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"user/module/auth"
	"user/module/dto"
	"user/module/helpers"
	"user/module/repository"
	"user/module/service"

	"golang.org/x/crypto/bcrypt"
)

// UserHandler is a http.Handler
type UserHandler struct {
	l        *log.Logger
	service  service.UserService
	jsonConv helpers.JsonConverters
	repo     repository.UserRepository
}

type ResponseEmail struct {
	Email string
}
type ErrorResponse struct {
	Err string
}

func NewUserHandler(l *log.Logger, service service.UserService, jsonConv helpers.JsonConverters, repo repository.UserRepository) *UserHandler {
	return &UserHandler{l, service, jsonConv, repo}
}

func (u *UserHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {

	u.l.Println("Handling GET Users")
	users, err := u.service.GetUsers()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	logInResponseJson, _ := json.Marshal(users)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(logInResponseJson)
}
func (u *UserHandler) AddUsers(rw http.ResponseWriter, req *http.Request) {
	// u.l.Println("Handling POST Users")
	// contentType := req.Header.Get("Content-Type")
	// mediatype, _, err := mime.ParseMediaType(contentType)
	// if err != nil {
	// 	http.Error(rw, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// if mediatype != "application/json" {
	// 	http.Error(rw, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
	// 	return
	// }

	var newUser dto.NewUser
	err := json.NewDecoder(req.Body).Decode(&newUser)

	if err != nil {

		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest)
	}
	//cuvamo password kao hash neki
	pass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(rw).Encode(err)
	}

	newUser.Password = string(pass)

	//zakucana rola za sad
	var role = "REGISTERED_USER"

	email := u.repo.CreateUser(nil, newUser.Username, newUser.Password, newUser.Email, newUser.PhoneNumber, newUser.FirstName, newUser.LastName, newUser.Gender, role)

	userEmail := ResponseEmail{
		Email: email,
	}

	userEmailJson, _ := json.Marshal(userEmail)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(userEmailJson)

}

func (u *UserHandler) UpdateUsers(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handling PUT Users")
}
func (u *UserHandler) LoginUser(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handling LOGIN Users")

	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {

		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest)

	}

	var claims = &auth.JwtClaims{}
	claims.Username = loginRequest.Username
	claims.Roles = []string{"admin", "user"}

	var tokenCreationTime = time.Now().UTC()
	var tokenExpirationTime = tokenCreationTime.Add(time.Duration(30) * time.Minute)

	token, err := auth.GenerateToken(claims, tokenExpirationTime)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	logInResponse := dto.LogInResponseDto{
		Token: token,
	}

	logInResponseJson, _ := json.Marshal(logInResponse)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(logInResponseJson)

}
