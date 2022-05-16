package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"user/module/auth"
	"user/module/dto"
	"user/module/helpers"
	"user/module/model"
	"user/module/repository"
	"user/module/service"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

// UserHandler is a http.Handler
type UserHandler struct {
	l               *log.Logger
	service         *service.UserService
	registerService *service.RegisteredUserService
	jsonConv        *helpers.JsonConverters
	repo            *repository.UserRepository
	validator       *validator.Validate
	passwordUtil    *helpers.PasswordUtil
}

type ResponseEmail struct {
	Email string
}
type ErrorResponse struct {
	Err string
}

func NewUserHandler(l *log.Logger, service *service.UserService, registerService *service.RegisteredUserService, jsonConv *helpers.JsonConverters, repo *repository.UserRepository, validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil) *UserHandler {
	return &UserHandler{l, service, registerService, jsonConv, repo, validator, passwordUtil}
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

//create registered user function
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
		//	rw.WriteHeader(http.StatusBadRequest)
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest) //400
	}

	if err := u.validator.Struct(&newUser); err != nil {
		http.Error(rw, "New user dto fields aren't entered in valid format! error:"+err.Error(), http.StatusExpectationFailed) //400

	}

	var er error
	_, er = u.service.GetByUsername(context.TODO(), newUser.Username)
	if er != nil {
		http.Error(rw, "User with entered username already exists! error:"+err.Error(), http.StatusConflict) //409
	}

	salt := ""
	password := ""
	validPassword := u.passwordUtil.IsValidPassword(newUser.Password)

	if validPassword {
		//PASSWORD SALT
		salt, password = u.passwordUtil.GeneratePasswordWithSalt(newUser.Password)

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
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(rw).Encode(err)
	}

	newUser.Password = string(pass)

	//zakucana rola za sad
	//	var role = "REGISTERED_USER"
	//var salt = ""
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, newUser.DateOfBirth)
	email, er := u.registerService.CreateRegisteredUser(newUser.Username, password, newUser.Email, newUser.PhoneNumber, newUser.FirstName, newUser.LastName, gender, model.REGISTERED_USER, salt, dateOfBirth)

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
