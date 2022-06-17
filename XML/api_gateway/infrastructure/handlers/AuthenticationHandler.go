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
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type AuthenticationHandler struct {
	l            *log.Logger
	userService  *services.UserService
	tfaService   *services.TFAuthService
	validator    *validator.Validate
	passwordUtil *helpers.PasswordUtil
}

func NewAuthenticationHandler(l *log.Logger, service *services.UserService, tfaService *services.TFAuthService,
	validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil) Handler {
	return &AuthenticationHandler{l, service, tfaService, validator, passwordUtil}
}

func (a AuthenticationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/users/auth/user", a.AuthenticateUser)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/users/login/user", a.LoginUser)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/2fa/check", a.Check2FaForUser)
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

func (a AuthenticationHandler) Check2FaForUser(rw http.ResponseWriter, r *http.Request, params map[string]string) {
	a.l.Printf("Handling Check2FaForUser Users ")
	var request dto.UsernameRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(rw, "Error decoding request", http.StatusBadRequest)
		return
	}
	policy := bluemonday.UGCPolicy()

	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))
	res, err := a.tfaService.Check2FaForUser(request.Username)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	response, _ := json.Marshal(res)
	rw.Write(response)

	rw.Header().Set("Content-Type", "application/json")

}

func (a AuthenticationHandler) Enable2FaForUser(rw http.ResponseWriter, r *http.Request, params map[string]string) {
	a.l.Printf("Handling Check2FaForUser Users ")
	var request dto.UsernameRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(rw, "Error decoding request", http.StatusBadRequest)
		return
	}
	policy := bluemonday.UGCPolicy()
	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))

	res, uri, _ := a.tfaService.Enable2FaForUser(request.Username)
	enable2FaResponse := dto.Enable2FaResponse{
		Res: res,
		Uri: uri,
	}

	response, _ := json.Marshal(enable2FaResponse)
	rw.Write(response)
	rw.Header().Set("Content-Type", "application/json")
}

func (a AuthenticationHandler) Disable2FaForUser(rw http.ResponseWriter, r *http.Request, params map[string]string) {
	a.l.Printf("Handling Disable2FaForUser Users ")
	var request dto.UsernameRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(rw, "Error decoding request", http.StatusBadRequest)
		return
	}
	policy := bluemonday.UGCPolicy()

	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))
	res, _ := a.tfaService.Disable2FaForUser(request.Username)

	response, _ := json.Marshal(res)
	rw.Write(response)

	rw.Header().Set("Content-Type", "application/json")
}

func (a AuthenticationHandler) Authenticate2Fa(rw http.ResponseWriter, r *http.Request, parameters map[string]string) {
	a.l.Printf("Handling Authenticate2Fa Users ")
	var request dto.AuthenticateRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(rw, "Error decoding request", http.StatusBadRequest)
		return
	}
	policy := bluemonday.UGCPolicy()

	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))
	userSecret, err := a.tfaService.GetUserSecret(request.Username)
	if err != nil {
		http.Error(rw, "Error user secret", http.StatusBadRequest)
		return
	}
	otpc := services.NewOTPConfig(userSecret)

	val, err := otpc.Authenticate(strconv.Itoa(request.Token))
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

func (a AuthenticationHandler) LoginUser(rw http.ResponseWriter, r *http.Request, params map[string]string) {
	a.l.Println("Handling LOGIN Users")

	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.userService.GetByUsername(context.TODO(), loginRequest.Username)
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

	userRoles, err := a.userService.GetUserRole(loginRequest.Username)
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

func (a AuthenticationHandler) AuthenticateUser(rw http.ResponseWriter, r *http.Request, params map[string]string) {
	a.l.Println("Handling LOGIN Users")

	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.userService.GetByUsername(context.TODO(), loginRequest.Username)
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

	policy := bluemonday.UGCPolicy()

	loginRequest.Username = strings.TrimSpace(policy.Sanitize(loginRequest.Username))
	res, err := a.tfaService.Check2FaForUser(loginRequest.Username)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	response, _ := json.Marshal(res)
	rw.Write(response)

	rw.Header().Set("Content-Type", "application/json")

}
