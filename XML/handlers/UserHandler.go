package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"user/module/auth"
	"user/module/dto"
	"user/module/service"
)

// UserHandler is a http.Handler
type UserHandler struct {
	l       *log.Logger
	service service.UserService
}

func NewUserHandler(l *log.Logger, service service.UserService) *UserHandler {
	return &UserHandler{l, service}
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
func (u *UserHandler) AddUsers(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handling POST Users")
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
