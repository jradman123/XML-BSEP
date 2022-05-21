package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"user/module/auth"
	"user/module/dto"
	"user/module/helpers"
	"user/module/model"
	"user/module/repository"
	"user/module/service"

	hibp "github.com/mattevans/pwned-passwords"
	"github.com/microcosm-cc/bluemonday"
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
	pwnedClient     *hibp.Client
}

type ResponseEmail struct {
	Email string
}
type ErrorResponse struct {
	Err string
}

func NewUserHandler(l *log.Logger, service *service.UserService, registerService *service.RegisteredUserService, jsonConv *helpers.JsonConverters, repo *repository.UserRepository, validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil, pwnedClient *hibp.Client) *UserHandler {
	return &UserHandler{l, service, registerService, jsonConv, repo, validator, passwordUtil, pwnedClient}
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
	u.l.Println("Handling POST Users")
	//TODO: Ask yourself a question and answer it

	var newUser dto.NewUser
	err := json.NewDecoder(req.Body).Decode(&newUser)

	if err != nil {
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest) //400
	}

	if err := u.validator.Struct(&newUser); err != nil {
		http.Error(rw, "New user dto fields aren't entered in valid format! error:"+err.Error(), http.StatusExpectationFailed) //400

	}

	policy := bluemonday.UGCPolicy()
	//sanitize everything
	newUser.Username = strings.TrimSpace(policy.Sanitize(newUser.Username))
	newUser.FirstName = strings.TrimSpace(policy.Sanitize(newUser.FirstName))
	newUser.LastName = strings.TrimSpace(policy.Sanitize(newUser.LastName))
	newUser.Email = strings.TrimSpace(policy.Sanitize(newUser.Email))
	newUser.Password = strings.TrimSpace(policy.Sanitize(newUser.Password))
	newUser.Gender = strings.TrimSpace(policy.Sanitize(newUser.Gender))
	newUser.DateOfBirth = strings.TrimSpace(policy.Sanitize(newUser.DateOfBirth))
	newUser.PhoneNumber = strings.TrimSpace(policy.Sanitize(newUser.PhoneNumber))

	if newUser.Username == "" || newUser.FirstName == "" || newUser.LastName == "" ||
		newUser.Gender == "" || newUser.DateOfBirth == "" || newUser.PhoneNumber == "" ||
		newUser.Password == "" || newUser.Email == "" {
		http.Error(rw, "Fields are empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return
	}

	var er error
	_, er = u.service.GetByUsername(context.TODO(), newUser.Username)
	if er == nil {
		http.Error(rw, "User with entered username already exists! error:"+err.Error(), http.StatusConflict) //409
	}

	var hashedSaltedPassword = ""
	validPassword := u.passwordUtil.IsValidPassword(newUser.Password)

	if validPassword {
		//PASSWORD SALT
		//salt, password = u.passwordUtil.GeneratePasswordWithSalt(newUser.Password)
		//cuvamo password kao hash neki
		pass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			err := ErrorResponse{
				Err: "Password Encryption  failed",
			}
			json.NewEncoder(rw).Encode(err)
		}

		hashedSaltedPassword = string(pass)

	} else {
		http.Error(rw, "Password format is not valid! error:"+err.Error(), http.StatusBadRequest) //400
		return
	}

	answer, err := bcrypt.GenerateFromPassword([]byte(newUser.HashedAnswer), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Answer Encryption  failed",
		}
		json.NewEncoder(rw).Encode(err)
	}

	newUser.HashedAnswer = string(answer)

	gender := model.OTHER
	switch newUser.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
	}

	//zakucana rola za sad
	//	var role = "REGISTERED_USER"
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, newUser.DateOfBirth)
	email, er := u.registerService.CreateRegisteredUser(newUser.Username, hashedSaltedPassword, newUser.Email, newUser.PhoneNumber, newUser.FirstName, newUser.LastName, gender, model.REGISTERED_USER, dateOfBirth, newUser.Question, newUser.HashedAnswer)

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
		return

	}

	user, err := u.service.GetByUsername(context.TODO(), loginRequest.Username)
	if err != nil {
		http.Error(rw, "User not found! "+err.Error(), http.StatusBadRequest)
		return

	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	var claims = &auth.JwtClaims{}
	claims.Username = loginRequest.Username

	userRoles, err := u.service.GetUserRole(loginRequest.Username)
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

func (u *UserHandler) CheckIfPwned(rw http.ResponseWriter, r *http.Request) {

	u.l.Println("Handling PWNED PASSWORD")
	var pwnedPassword dto.PwnedPasswordRequest
	err := json.NewDecoder(r.Body).Decode(&pwnedPassword)
	if err != nil {
		http.Error(rw, "Error decoding PwnedPasswordRequest:"+err.Error(), http.StatusBadRequest)
		return

	}

	pwned, err := u.pwnedClient.Compromised(pwnedPassword.PwnedPassword)
	if err != nil {
		http.Error(rw, "Error checkinf if password is pwaned!"+err.Error(), http.StatusBadRequest)
		u.l.Println(pwnedPassword.PwnedPassword)
	}

	pwnedResponse := dto.PwnedResponse{
		IsPwned: pwned,
		Message: "",
	}

	if pwned {
		// Oh dear! 😱 -- You should avoid using that password
		fmt.Print("Found to be compromised")
		pwnedResponse.Message = "Password is pwned,please chose another one!"
	} else {
		pwnedResponse.Message = "Password is OK!"
	}

	pwnedResponseJson, _ := json.Marshal(pwnedResponse)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(pwnedResponseJson)

}
