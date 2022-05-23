package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

}

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

type PassRecoveryForUser struct {
	Username string
}

type ActivateAccountWithCodeRequest struct {
	Username string
	Code     string
}
type ResponseEmail struct {
	Email string
}
type ErrorResponse struct {
	Err string
}
type ActivationResponse struct {
	Username  string
	Activated bool
}

func NewUserHandler(l *log.Logger, service *service.UserService, registerService *service.RegisteredUserService, jsonConv *helpers.JsonConverters, repo *repository.UserRepository, validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil, pwnedClient *hibp.Client) *UserHandler {
	return &UserHandler{l, service, registerService, jsonConv, repo, validator, passwordUtil, pwnedClient}
}

func (u *UserHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	enableCors(&rw)
	u.l.Println("Handling GET Users")
	users, err := u.service.GetUsers()
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	logInResponseJson, _ := json.Marshal(users)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(logInResponseJson)
}

func (u *UserHandler) CreateNewPassword(rw http.ResponseWriter, r *http.Request) {
	enableCors(&rw)
	u.l.Println("Handling PASSWORD RECCOVERY ")

	var requestBody dto.NewRecoveryPasword
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "Error decoding request:"+err.Error(), http.StatusBadRequest) //400
	}
	policy := bluemonday.UGCPolicy()
	requestBody.Username = strings.TrimSpace(policy.Sanitize(requestBody.Username))
	requestBody.NewPassword = strings.TrimSpace(policy.Sanitize(requestBody.NewPassword))
	requestBody.Code = strings.TrimSpace(policy.Sanitize(requestBody.Code))

	if requestBody.Username == "" || requestBody.NewPassword == "" || requestBody.Code == "" {
		fmt.Println("usrnname empty or xss")
		http.Error(rw, "Field empty or xss attack happened! error:", http.StatusExpectationFailed) //400
		return
	}
	var exists = u.registerService.UsernameExists(requestBody.Username)
	if !exists {
		fmt.Println(exists)
		http.Error(rw, "No user with such username!", http.StatusConflict) //409
	}

	///////////////////////////
	var hashedSaltedPassword = ""
	validPassword := u.passwordUtil.IsValidPassword(requestBody.NewPassword)

	if validPassword {
		//PASSWORD SALT
		//salt, password = u.passwordUtil.GeneratePasswordWithSalt(newUser.Password)
		//cuvamo password kao hash neki
		pass, err := bcrypt.GenerateFromPassword([]byte(requestBody.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			err := ErrorResponse{
				Err: "Password Encryption  failed",
			}
			json.NewEncoder(rw).Encode(err)
		}

		hashedSaltedPassword = string(pass)

	} else {
		fmt.Println("Password format is not valid!")
		http.Error(rw, "Password format is not valid! error:"+err.Error(), http.StatusBadRequest) //400
		return
	}
	////////////////////

	passChanged, err := u.registerService.CreateNewPassword(requestBody.Username, hashedSaltedPassword, requestBody.Code)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusConflict) //409
	}
	if !passChanged {
		http.Error(rw, "error changing password", http.StatusConflict) //409
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	//redirect na login

}

func (u *UserHandler) RecoverPasswordRequest(rw http.ResponseWriter, r *http.Request) {
	u.l.Println("Handling PASSWORD RECCOVERY ")
	enableCors(&rw)
	var requestBody PassRecoveryForUser
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "Error decoding request:"+err.Error(), http.StatusBadRequest) //400
	}
	policy := bluemonday.UGCPolicy()
	requestBody.Username = strings.TrimSpace(policy.Sanitize(requestBody.Username))

	if requestBody.Username == "" {
		fmt.Println("usrnname empty or xss")
		http.Error(rw, "Field empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return
	}
	var exists = u.registerService.UsernameExists(requestBody.Username)
	if !exists {
		fmt.Println(exists)
		http.Error(rw, "No user with such username!", http.StatusConflict) //409
	}
	////////////////////

	codeSent := u.registerService.SendCodeToRecoveryMail(requestBody.Username)
	if !codeSent {
		http.Error(rw, "error sending code to recovery mail", http.StatusConflict) //409
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
}

func (u *UserHandler) ActivateUserAccount(rw http.ResponseWriter, req *http.Request) {
	u.l.Println("Handling ACTIVATING ACCOUNT POST ")
	enableCors(&rw)
	var requestBody ActivateAccountWithCodeRequest
	err := json.NewDecoder(req.Body).Decode(&requestBody)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "Error decoding request:"+err.Error(), http.StatusBadRequest) //400
	}
	policy := bluemonday.UGCPolicy()
	requestBody.Username = strings.TrimSpace(policy.Sanitize(requestBody.Username))
	requestBody.Code = strings.TrimSpace(policy.Sanitize(requestBody.Code))

	var code int
	code, convertError := strconv.Atoi(requestBody.Code)
	if convertError != nil {
		fmt.Println(convertError)
		http.Error(rw, "Error converting code from string to int! error:"+convertError.Error(), http.StatusConflict) //409
	}

	if requestBody.Username == "" {
		fmt.Println("usrnname empty or xss")
		http.Error(rw, "Field empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return
	}
	var exists = u.registerService.UsernameExists(requestBody.Username)
	if !exists {
		fmt.Println(exists)
		http.Error(rw, "No user with such username!", http.StatusConflict) //409
	}

	activated, e := u.registerService.ActivateUserAccount(requestBody.Username, code)
	if e != nil {
		fmt.Println(e)
		http.Error(rw, e.Error(), http.StatusConflict) //409
	}
	if !activated {
		fmt.Println("account activation failed")
		http.Error(rw, "Account activation failed!", http.StatusConflict) //409
	}

	activation := ActivationResponse{
		Username:  requestBody.Username,
		Activated: true,
	}

	activationJson, _ := json.Marshal(activation)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(activationJson)
}

