package handlers

import (
	common "common/module"
	"common/module/interceptor"
	"common/module/logger"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gateway/module/application/helpers"
	"gateway/module/application/services"
	"gateway/module/auth"
	"gateway/module/domain/dto"
	modelGateway "gateway/module/domain/model"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/microcosm-cc/bluemonday"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	tracer "monitoring/module"
	"net/http"
	"strconv"
	"strings"
)

type AuthenticationHandler struct {
	l                   *log.Logger
	logInfo             *logger.Logger
	logError            *logger.Logger
	userService         *services.UserService
	tfaService          *services.TFAuthService
	validator           *validator.Validate
	passwordUtil        *helpers.PasswordUtil
	passwordLessService *services.PasswordLessService
}

func NewAuthenticationHandler(l *log.Logger, logInfo *logger.Logger, logError *logger.Logger, userService *services.UserService,
	tfaService *services.TFAuthService,
	validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil, passwordLessService *services.PasswordLessService) Handler {
	return &AuthenticationHandler{l, logInfo, logError, userService, tfaService, validator, passwordUtil, passwordLessService}
}

func (a AuthenticationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/users/auth/user", a.AuthenticateUser)
	if err != nil {
		panic(err)
	}

	err = mux.HandlePath("POST", "/users/auth/user/regular", a.AuthenticateUserRegular)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/2fa/authenticate", a.Authenticate2Fa)
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

	err = mux.HandlePath("POST", "/users/login/passwordless", a.PasswordLessLoginReq)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/users/login/passwordless/{id}", a.PasswordlessLogin)
	if err != nil {
		panic(err)
	}

}

func (a AuthenticationHandler) Check2FaForUser(rw http.ResponseWriter, r *http.Request, _ map[string]string) {
	span := tracer.StartSpanFromRequest("check2FaForUser", otgo.GlobalTracer(), r)
	defer span.Finish()

	ctx := tracer.ContextWithSpan(context.Background(), span)
	a.l.Printf("Handling Check2FaForUser Users ")
	var request dto.UsernameRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(rw, "Error decoding request", http.StatusBadRequest)
		return
	}
	policy := bluemonday.UGCPolicy()

	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))
	res, err := a.tfaService.Check2FaForUser(request.Username, ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	response, _ := json.Marshal(res)
	_, err = rw.Write(response)
	if err != nil {
		return
	}

	rw.Header().Set("Content-Type", "application/json")

}

func (a AuthenticationHandler) Enable2FaForUser(rw http.ResponseWriter, r *http.Request, _ map[string]string) {
	span := tracer.StartSpanFromRequest("enable2FaForUser", otgo.GlobalTracer(), r)
	defer span.Finish()

	ctx := tracer.ContextWithSpan(context.Background(), span)
	a.l.Printf("Handling Check2FaForUser Users ")
	var request dto.UsernameRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(rw, "Error decoding request", http.StatusBadRequest)
		return
	}
	policy := bluemonday.UGCPolicy()
	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))

	res, uri, _ := a.tfaService.Enable2FaForUser(request.Username, ctx)
	enable2FaResponse := dto.Enable2FaResponse{
		Res: res,
		Uri: uri,
	}

	response, _ := json.Marshal(enable2FaResponse)
	_, err = rw.Write(response)
	if err != nil {
		return
	}
	rw.Header().Set("Content-Type", "application/json")
}

func (a AuthenticationHandler) Disable2FaForUser(rw http.ResponseWriter, r *http.Request, _ map[string]string) {
	span := tracer.StartSpanFromRequest("disable2FaForUser", otgo.GlobalTracer(), r)
	defer span.Finish()

	ctx := tracer.ContextWithSpan(context.Background(), span)
	a.l.Printf("Handling Disable2FaForUser Users ")
	var request dto.UsernameRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(rw, "Error decoding request", http.StatusBadRequest)
		return
	}
	policy := bluemonday.UGCPolicy()

	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))
	res, _ := a.tfaService.Disable2FaForUser(request.Username, ctx)

	response, _ := json.Marshal(res)
	_, err = rw.Write(response)
	if err != nil {
		return
	}

	rw.Header().Set("Content-Type", "application/json")
}

