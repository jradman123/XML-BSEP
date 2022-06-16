package handlers

import (
	"common/module/interceptor"
	"context"
	"encoding/json"
	"fmt"
	"gateway/module/application/helpers"
	"gateway/module/application/services"
	"gateway/module/auth"
	"gateway/module/domain/dto"
	modelGateway "gateway/module/domain/model"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/skip2/go-qrcode"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AuthenticationHandler struct {
	l            *log.Logger
	service      *services.UserService
	validator    *validator.Validate
	passwordUtil *helpers.PasswordUtil
}

func NewAuthenticationHandler(l *log.Logger, service *services.UserService, validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil) Handler {
	return &AuthenticationHandler{l, service, validator, passwordUtil}
}

func (a AuthenticationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/users/login/user", a.LoginUser)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/2fa/check", a.Check2FaForUser)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/2fa", a.TwoFactorAuth) //TODO:MAKE REDUNDANT
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/2fa/enable", a.Enable2FaForUser)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/2fa/disable", a.Disable2FaForUser)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/2fa/authenticate", a.Authenticate2Fa)
	if err != nil {
		panic(err)
	}
}

func (a AuthenticationHandler) Check2FaForUser(w http.ResponseWriter, r *http.Request, params map[string]string) {

}

func (a AuthenticationHandler) Enable2FaForUser(w http.ResponseWriter, r *http.Request, params map[string]string) {

}

func (a AuthenticationHandler) Disable2FaForUser(w http.ResponseWriter, r *http.Request, params map[string]string) {

}
func (a AuthenticationHandler) Authenticate2Fa(rw http.ResponseWriter, r *http.Request, parameters map[string]string) {
	//get user secret from username

	otpc := NewOTPConfig("TopSecret1234")

	var token int
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		http.Error(rw, "Error decoding Authenticate2Fa:"+err.Error(), http.StatusBadRequest)
		return
	}

	val, err := otpc.Authenticate(strconv.Itoa(token))
	if err != nil {
		fmt.Println(err)
		return
	}

	if !val {
		fmt.Println("Sorry, Not Authenticated")
		return
	}

	fmt.Println("Authenticated!")
}

func (a AuthenticationHandler) TwoFactorAuth(rw http.ResponseWriter, r *http.Request, parameters map[string]string) {
	twofa := NewOTPConfig("TopSecret1234")
	uri := twofa.ProvisionURI("testuser")
	log.Println("This is URI: " + uri)
	err := qrcode.WriteFile(uri, qrcode.Medium, 256, "qr1.png")
	if err != nil {
		http.Error(rw, "Err generating qr "+err.Error(), http.StatusBadRequest)
		return
	}
}
func (a AuthenticationHandler) LoginUser(rw http.ResponseWriter, r *http.Request, params map[string]string) {
	a.l.Println("Handling LOGIN Users")

	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.service.GetByUsername(context.TODO(), loginRequest.Username)
	if err != nil {
		http.Error(rw, "Invalid credentials!", http.StatusBadRequest)
		return
	}
	if !user.IsConfirmed {
		http.Error(rw, "Check your mail for activation code! ", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		http.Error(rw, "Invalid credentials!", http.StatusBadRequest)
		return
	}

	var claims = &interceptor.JwtClaims{}
	claims.Username = loginRequest.Username

	userRoles, err := a.service.GetUserRole(loginRequest.Username)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return

	}

	claims.Roles = append(claims.Roles, userRoles)
	var tokenCreationTime = time.Now().UTC()
	var tokenExpirationTime = tokenCreationTime.Add(time.Duration(30) * time.Minute)

	token, err := auth.GenerateToken(claims, tokenExpirationTime)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return

	}

	var roleString string

	if user.Role == modelGateway.Admin {
		roleString = "Admin"
	} else if user.Role == modelGateway.Agent {
		roleString = "Agent"
	} else if user.Role == modelGateway.Regular {
		roleString = "Regular"
	}

	logInResponse := dto.LogInResponseDto{
		Token:    token,
		Role:     roleString,
		Email:    user.Email,
		Username: user.Username,
	}

	logInResponseJson, _ := json.Marshal(logInResponse)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(logInResponseJson)

}