//create registered user function
func (u *UserHandler) AddUsers(rw http.ResponseWriter, req *http.Request) {
	u.l.Println("Handling POST Users")
	enableCors(&rw)
	var newUser dto.NewUser
	err := json.NewDecoder(req.Body).Decode(&newUser)

	if err != nil {
		fmt.Println(err)
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest) //400
	}

	if err := u.validator.Struct(&newUser); err != nil {
		fmt.Println(err)
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
	newUser.RecoveryEmail = strings.TrimSpace(policy.Sanitize(newUser.RecoveryEmail))

	if newUser.Username == "" || newUser.FirstName == "" || newUser.LastName == "" ||
		newUser.Gender == "" || newUser.DateOfBirth == "" || newUser.PhoneNumber == "" ||
		newUser.Password == "" || newUser.Email == "" || newUser.RecoveryEmail == "" {
		fmt.Println("fields are empty or xss")
		http.Error(rw, "Fields are empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return
	}

	exists := u.registerService.UsernameExists(newUser.Username)
	if exists {
		fmt.Println(exists)
		http.Error(rw, "User with entered username already exists!", http.StatusConflict) //409
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
		fmt.Println("Password format is not valid!")
		http.Error(rw, "Password format is not valid! error:"+err.Error(), http.StatusBadRequest) //400
		return
	}

	// answer, err := bcrypt.GenerateFromPassword([]byte(newUser.HashedAnswer), bcrypt.DefaultCost)
	// if err != nil {
	// 	fmt.Println(err)
	// 	err := ErrorResponse{
	// 		Err: "Answer Encryption  failed",
	// 	}
	// 	json.NewEncoder(rw).Encode(err)
	// }

	// newUser.HashedAnswer = string(answer)

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
	email, er := u.registerService.CreateRegisteredUser(newUser.Username, hashedSaltedPassword, newUser.Email, newUser.PhoneNumber, newUser.FirstName, newUser.LastName, gender, model.REGISTERED_USER, dateOfBirth, newUser.RecoveryEmail)

	if er != nil {
		fmt.Println(er)
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
	enableCors(&rw)
	u.l.Println("Handling PUT Users")
}
func (u *UserHandler) LoginUser(rw http.ResponseWriter, r *http.Request) {
	//enableCors(&rw)
	//rw.Header().Set("Access-Control-Allow-Origin", "https://localhost:4200")
	//rw.Header().Set("Access-Control-Allow-Headers", "authentication, content-type")
	// rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	u.l.Println("Handling LOGIN Users")

	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "Error decoding loginRequest:"+err.Error(), http.StatusBadRequest)
		return

	}

	user, err := u.service.GetByUsername(context.TODO(), loginRequest.Username)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "User not found! "+err.Error(), http.StatusBadRequest)
		return

	}
	fmt.Println("NE MOGU VISEEEEE")

	fmt.Println(user.IsConfirmed)
	if !user.IsConfirmed {
		fmt.Println("NIIJE AKTIVIRAN NALOF")
		http.Error(rw, "User account not activated! ", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	var claims = &auth.JwtClaims{}
	claims.Username = loginRequest.Username

	userRoles, err := u.service.GetUserRole(loginRequest.Username)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return

	}

	claims.Roles = append(claims.Roles, userRoles)
	var tokenCreationTime = time.Now().UTC()
	var tokenExpirationTime = tokenCreationTime.Add(time.Duration(30) * time.Minute)

	token, err := auth.GenerateToken(claims, tokenExpirationTime)

	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return

	}

	var roleString string

	if user.Role == model.ADMIN {
		roleString = "ADMIN"
	} else if user.Role == model.AGENT {
		roleString = "AGENT"
	} else if user.Role == model.REGISTERED_USER {
		roleString = "REGISTERED_USER"
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

func (u *UserHandler) CheckIfPwned(rw http.ResponseWriter, r *http.Request) {
	enableCors(&rw)
	u.l.Println("Handling PWNED PASSWORD")
	var pwnedPassword dto.PwnedPasswordRequest
	err := json.NewDecoder(r.Body).Decode(&pwnedPassword)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "Error decoding PwnedPasswordRequest:"+err.Error(), http.StatusBadRequest)
		return

	}

	pwned, err := u.pwnedClient.Compromised(pwnedPassword.PwnedPassword)
	if err != nil {
		fmt.Println(err)
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
