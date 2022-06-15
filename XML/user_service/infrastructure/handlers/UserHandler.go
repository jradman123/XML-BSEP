package handlers

import (
	"bytes"
	common "common/module"
	"common/module/interceptor"
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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	logInfo      *logger.Logger
	logError     *logger.Logger
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
	u.logInfo.Logger.Infof("Handling MustEmbedUnimplementedUserServiceServer Users")
}

func NewUserHandler(logInfo *logger.Logger, logError *logger.Logger, service *services.UserService, jsonConv *helpers.JsonConverters, validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil, pwnedClient *hibp.Client, tokenService *services.ApiTokenService) *UserHandler {
	return &UserHandler{logInfo, logError, service, jsonConv, validator, passwordUtil, pwnedClient, tokenService}
}

func (u UserHandler) GenerateAPIToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.ApiToken, error) {
	username := request.Username.Username
	policy := bluemonday.UGCPolicy()
	username = strings.TrimSpace(policy.Sanitize(username))
	//sqlInj := common.CheckRegexSQL(username)
	sqlInj := common.BadUsername(username)
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	if username == "" {
		u.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:XSS")
		return nil, status.Error(codes.FailedPrecondition, "fields are empty or xss happened")
	} else if sqlInj {
		u.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
		return nil, status.Error(codes.FailedPrecondition, "there is chance of sql injection happening")
	} else {
		u.logInfo.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Infof("INFO:GenerateAPIToken")
	}
	existsErr := u.service.UserExists(username)
	if existsErr != nil {
		u.logError.Logger.Errorf("ERR:USER DOES NOT EXIST:" + username)
		return nil, status.Error(codes.Internal, existsErr.Error())
	}
	user, er := u.service.GetByUsername(context.TODO(), username)
	if er != nil {
		return nil, status.Error(codes.Internal, er.Error())
	}
	token, tokenErr := u.tokenService.GenerateApiToken(user)
	if tokenErr != nil {
		u.logError.Logger.WithFields(logrus.Fields{
			"user": userNameCtx,
		}).Errorf("ERR:GEN TOKEN")
		return nil, status.Error(codes.Internal, tokenErr.Error())
	}
	return &pb.ApiToken{ApiToken: token}, nil
}

func (u UserHandler) ShareJobOffer(ctx context.Context, request *pb.ShareJobOfferRequest) (*pb.EmptyRequest, error) {

	policy := bluemonday.UGCPolicy()
	request.ShareJobOffer.ApiToken = strings.TrimSpace(policy.Sanitize(request.ShareJobOffer.ApiToken))
	request.ShareJobOffer.JobOffer.Publisher = strings.TrimSpace(policy.Sanitize(request.ShareJobOffer.JobOffer.Publisher))
	request.ShareJobOffer.JobOffer.Position = strings.TrimSpace(policy.Sanitize(request.ShareJobOffer.JobOffer.Position))
	request.ShareJobOffer.JobOffer.JobDescription = strings.TrimSpace(policy.Sanitize(request.ShareJobOffer.JobOffer.JobDescription))
	for i, _ := range request.ShareJobOffer.JobOffer.Requirements {
		request.ShareJobOffer.JobOffer.Requirements[i] = strings.TrimSpace(policy.Sanitize(request.ShareJobOffer.JobOffer.Requirements[i]))
	}
	request.ShareJobOffer.JobOffer.DatePosted = strings.TrimSpace(policy.Sanitize(request.ShareJobOffer.JobOffer.DatePosted))
	request.ShareJobOffer.JobOffer.Duration = strings.TrimSpace(policy.Sanitize(request.ShareJobOffer.JobOffer.Duration))
	p1 := common.BadJWTToken(request.ShareJobOffer.ApiToken)
	p2 := common.BadText(request.ShareJobOffer.JobOffer.Publisher)
	p3 := common.BadText(request.ShareJobOffer.JobOffer.Position)
	p4 := common.BadText(request.ShareJobOffer.JobOffer.JobDescription)
	p5 := common.BadTexts(request.ShareJobOffer.JobOffer.Requirements)
	p6 := common.BadDate(request.ShareJobOffer.JobOffer.DatePosted)
	p7 := common.BadDate(request.ShareJobOffer.JobOffer.Duration)
	//sqlInj := common.CheckForSQLInjection([]string{request.ShareJobOffer.ApiToken, request.ShareJobOffer.JobOffer.Publisher,
	//	request.ShareJobOffer.JobOffer.Position, request.ShareJobOffer.JobOffer.JobDescription, request.ShareJobOffer.JobOffer.DatePosted,
	//	request.ShareJobOffer.JobOffer.Duration})
	//sqlInj2 := common.CheckForSQLInjection(request.ShareJobOffer.JobOffer.Requirements)

	if request.ShareJobOffer.ApiToken == "" || request.ShareJobOffer.JobOffer.Publisher == "" ||
		request.ShareJobOffer.JobOffer.Position == "" || request.ShareJobOffer.JobOffer.JobDescription == "" ||
		request.ShareJobOffer.JobOffer.DatePosted == "" || request.ShareJobOffer.JobOffer.Duration == "" {
		u.logError.Logger.Errorf("ERR:XSS")
	} else if p1 || p2 || p3 || p4 || p5 || p6 || p7 {
		u.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
	} else {
		u.logInfo.Logger.Infof("INFO:Handling ShareJobOffer")
	}

	token := request.ShareJobOffer.ApiToken
	hasAccess, er := u.tokenService.CheckIfHasAccess(token)
	if er != nil {
		return &pb.EmptyRequest{}, er
	}
	if !hasAccess {
		u.logError.Logger.Errorf("ERR:DOES NOT HAVE ACCCESS")
		return &pb.EmptyRequest{}, errors.New("you don't have access")
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
		u.logError.Logger.Errorf("ERR:POST")
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)

	u.logInfo.Logger.Infof("INFO:REQ SENT TO POST SER")
	return &pb.EmptyRequest{}, nil
}

