package handlers

import (
	"bytes"
	common "common/module"
	"common/module/logger"
	pb "common/module/proto/user_service"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	hibp "github.com/mattevans/pwned-passwords"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"user/module/application/helpers"
	"user/module/application/services"
	"user/module/infrastructure/api"
)

type UserHandler struct {
	l            *logger.Logger
	service      *services.UserService
	jsonConv     *helpers.JsonConverters
	validator    *validator.Validate
	passwordUtil *helpers.PasswordUtil
	pwnedClient  *hibp.Client
	tokenService *services.ApiTokenService
}

func (u UserHandler) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (u UserHandler) MustEmbedUnimplementedUserServiceServer() {
	//u.l.Println("Handling MustEmbedUnimplementedUserServiceServer Users")
	u.l.Logger.Infof("Handling MustEmbedUnimplementedUserServiceServer Users")
}

func NewUserHandler(l *logger.Logger, service *services.UserService, jsonConv *helpers.JsonConverters, validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil, pwnedClient *hibp.Client, tokenService *services.ApiTokenService) *UserHandler {
	return &UserHandler{l, service, jsonConv, validator, passwordUtil, pwnedClient, tokenService}
}

func (u UserHandler) GenerateAPIToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.ApiToken, error) {
	u.l.Logger.Println("GenerateAPIToken")
	username := request.Username.Username
	policy := bluemonday.UGCPolicy()
	username = strings.TrimSpace(policy.Sanitize(username))
	sqlInj := common.CheckRegexSQL(username)
	if username == "" {
		u.l.Logger.WithFields(logrus.Fields{
			"user": request.Username.Username,
		}).Warnf("XSS")
		return &pb.ApiToken{ApiToken: ""}, errors.New("fields are empty or xss happened")
	} else if sqlInj {
		u.l.Logger.WithFields(logrus.Fields{
			"user": request.Username.Username,
		}).Warnf("INJECTION")
		return &pb.ApiToken{ApiToken: ""}, errors.New("there is chance of sql injection happening")
	} else {
		u.l.Logger.WithFields(logrus.Fields{
			"user": request.Username.Username,
		}).Infof("GenerateAPIToken")
	}
	existsErr := u.service.UserExists(username)
	if existsErr != nil {
		u.l.Logger.WithFields(logrus.Fields{
			"user": request.Username.Username,
		}).Infof("USER DO NOT EXIST")
		return &pb.ApiToken{ApiToken: ""}, existsErr
	}
	user, er := u.service.GetByUsername(context.TODO(), username)
	if er != nil {
		u.l.Logger.WithFields(logrus.Fields{
			"user": request.Username.Username,
		}).Error("DB ERR")
		return &pb.ApiToken{ApiToken: ""}, er
	}
	token, tokenErr := u.tokenService.GenerateApiToken(user)
	if tokenErr != nil {
		u.l.Logger.WithFields(logrus.Fields{
			"user": request.Username.Username,
		}).Infof("GEN TOKEN ERR")
		return &pb.ApiToken{ApiToken: ""}, tokenErr
	}
	u.l.Logger.WithFields(logrus.Fields{
		"user": request.Username.Username,
	}).Infof("SUCCESS GenerateAPIToken")
	return &pb.ApiToken{ApiToken: token}, nil
}

