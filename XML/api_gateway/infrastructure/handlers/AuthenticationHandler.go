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
	logInfo             *logger.Logger
	logError            *logger.Logger
	service             *services.UserService
	validator           *validator.Validate
	passwordUtil        *helpers.PasswordUtil
	passwordlessService *services.PasswordLessService
}

func NewAuthenticationHandler(logInfo *logger.Logger, logError *logger.Logger, service *services.UserService, validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil, passwordlessService *services.PasswordLessService) Handler {
	return &AuthenticationHandler{logInfo, logError, service, validator, passwordUtil, passwordlessService}
}

func (a AuthenticationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/users/login/user", a.LoginUser)
	err2 := mux.HandlePath("POST", "/users/login/passwordless", a.PasswordLessLoginReq)
	err3 := mux.HandlePath("GET", "/users/login/passwordless/{id}", a.PasswordlessLogin)

	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
	if err3 != nil {
		panic(err3)
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
	//sqlInj := common.CheckForSQLInjection([]string{loginRequest.Username, loginRequest.Password})
	sqlInj := common.BadUsername(loginRequest.Username)
	sqlInj2 := common.BadPassword(loginRequest.Password)
	//fmt.Println(sqlInj2)
	fmt.Println("username")
	fmt.Println(sqlInj)
	if loginRequest.Username == "" || loginRequest.Password == "" {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:XSS")
		http.Error(rw, "XSS! ", http.StatusBadRequest)
		return
	} else if sqlInj || sqlInj2 {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
		http.Error(rw, "Chance for sql injection! ", http.StatusBadRequest)
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

func (a AuthenticationHandler) PasswordLessLoginReq(rw http.ResponseWriter, r *http.Request, params map[string]string) {
	var loginRequest dto.PasswordLessLoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest)
		return
	}
	ip := ReadUserIP(r)
	policy := bluemonday.UGCPolicy()
	loginRequest.Username = strings.TrimSpace(policy.Sanitize(loginRequest.Username))
	sqlInj := common.BadUsername(loginRequest.Username)

	if loginRequest.Username == "" {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:XSS")
		http.Error(rw, "XSS! ", http.StatusBadRequest)
		return
	} else if sqlInj {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
		http.Error(rw, "Chance for sql injection! ", http.StatusBadRequest)
		return
	} else {
		a.logInfo.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Infof("INFO:Handling PasswordLessLoginReq")
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
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:USER NOT ACTIVATED")
		fmt.Println("account not activated")
		http.Error(rw, "User account not activated! ", http.StatusBadRequest)
		return
	}
	a.passwordlessService.SendLink(context.TODO(), "https://localhost:4200", "http://localhost:9090/", user)

	rw.WriteHeader(http.StatusNoContent)

}

func (a AuthenticationHandler) PasswordlessLogin(rw http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Println("PasswordlessLogin")
	ip := ReadUserIP(r)
	var code string
	p := strings.Split(r.URL.Path, "/")
	if len(p) == 1 {
		http.Error(rw, "no path param", http.StatusBadRequest)
		return
	} else if len(p) > 1 {
		code = p[len(p)-1]
	}

	policy := bluemonday.UGCPolicy()
	code = strings.TrimSpace(policy.Sanitize(code))
	sqlInj := common.BadId(code)

	if code == "" {
		a.logError.Logger.WithFields(logrus.Fields{
			"userIP": ip,
		}).Errorf("ERR:XSS")
		http.Error(rw, "XSS! ", http.StatusBadRequest)
		return
	} else if sqlInj {
		a.logError.Logger.WithFields(logrus.Fields{
			"userIP": ip,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
		http.Error(rw, "Chance for sql injection! ", http.StatusBadRequest)
		return
	} else {
		a.logInfo.Logger.WithFields(logrus.Fields{
			"userIP": ip,
		}).Infof("INFO:Handling PasswordlessLogin")
	}

	ver, err := a.passwordlessService.GetUsernameByCode(code)
	username := ver.Username
	if err != nil {
		http.Error(rw, "code doesn't exist ot is invalid", http.StatusBadRequest)
		return
	}

	user, err := a.service.GetByUsername(context.TODO(), username)
	if err != nil {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   user.Username,
			"userIP": ip,
		}).Errorf("ERR:USER NOT FOUND")
		http.Error(rw, "User not found! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !user.IsConfirmed {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   user.Username,
			"userIP": ip,
		}).Errorf("ERR:USER NOT ACTIVATED")
		fmt.Println("account not activated")
		http.Error(rw, "User account not activated! ", http.StatusBadRequest)
		return
	}
	validCode, err := a.passwordlessService.PasswordlesLogin(ver)
	if !validCode {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   user.Username,
			"userIP": ip,
		}).Errorf("ERR:CODE INVALID")
		fmt.Println("code not valid")
		http.Error(rw, "Code invalid! ", http.StatusBadRequest)
		return
	}
	if err != nil {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   user.Username,
			"userIP": ip,
		}).Errorf("ERR:CORE FLAG CHANGE")
		fmt.Println("error while changing code flag for used")
		http.Error(rw, "Login code error! ", http.StatusBadRequest)
		return
	}

	var claims = &interceptor.JwtClaims{}
	claims.Username = username

	userRoles, err := a.service.GetUserRole(username)
	if err != nil {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   user.Username,
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
			"user":   user.Username,
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