type myJSON struct {
	Array []string
}

func (u UserHandler) ActivateUserAccount(ctx context.Context, request *pb.ActivationRequest) (*pb.ActivationResponse, error) {
	//TODO:mzd dodati provjeru da li se uspelo ok mapirati?
	requestDto := api.MapPbToUserActivateRequest(request)
	err := u.validator.Struct(requestDto)
	if err != nil {
		u.logError.Logger.Errorf("ERR:INVALID REQ FIELDS")
		return &pb.ActivationResponse{Activated: false, Username: requestDto.Username}, err
	}
	policy := bluemonday.UGCPolicy()
	//sanitize everything
	requestDto.Username = strings.TrimSpace(policy.Sanitize(requestDto.Username))
	requestDto.Code = strings.TrimSpace(policy.Sanitize(requestDto.Code))
	//sqlInj := common.CheckForSQLInjection([]string{requestDto.Username, requestDto.Code})
	p1 := common.BadUsername(requestDto.Username)
	p2 := common.BadNumber(requestDto.Code)
	if requestDto.Username == "" || requestDto.Code == "" {
		u.logError.Logger.Errorf("ERR:XSS")
		return &pb.ActivationResponse{Activated: false, Username: requestDto.Username}, errors.New("fields are empty or xss happened")
	} else if p1 || p2 {
		u.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
		return &pb.ActivationResponse{Activated: false, Username: requestDto.Username}, errors.New("there is chance of sql injection happening")
	} else {
		u.logInfo.Logger.Infof("INFO:Handling ActivateUserAccount")
	}
	existsErr := u.service.UserExists(requestDto.Username)
	if existsErr != nil {
		u.logError.Logger.Errorf("ERR:USER DOES NOT EXIST")
		return &pb.ActivationResponse{Activated: false, Username: requestDto.Username}, existsErr
	}

	var code int
	code, convertError := strconv.Atoi(requestDto.Code)
	if convertError != nil {
		u.logError.Logger.Errorf("ERR:CONVERT")
		return &pb.ActivationResponse{Activated: false, Username: requestDto.Username}, convertError
	}
	activated, e := u.service.ActivateUserAccount(requestDto.Username, code)
	if e != nil {
		//u.logError.Logger.Errorf("ERR:")
		return &pb.ActivationResponse{Activated: false, Username: requestDto.Username}, e
	}
	if !activated {
		//u.logError.Logger.Errorf("account activation failed")
		return &pb.ActivationResponse{Activated: false, Username: requestDto.Username}, errors.New("account activation failed")
	}
	//u.logInfo.Logger.Infof("SUCCESS GenerateAPIToken")
	return &pb.ActivationResponse{Activated: activated, Username: requestDto.Username}, nil
}