func (a AuthenticationHandler) AuthenticateUser(rw http.ResponseWriter, r *http.Request, _ map[string]string) {
	span := tracer.StartSpanFromRequest("authenticateUser", otgo.GlobalTracer(), r)
	defer span.Finish()

	ctx := tracer.ContextWithSpan(context.Background(), span)
	a.l.Println("Handling AuthenticateUser Users")

	var loginRequest dto.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest)
		return
	}
	ip := ReadUserIP(r)
	err = CheckForAttack(loginRequest, ip, a)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	a.logInfo.Logger.WithFields(logrus.Fields{
		"user":   loginRequest.Username,
		"userIP": ip,
	}).Infof("INFO:Handling LOGIN")

	user, err := a.userService.GetByUsername(ctx, loginRequest.Username)
	if err != nil {
		http.Error(rw, "Invalid credentials!", http.StatusBadRequest)
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:USER NOT FOUND")
		http.Error(rw, "User not found! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !user.IsConfirmed {
		http.Error(rw, "Check your mail for activation code! ", http.StatusBadRequest)
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:USER NOT ACTIVATED")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		http.Error(rw, "Invalid credentials!", http.StatusBadRequest)
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:INCORRECT PASSWORD")
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	policy := bluemonday.UGCPolicy()

	loginRequest.Username = strings.TrimSpace(policy.Sanitize(loginRequest.Username))
	twofa, err := a.tfaService.Check2FaForUser(loginRequest.Username, ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	res := dto.AuthenticateResponse{
		Username: loginRequest.Username,
		TwoFa:    twofa,
	}
	response, _ := json.Marshal(res)
	_, err = rw.Write(response)
	if err != nil {
		return
	}

	rw.Header().Set("Content-Type", "application/json")
}

func (a AuthenticationHandler) Authenticate2Fa(rw http.ResponseWriter, r *http.Request, _ map[string]string) {
	span := tracer.StartSpanFromRequest("authenticate2Fa", otgo.GlobalTracer(), r)
	defer span.Finish()

	ctx := tracer.ContextWithSpan(context.Background(), span)
	a.l.Printf("Handling Authenticate2Fa Users ")
	var request dto.AuthenticateRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(rw, "Error decoding request", http.StatusBadRequest)
		return
	}
	policy := bluemonday.UGCPolicy()

	request.Username = strings.TrimSpace(policy.Sanitize(request.Username))
	userSecret, err := a.tfaService.GetUserSecret(request.Username, ctx)
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

	var claims = &interceptor.JwtClaims{}
	claims.Username = request.Username

	userRoles, err := a.userService.GetUserRole(request.Username, ctx)
	ip := ReadUserIP(r)
	if err != nil {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   request.Username,
			"userIP": ip,
		}).Errorf("ERR:THIS USER HAS NO ROLE")
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return

	}

	claims.Roles = append(claims.Roles, userRoles)

	token, expirationTime, err := auth.GenerateToken(claims)

	if err != nil {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   request.Username,
			"userIP": ip,
		}).Errorf("ERR:GENERATING TOKEN")
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return

	}
	user, err := a.userService.GetByUsername(ctx, request.Username)
	var roleString string

	if user.Role == modelGateway.Admin {
		roleString = "Admin"
	} else if user.Role == modelGateway.Agent {
		roleString = "Agent"
	} else if user.Role == modelGateway.Regular {
		roleString = "Regular"
	}

	logInResponse := dto.LogInResponseDto{
		Token:          token,
		Role:           roleString,
		Email:          user.Email,
		Username:       user.Username,
		ExpirationTime: expirationTime,
	}

	logInResponseJson, _ := json.Marshal(logInResponse)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(logInResponseJson)
	if err != nil {
		return
	}
}

func (a AuthenticationHandler) AuthenticateUserRegular(rw http.ResponseWriter, r *http.Request, _ map[string]string) {
	span := tracer.StartSpanFromRequest("authenticateUserRegular", otgo.GlobalTracer(), r)
	defer span.Finish()

	ctx := tracer.ContextWithSpan(context.Background(), span)
	a.l.Println("Handling AuthenticateUserRegular Users")

	var loginRequest dto.UsernameRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest)
		return
	}
	policy := bluemonday.UGCPolicy()
	loginRequest.Username = strings.TrimSpace(policy.Sanitize(loginRequest.Username))

	var claims = &interceptor.JwtClaims{}
	claims.Username = loginRequest.Username

	userRoles, err := a.userService.GetUserRole(loginRequest.Username, ctx)
	ip := ReadUserIP(r)
	if err != nil {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:THIS USER HAS NO ROLE")
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return

	}

	claims.Roles = append(claims.Roles, userRoles)
	token, expirationTime, err := auth.GenerateToken(claims)

	if err != nil {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:GENERATING TOKEN")
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return

	}
	user, err := a.userService.GetByUsername(ctx, loginRequest.Username)
	var roleString string

	if user.Role == modelGateway.Admin {
		roleString = "Admin"
	} else if user.Role == modelGateway.Agent {
		roleString = "Agent"
	} else if user.Role == modelGateway.Regular {
		roleString = "Regular"
	}

	logInResponse := dto.LogInResponseDto{
		Token:          token,
		Role:           roleString,
		Email:          user.Email,
		Username:       user.Username,
		ExpirationTime: expirationTime,
	}

	logInResponseJson, _ := json.Marshal(logInResponse)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(logInResponseJson)
	if err != nil {
		return
	}
}