func (u UserHandler) ShareJobOffer(ctx context.Context, request *pb.ShareJobOfferRequest) (*pb.EmptyRequest, error) {
	token := request.ShareJobOffer.ApiToken
	hasAccess, er := u.tokenService.CheckIfHasAccess(token)
	if er != nil {
		return &pb.EmptyRequest{}, er
	}
	if !hasAccess {
		return &pb.EmptyRequest{}, errors.New("you don't have acccess")
	}

	postBody, _ := json.Marshal(map[string]any{
		"Publisher":      request.ShareJobOffer.JobOffer.Publisher,
		"Position":       request.ShareJobOffer.JobOffer.Position,
		"JobDescription": request.ShareJobOffer.JobOffer.JobDescription,
		"Requirements":   request.ShareJobOffer.JobOffer.Requirements,
		"DatePosted":     request.ShareJobOffer.JobOffer.DatePosted,
		"Duration":       request.ShareJobOffer.JobOffer.Duration,
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:9090/job_offer", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)

	return &pb.EmptyRequest{}, nil
}

type myJSON struct {
	Array []string
}

func (u UserHandler) ActivateUserAccount(ctx context.Context, request *pb.ActivationRequest) (*pb.ActivationResponse, error) {
	u.l.Logger.Println("Handling ActivateUserAccount ")
	//TODO:mzd dodati provjeru da li se uspelo ok mapirati?
	requstDto := api.MapPbToUserActivateRequest(request)
	err := u.validator.Struct(requstDto)
	if err != nil {
		u.l.Logger.Warnf("XSS")
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, err
	}
	policy := bluemonday.UGCPolicy()
	//sanitize everything
	requstDto.Username = strings.TrimSpace(policy.Sanitize(requstDto.Username))
	requstDto.Code = strings.TrimSpace(policy.Sanitize(requstDto.Code))
	if requstDto.Username == "" || requstDto.Code == "" {
		u.l.Logger.Warnf("XSS")
		//http.Error(rw, "Fields are empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, errors.New("fields are empty or xss happened")
	}
	existsErr := u.service.UserExists(requstDto.Username)
	if existsErr != nil {
		u.l.Logger.Warnf("USER DO NOT EXIST")
		//http.Error(rw, "User with entered username already exists!", http.StatusConflict) //409
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, existsErr
	}
	sqlInj := common.CheckForSQLInjection([]string{requstDto.Username, requstDto.Code})
	if sqlInj {
		u.l.Logger.Warnf("INJECTION")
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, errors.New("there is chance of sql injection happening")
	}
	var code int
	code, convertError := strconv.Atoi(requstDto.Code)
	if convertError != nil {
		u.l.Logger.Warnf("CONVERT ERR")
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, convertError
	}
	activated, e := u.service.ActivateUserAccount(requstDto.Username, code)
	if e != nil {
		u.l.Logger.Warnf("ERR")
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, e
	}
	if !activated {
		u.l.Logger.Warnf("account activation failed")
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, errors.New("account activation failed")
	}
	u.l.Logger.Infof("SUCCESS GenerateAPIToken")
	return &pb.ActivationResponse{Activated: activated, Username: requstDto.Username}, nil
}

func (u UserHandler) GetAll(ctx context.Context, request *pb.EmptyRequest) (*pb.GetAllResponse, error) {
	u.l.Logger.Println("Handling GetAll Users")
	users, err := u.service.GetUsers()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Users: []*pb.User{},
	}
	for _, user := range users {
		current := api.MapProduct(&user)
		response.Users = append(response.Users, current)
	}
	return response, nil
}

func (u UserHandler) UpdateUser(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateUserResponse, error) {
	u.l.Logger.Println("Handling UpdateUser")

	return &pb.UpdateUserResponse{UpdatedUser: nil}, nil
}

func (u UserHandler) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	u.l.Logger.Println("Handling RegisterUser")
	//TODO:mzd dodati provjeru da li se uspelo ok mapirati?
	newUser := api.MapPbUserToNewUserDto(request)
	//TODO:dodati validaciju u obliku regexa, spreciti injection napad
	if err := u.validator.Struct(newUser); err != nil {
		u.l.Logger.Println("Invalid values")
		return nil, err
		//http.Error(rw, "New user dto fields aren't entered in valid format! error:"+err.Error(), http.StatusExpectationFailed) //400
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
		u.l.Logger.Warnf("XSS")
		return nil, errors.New("fields are empty or xss happened")
	}
	sqlInj := common.CheckForSQLInjection([]string{newUser.Username, newUser.FirstName, newUser.LastName, newUser.Email,
		newUser.Gender, newUser.DateOfBirth, newUser.PhoneNumber, newUser.RecoveryEmail})
	if sqlInj {
		u.l.Logger.Warnf("INJECTION")
		return nil, errors.New("there is chance for sql injection")
	}
	err := u.service.UserExists(newUser.Username)
	if err == nil {
		u.l.Logger.Infof("USER EXISTS")
		return nil, err
	}

	var hashedSaltedPassword = ""
	validPassword := u.passwordUtil.IsValidPassword(newUser.Password)

	if validPassword {
		pass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		hashedSaltedPassword = string(pass)
	} else {
		u.l.Logger.Infof("PASSWORD FORMAT NOT VALID")
		return nil, errors.New("password format is not valid")
	}
	newUser.Password = hashedSaltedPassword
	registeredUser, er := u.service.CreateRegisteredUser(api.MapDtoToUser(newUser))

	if er != nil {
		u.l.Logger.Infof(er.Error())
		return nil, er
	}

	u.l.Logger.Infof("SUCCESS RegisterUser")
	return &pb.RegisterUserResponse{RegisteredUser: api.MapUserToPbResponseUser(registeredUser)}, nil
}