func (u UserHandler) GetAll(ctx context.Context, request *pb.EmptyRequest) (*pb.GetAllResponse, error) {
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	u.logInfo.Logger.WithFields(logrus.Fields{
		"user": userNameCtx,
	}).Infof("INFO:Handling GetAll Users")
	users, err := u.service.GetUsers()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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

// UpdateUser TODO: SPOJITI S GRANOM EDITA
func (u UserHandler) UpdateUser(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateUserResponse, error) {
	userNameCtx := fmt.Sprintf(ctx.Value(interceptor.LoggedInUserKey{}).(string))
	u.logInfo.Logger.WithFields(logrus.Fields{
		"user": userNameCtx,
	}).Infof("INFO:Handling UpdateUser")

	return &pb.UpdateUserResponse{UpdatedUser: nil}, nil
}

func (u UserHandler) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	newUser := api.MapPbUserToNewUserDto(request)
	if err := u.validator.Struct(newUser); err != nil {
		u.logError.Logger.Errorf("ERR:INVALID REQ FILEDS")
		return nil, status.Error(codes.FailedPrecondition, err.Error())
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

	p1 := common.BadUsername(newUser.Username)
	p2 := common.BadName(newUser.FirstName)
	p3 := common.BadName(newUser.LastName)
	p4 := common.BadEmail(newUser.Email)
	p5 := common.BadText(newUser.Gender)
	p6 := common.BadDate(newUser.DateOfBirth)
	p7 := common.BadNumber(newUser.PhoneNumber)
	p8 := common.BadEmail(newUser.RecoveryEmail)
	p9 := common.BadPassword(newUser.Password)

	//sqlInj := common.CheckForSQLInjection([]string{newUser.Username, newUser.FirstName, newUser.LastName, newUser.Email,
	//	newUser.Gender, newUser.DateOfBirth, newUser.PhoneNumber, newUser.RecoveryEmail})

	if newUser.Username == "" || newUser.FirstName == "" || newUser.LastName == "" ||
		newUser.Gender == "" || newUser.DateOfBirth == "" || newUser.PhoneNumber == "" ||
		newUser.Password == "" || newUser.Email == "" || newUser.RecoveryEmail == "" {
		u.logError.Logger.Errorf("ERR:XSS")
		return nil, status.Error(codes.FailedPrecondition, "fields are empty or xss happened")
		//return nil, errors.New("fields are empty or xss happened")
	} else if p1 || p2 || p3 || p4 || p5 || p6 || p7 || p8 || p9 {
		u.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
		//return nil, errors.New("there is chance for sql injection")
		return nil, status.Error(codes.FailedPrecondition, "there is chance for sql injection")
	} else {
		u.logInfo.Logger.Infof("INFO:Handling RegisterUser")
	}
	err := u.service.UserExists(newUser.Username)
	if err == nil {
		u.logError.Logger.Errorf("USER ALREADY EXISTS")
		return nil, status.Error(codes.AlreadyExists, "user already exists")
	}

	var hashedSaltedPassword = ""
	validPassword := u.passwordUtil.IsValidPassword(newUser.Password)

	if validPassword {
		pass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			u.logError.Logger.Errorf("ERR:BCRYPT")
			return nil, status.Error(codes.Internal, err.Error())
		}
		hashedSaltedPassword = string(pass)
	} else {
		u.logError.Logger.Errorf("ERR:PASSWORD FORMAT NOT VALID")
		return nil, status.Error(codes.FailedPrecondition, "password format is not valid")
	}
	newUser.Password = hashedSaltedPassword
	registeredUser, er := u.service.CreateRegisteredUser(api.MapDtoToUser(newUser))

	if er != nil {
		return nil, status.Error(codes.Internal, er.Error())
	}

	//u.logInfo.Logger.Infof("SUCCESS RegisterUser")
	return &pb.RegisterUserResponse{RegisteredUser: api.MapUserToPbResponseUser(registeredUser)}, nil
}

func (u UserHandler) SendRequestForPasswordRecovery(ctx context.Context, request *pb.PasswordRecoveryRequest) (*pb.PasswordRecoveryResponse, error) {
	var requestUsername = request.Username.Username
	policy := bluemonday.UGCPolicy()
	requestUsername = strings.TrimSpace(policy.Sanitize(requestUsername))
	//sqlInj := common.CheckRegexSQL(requestUsername)
	sqlInj := common.BadUsername(requestUsername)
	if requestUsername == "" {
		u.logError.Logger.Errorf("XSS")
		//http.Error(rw, "Field empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return &pb.PasswordRecoveryResponse{CodeSent: false}, errors.New("fields are empty or xss happened")
	} else if sqlInj {
		u.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
		return nil, errors.New("there is chance for sql injection")
	} else {
		u.logInfo.Logger.Println("INFO:Handling PASSWORD RECCOVERY ")
	}
	existsErr := u.service.UserExists(requestUsername)
	if existsErr != nil {
		u.logError.Logger.Errorf("USER DOES NOT EXIST")
		return &pb.PasswordRecoveryResponse{CodeSent: false}, existsErr
	}

	codeSent, codeErr := u.service.SendCodeToRecoveryMail(requestUsername)
	if codeErr != nil {
		//u.logError.Logger.Errorf(codeErr.Error())
		return &pb.PasswordRecoveryResponse{CodeSent: false}, codeErr
	}
	if !codeSent {
		u.logError.Logger.Errorf("ERR:ACCOUNT ACTIVATION FAILED")
		return &pb.PasswordRecoveryResponse{CodeSent: false}, errors.New("account activation failed")
	}

	return &pb.PasswordRecoveryResponse{CodeSent: true}, nil
}

