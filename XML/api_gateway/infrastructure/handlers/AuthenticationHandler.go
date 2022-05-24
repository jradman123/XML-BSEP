package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	myerr "gateway/module/application/errors"
	"gateway/module/application/helpers"
	"gateway/module/application/services"
	"gateway/module/auth"
	"gateway/module/domain/dto"
	"gateway/module/domain/model"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
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
	err = mux.HandlePath("POST", "/users/register/user", a.RegisterUser)
	if err != nil {
		panic(err)
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
		http.Error(rw, "User not found! "+err.Error(), http.StatusBadRequest)
		return

	}

	salt, err := a.service.GetUserSalt(loginRequest.Username)
	a.l.Printf("so:" + salt)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.passwordUtil.ValidateLoginPassword(salt, user.Password, loginRequest.Password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	var claims = &auth.JwtClaims{}
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

	logInResponse := dto.LogInResponseDto{
		Token: token,
	}

	logInResponseJson, _ := json.Marshal(logInResponse)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(logInResponseJson)
}

func (a AuthenticationHandler) RegisterUser(rw http.ResponseWriter, r *http.Request, params map[string]string) {

	var newUser dto.NewUser
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest) //400
	}

	if err := a.validator.Struct(&newUser); err != nil {
		http.Error(rw, "New user dto fields aren't entered in valid format! error:"+err.Error(), http.StatusExpectationFailed) //400

	}

	err = a.service.UserExists(newUser.Username)
	if err == nil {
		http.Error(rw, "User with entered username already exists! error:"+err.Error(), http.StatusConflict) //409
	}

	salt := ""
	password := ""
	validPassword := a.passwordUtil.IsValidPassword(newUser.Password)

	if validPassword {
		//PASSWORD SALT
		salt, password = a.passwordUtil.GeneratePasswordWithSalt(newUser.Password)

	} else {
		http.Error(rw, "Password format is not valid! error:"+err.Error(), http.StatusBadRequest) //400
		return
	}

	gender := model.OTHER
	switch newUser.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	//cuvamo password kao hash neki
	pass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := myerr.AuthenticationError{StatusCode: 404, Err: err, Message: "Password Encryption  failed"}
		json.NewEncoder(rw).Encode(err)
	}

	newUser.Password = string(pass)

	//zakucana rola za sad
	//	var role = "REGISTERED_USER"
	//var salt = ""
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, newUser.DateOfBirth)
	email, er := a.service.CreateRegisteredUser(newUser.Username, password, newUser.Email, newUser.PhoneNumber, newUser.FirstName, newUser.LastName, gender, model.Regular, salt, dateOfBirth)

	if er != nil {
		http.Error(rw, "Failed creating registered user! error:"+er.Error(), http.StatusExpectationFailed) //
		return
	}
	userEmail := ResponseEmail{
		Email: email,
	}

	userEmailJson, _ := json.Marshal(userEmail)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(userEmailJson)
}

//??????????//
type ResponseEmail struct {
	Email string
}