func (u UserHandler) SendRequestForPasswordRecovery(ctx context.Context, request *pb.PasswordRecoveryRequest) (*pb.PasswordRecoveryResponse, error) {

	u.l.Logger.Println("Handling PASSWORD RECCOVERY ")
	//TODO: injection
	var requestUsername = request.Username.Username
	policy := bluemonday.UGCPolicy()
	requestUsername = strings.TrimSpace(policy.Sanitize(requestUsername))

	if requestUsername == "" {
		fmt.Println("usernname empty or xss")
		//http.Error(rw, "Field empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return &pb.PasswordRecoveryResponse{CodeSent: false}, errors.New("fields are empty or xss happened")
	}
	sqlInj := common.CheckRegexSQL(requestUsername)
	if sqlInj {
		u.l.Logger.Warnf("INJECTION")
		return nil, errors.New("there is chance for sql injection")
	}
	existsErr := u.service.UserExists(requestUsername)
	if existsErr != nil {
		u.l.Logger.Warnf("USER DOES NOT EXIST")
		return &pb.PasswordRecoveryResponse{CodeSent: false}, existsErr
	}

	codeSent, codeErr := u.service.SendCodeToRecoveryMail(requestUsername)
	if codeErr != nil {
		u.l.Logger.Warnf(codeErr.Error())
		return &pb.PasswordRecoveryResponse{CodeSent: false}, codeErr
	}
	if !codeSent {
		u.l.Logger.Warnf("ACCOUNT ACTIVATION FAILED")
		return &pb.PasswordRecoveryResponse{CodeSent: false}, errors.New("account activation failed")
	}
	u.l.Logger.Infof("SUCCESS SendRequestForPasswordRecovery")

	return &pb.PasswordRecoveryResponse{CodeSent: true}, nil
}