func (u UserHandler) RecoverPassword(ctx context.Context, request *pb.NewPasswordRequest) (*pb.NewPasswordResponse, error) {
	requestDto := api.MapPbToNewPasswordRequestDto(request)

	err := u.validator.Struct(requestDto)
	if err != nil {
		u.logError.Logger.Error("ERR:INVALID REQ FIELDS")
		return &pb.NewPasswordResponse{PasswordChanged: false}, err
	}
	policy := bluemonday.UGCPolicy()
	//sanitize everything
	requestDto.Username = strings.TrimSpace(policy.Sanitize(requestDto.Username))
	requestDto.Code = strings.TrimSpace(policy.Sanitize(requestDto.Code))
	requestDto.NewPassword = strings.TrimSpace(policy.Sanitize(requestDto.NewPassword))
	//sqlInj := common.CheckForSQLInjection([]string{requestDto.Username, requestDto.Code})
	p1 := common.BadUsername(requestDto.Username)
	p2 := common.BadNumber(requestDto.Code)
	p3 := common.BadPassword(requestDto.NewPassword)
	if requestDto.Username == "" || requestDto.Code == "" || requestDto.NewPassword == "" {
		u.logError.Logger.Error("ERR:XSS")
		return &pb.NewPasswordResponse{PasswordChanged: false}, errors.New("fields are empty or xss happened")
	} else if p1 || p2 || p3 {
		u.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
		return nil, errors.New("there is chance for sql injection")
	} else {
		u.logInfo.Logger.Println("INFO:Handling RecoverPassword")
	}

	existsErr := u.service.UserExists(requestDto.Username)
	if existsErr != nil {
		u.logError.Logger.Errorf("USER DOES NOT EXIST")
		return &pb.NewPasswordResponse{PasswordChanged: false}, existsErr
	}
	///////////////////
	var hashedSaltedPassword = ""
	validPassword := u.passwordUtil.IsValidPassword(requestDto.NewPassword)

	if validPassword {
		pass, err := bcrypt.GenerateFromPassword([]byte(requestDto.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			u.logError.Logger.Errorf("ERR:BCRYPT")
			return &pb.NewPasswordResponse{PasswordChanged: false}, err
		}

		hashedSaltedPassword = string(pass)

	} else {
		fmt.Println("Password format is not valid!")
		u.logError.Logger.Errorf("ERR:PASSWORD FORMAT NOT VALID")
		return &pb.NewPasswordResponse{PasswordChanged: false}, errors.New("password format is not valid")
	}
	passChanged, err := u.service.CreateNewPassword(requestDto.Username, hashedSaltedPassword, requestDto.Code)
	if err != nil {
		//http.Error(rw, err.Error(), http.StatusConflict) //409
		return &pb.NewPasswordResponse{PasswordChanged: false}, err
	}
	if !passChanged {
		//http.Error(rw, "error changing password", http.StatusConflict) //409
		u.logError.Logger.Errorf("ERR:CHANGING PASSWORD")
		return &pb.NewPasswordResponse{PasswordChanged: false}, errors.New("error changing password")
	}
	return &pb.NewPasswordResponse{PasswordChanged: true}, nil
}

func (u UserHandler) PwnedPassword(ctx context.Context, request *pb.PwnedRequest) (*pb.PwnedResponse, error) {
	pwnedPassword := request.Password.Password
	policy := bluemonday.UGCPolicy()
	//sanitize everything
	pwnedPassword = strings.TrimSpace(policy.Sanitize(pwnedPassword))
	sqlInj := common.BadPassword(pwnedPassword)
	if pwnedPassword == "" {
		u.logError.Logger.Errorf("XSS")
		//http.Error(rw, "Fields are empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return &pb.PwnedResponse{Pwned: true, Message: "fields are empty or xss happened"}, errors.New("fields are empty or xss happened")
	} else if sqlInj {
		u.logError.Logger.Errorf("ERR:BAD VALIDATION: POSIBLE INJECTION")
		return nil, errors.New("there is chance for sql injection")
	} else {
		u.logInfo.Logger.Println("INFO:Handling PWNED PASSWORD")
	}

	pwned, err := u.pwnedClient.Compromised(pwnedPassword)
	if err != nil {
		u.logError.Logger.Errorf("ERR:PWNED CLIENT")
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
	return &pb.PwnedResponse{Pwned: pwned, Message: mess}, nil

}
