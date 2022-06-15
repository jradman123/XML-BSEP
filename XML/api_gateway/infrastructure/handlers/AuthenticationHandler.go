package handlers

import (
	common "common/module"
	"common/module/interceptor"
	"common/module/logger"
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
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
	"time"
)

type AuthenticationHandler struct {
	logInfo      *logger.Logger
	logError     *logger.Logger
	service      *services.UserService
	validator    *validator.Validate
	passwordUtil *helpers.PasswordUtil
}

func NewAuthenticationHandler(logInfo *logger.Logger, logError *logger.Logger, service *services.UserService, validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil) Handler {
	return &AuthenticationHandler{logInfo, logError, service, validator, passwordUtil}
}

func (a AuthenticationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/users/login/user", a.LoginUser)
	if err != nil {
		panic(err)
	}
}

func (a AuthenticationHandler) LoginUser(rw http.ResponseWriter, r *http.Request, params map[string]string) {
	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest)
		return
	}
	ip := ReadUserIP(r)
	policy := bluemonday.UGCPolicy()
	loginRequest.Password = strings.TrimSpace(policy.Sanitize(loginRequest.Password))
	loginRequest.Username = strings.TrimSpace(policy.Sanitize(loginRequest.Username))
	sqlInj := common.CheckForSQLInjection([]string{loginRequest.Username, loginRequest.Password})
	if loginRequest.Username == "" || loginRequest.Password == "" {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:XSS")
		http.Error(rw, "XSS! "+err.Error(), http.StatusBadRequest)
		return
	} else if sqlInj {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:INJECTION")
		http.Error(rw, "Chance for sql injection! "+err.Error(), http.StatusBadRequest)
		return
	} else {
		a.logInfo.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Infof("INFO:Handling LOGIN")
	}

	user, err := a.service.GetByUsername(context.TODO(), loginRequest.Username)
	if err != nil {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:USER NOT FOUND")
		http.Error(rw, "User not found! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !user.IsConfirmed {
		fmt.Println("account not activated")
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:USER NOT ACTIVATED")
		http.Error(rw, "User account not activated! ", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		fmt.Println(err)
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:INCORRECT PASSWORD")
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	var claims = &interceptor.JwtClaims{}
	claims.Username = loginRequest.Username

	userRoles, err := a.service.GetUserRole(loginRequest.Username)
	if err != nil {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:THIS USER HAS NO ROLE")
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return

	}

	claims.Roles = append(claims.Roles, userRoles)
	var tokenCreationTime = time.Now().UTC()
	var tokenExpirationTime = tokenCreationTime.Add(time.Duration(30) * time.Minute)

	token, err := auth.GenerateToken(claims, tokenExpirationTime)

	if err != nil {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:GENERATING TOKEN")
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

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