func (u UserHandler) RecoverPassword(ctx context.Context, request *pb.NewPasswordRequest) (*pb.NewPasswordResponse, error) {

	u.l.Logger.Println("Handling RecoverPassword handler ")
	//TODO:mzd dodati provjeru da li se uspelo ok mapirati?
	requestDto := api.MapPbToNewPasswordRequestDto(request)

	err := u.validator.Struct(requestDto)
	if err != nil {
		u.l.Logger.Warnf("Invalid values")
		return &pb.NewPasswordResponse{PasswordChanged: false}, err
		//http.Error(rw, "New user dto fields aren't entered in valid format! error:"+err.Error(), http.StatusExpectationFailed) //400
	}
	policy := bluemonday.UGCPolicy()
	//sanitize everything
	requestDto.Username = strings.TrimSpace(policy.Sanitize(requestDto.Username))
	requestDto.Code = strings.TrimSpace(policy.Sanitize(requestDto.Code))
	requestDto.NewPassword = strings.TrimSpace(policy.Sanitize(requestDto.NewPassword))
	if requestDto.Username == "" || requestDto.Code == "" || requestDto.NewPassword == "" {
		u.l.Logger.Warnf("XSS")
		return &pb.NewPasswordResponse{PasswordChanged: false}, errors.New("fields are empty or xss happened")
	}
	sqlInj := common.CheckForSQLInjection([]string{requestDto.Username, requestDto.Code})
	if sqlInj {
		u.l.Logger.Warnf("INJECTION")
		return nil, errors.New("there is chance for sql injection")
	}

	existsErr := u.service.UserExists(requestDto.Username)
	if existsErr != nil {
		u.l.Logger.Infof("USER DOES NOT EXIST")
		return &pb.NewPasswordResponse{PasswordChanged: false}, existsErr
	}
	///////////////////
	var hashedSaltedPassword = ""
	validPassword := u.passwordUtil.IsValidPassword(requestDto.NewPassword)

	if validPassword {
		pass, err := bcrypt.GenerateFromPassword([]byte(requestDto.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			return &pb.NewPasswordResponse{PasswordChanged: false}, err
		}

		hashedSaltedPassword = string(pass)

	} else {
		fmt.Println("Password format is not valid!")
		//http.Error(rw, "Password format is not valid! error:"+err.Error(), http.StatusBadRequest) //400
		//return
		return &pb.NewPasswordResponse{PasswordChanged: false}, errors.New("password format is not valid")
	}
	passChanged, err := u.service.CreateNewPassword(requestDto.Username, hashedSaltedPassword, requestDto.Code)
	if err != nil {
		//http.Error(rw, err.Error(), http.StatusConflict) //409
		return &pb.NewPasswordResponse{PasswordChanged: false}, err
	}
	if !passChanged {
		//http.Error(rw, "error changing password", http.StatusConflict) //409
		return &pb.NewPasswordResponse{PasswordChanged: false}, errors.New("error changing password")
	}
	return &pb.NewPasswordResponse{PasswordChanged: true}, nil
}

func (u UserHandler) PwnedPassword(ctx context.Context, request *pb.PwnedRequest) (*pb.PwnedResponse, error) {

	u.l.Logger.Println("Handling PWNED PASSWORD")
	pwnedPassword := request.Password.Password
	policy := bluemonday.UGCPolicy()
	//sanitize everything
	pwnedPassword = strings.TrimSpace(policy.Sanitize(pwnedPassword))
	if pwnedPassword == "" {
		u.l.Logger.Warnf("XSS")
		//http.Error(rw, "Fields are empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return &pb.PwnedResponse{Pwned: true, Message: "fields are empty or xss happened"}, errors.New("fields are empty or xss happened")
	}
	sqlInj := common.CheckRegexSQL(pwnedPassword)
	if sqlInj {
		u.l.Logger.Warnf("INJECTION")
		return nil, errors.New("there is chance for sql injection")
	}

	pwned, err := u.pwnedClient.Compromised(pwnedPassword)
	if err != nil {
		u.l.Logger.Infof("pwnedClient ERR")
		return &pb.PwnedResponse{Pwned: pwned, Message: "error checking if password is pwned"}, errors.New("error checkinf if password is pwaned")
	}
	var mess string
	if pwned {
		// Oh dear! 😱 -- You should avoid using that password
		fmt.Print("Found to be compromised")
		mess = "Password is pwned,please chose another one!"
	} else {
		mess = "Password is OK!"
	}
	u.l.Logger.Infof("SUCCESS PwnedPassword")
	return &pb.PwnedResponse{Pwned: pwned, Message: mess}, nil

}