func (a AuthenticationHandler) PasswordLessLoginReq(rw http.ResponseWriter, r *http.Request, _ map[string]string) {
	span := tracer.StartSpanFromRequest("passwordLessLoginReq", otgo.GlobalTracer(), r)
	defer span.Finish()

	ctx := tracer.ContextWithSpan(context.Background(), span)
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

	user, err := a.userService.GetByUsername(ctx, loginRequest.Username)
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
	err = a.passwordLessService.SendLink(ctx, "https://localhost:4200", "http://localhost:9090/", user)
	if err != nil {
		return
	}

	rw.WriteHeader(http.StatusNoContent)

}

func (a AuthenticationHandler) PasswordlessLogin(rw http.ResponseWriter, r *http.Request, _ map[string]string) {
	span := tracer.StartSpanFromRequest("passwordLessLogin", otgo.GlobalTracer(), r)
	defer span.Finish()

	ctx := tracer.ContextWithSpan(context.Background(), span)
	fmt.Println("Handling PasswordlessLogin Request")
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
		a.LogError(ip, "", "XSS")
		http.Error(rw, "XSS! ", http.StatusBadRequest)
		return
	} else if sqlInj {
		a.LogError(ip, "", "BAD VALIDATION: POSSIBLE INJECTION")
		http.Error(rw, "Chance for sql injection! ", http.StatusBadRequest)
		return
	} else {
		a.LogInfo(ip, "Handling PasswordlessLogin")

	}

	ver, err := a.passwordLessService.GetUsernameByCode(code, ctx)
	username := ver.Username
	if err != nil {
		http.Error(rw, "code doesn't exist ot is invalid", http.StatusBadRequest)
		return
	}

	user, err := a.userService.GetByUsername(ctx, username)
	if err != nil {
		a.LogError(ip, user.Username, "USER NOT FOUND")
		http.Error(rw, "User not found! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !user.IsConfirmed {
		a.LogError(ip, user.Username, "USER NOT ACTIVATED")
		http.Error(rw, "User account not activated! ", http.StatusBadRequest)
		return
	}
	validCode, err := a.passwordLessService.PasswordlessLogin(ver, ctx)
	if !validCode {
		a.LogError(ip, user.Username, "CODE INVALID")
		http.Error(rw, "Code invalid! ", http.StatusBadRequest)
		return
	}
	if err != nil {
		a.LogError(ip, user.Username, "CORE FLAG CHANGE")
		http.Error(rw, "Login code error! ", http.StatusBadRequest)
		return
	}

	var claims = &interceptor.JwtClaims{}
	claims.Username = username

	userRoles, err := a.userService.GetUserRole(username, ctx)
	if err != nil {
		a.LogError(ip, user.Username, "THIS USER HAS NO ROLE")
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return

	}

	claims.Roles = append(claims.Roles, userRoles)

	token, expirationTime, err := auth.GenerateToken(claims)

	if err != nil {
		a.LogError(ip, user.Username, "GENERATING TOKEN")
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
		Token:          token,
		Role:           roleString,
		Email:          user.Email,
		Username:       user.Username,
		ExpirationTime: expirationTime,
	}

	logInResponseJson, _ := json.Marshal(logInResponse)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write(logInResponseJson)
	if err != nil {
		return
	}
}

func (a AuthenticationHandler) LogError(ip string, username string, message string) {
	a.logError.Logger.WithFields(logrus.Fields{
		"user":   username,
		"userIP": ip,
	}).Errorf(message)

}
func (a AuthenticationHandler) LogInfo(ip string, message string) {
	a.logInfo.Logger.WithFields(logrus.Fields{
		"userIP": ip,
	}).Infof(message)

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

func CheckForAttack(loginRequest dto.LoginRequest, ip string, a AuthenticationHandler) error {

	policy := bluemonday.UGCPolicy()
	loginRequest.Password = strings.TrimSpace(policy.Sanitize(loginRequest.Password))
	loginRequest.Username = strings.TrimSpace(policy.Sanitize(loginRequest.Username))

	sqlInj := common.BadUsername(loginRequest.Username)
	sqlInj2 := common.BadPassword(loginRequest.Password)
	if loginRequest.Username == "" || loginRequest.Password == "" {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:XSS")

		return errors.New("XSS! ")

	} else if sqlInj || sqlInj2 {
		a.logError.Logger.WithFields(logrus.Fields{
			"user":   loginRequest.Username,
			"userIP": ip,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")

		return errors.New("Chance for sql injection! ")

	}
	return nil
}
