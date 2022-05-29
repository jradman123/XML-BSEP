package handlers

import (
	pb "common/module/proto/user_service"
	"context"
	"errors"
	"fmt"
	hibp "github.com/mattevans/pwned-passwords"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"strconv"
	"strings"
	"user/module/application/helpers"
	"user/module/application/services"
	"user/module/infrastructure/api"
)

type UserHandler struct {
	l            *log.Logger
	service      *services.UserService
	jsonConv     *helpers.JsonConverters
	validator    *validator.Validate
	passwordUtil *helpers.PasswordUtil
	pwnedClient  *hibp.Client
}

func NewUserHandler(l *log.Logger, service *services.UserService, jsonConv *helpers.JsonConverters, validator *validator.Validate,
	passwordUtil *helpers.PasswordUtil, pwnedClient *hibp.Client) *UserHandler {
	return &UserHandler{l, service, jsonConv, validator, passwordUtil, pwnedClient}
}

func (u UserHandler) MustEmbedUnimplementedUserServiceServer() {
	u.l.Println("Handling MustEmbedUnimplementedUserServiceServer Users")
}

func (u UserHandler) ActivateUserAccount(ctx context.Context, request *pb.ActivationRequest) (*pb.ActivationResponse, error) {
	u.l.Println("Handling ActivateUserAccount ")
	//TODO:mzd dodati provjeru da li se uspelo ok mapirati?
	requstDto := api.MapPbToUserActivateRequest(request)
	//TODO:dodati validaciju u obliku regexa, spreciti injection napad
	err := u.validator.Struct(requstDto)
	if err != nil {
		u.l.Println(err)
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, err
		//http.Error(rw, "New user dto fields aren't entered in valid format! error:"+err.Error(), http.StatusExpectationFailed) //400
	}
	policy := bluemonday.UGCPolicy()
	//sanitize everything
	requstDto.Username = strings.TrimSpace(policy.Sanitize(requstDto.Username))
	requstDto.Code = strings.TrimSpace(policy.Sanitize(requstDto.Code))
	if requstDto.Username == "" || requstDto.Code == "" {
		u.l.Println("fields are empty or xss")
		//http.Error(rw, "Fields are empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, errors.New("fields are empty or xss happened")
	}
	existsErr := u.service.UserExists(requstDto.Username)
	if existsErr != nil {
		u.l.Println(existsErr)
		//http.Error(rw, "User with entered username already exists!", http.StatusConflict) //409
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, existsErr
	}
	var code int
	code, convertError := strconv.Atoi(requstDto.Code)
	if convertError != nil {
		u.l.Println(convertError)
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, convertError
		//http.Error(rw, "Error converting code from string to int! error:"+convertError.Error(), http.StatusConflict) //409
	}
	activated, e := u.service.ActivateUserAccount(requstDto.Username, code)
	if e != nil {
		u.l.Println(e)
		//http.Error(rw, e.Error(), http.StatusConflict) //409
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, e
	}
	if !activated {
		u.l.Println("account activation failed")
		//http.Error(rw, "Account activation failed!", http.StatusConflict) //409
		return &pb.ActivationResponse{Activated: false, Username: requstDto.Username}, errors.New("account activation failed")
	}
	u.l.Println("skoro pa kraj")
	return &pb.ActivationResponse{Activated: activated, Username: requstDto.Username}, nil
}

func (u UserHandler) GetAll(ctx context.Context, request *pb.EmptyRequest) (*pb.GetAllResponse, error) {
	u.l.Println("Handling GetAll Users")
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
	u.l.Println("Handling UpdateUser Users")

	return &pb.UpdateUserResponse{UpdatedUser: nil}, nil
}

func (u UserHandler) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	//	fmt.Println(request.UserRequest.Email)
	u.l.Println("Handling RegisterUser")
	//TODO:mzd dodati provjeru da li se uspelo ok mapirati?
	newUser := api.MapPbUserToNewUserDto(request)
	//TODO:dodati validaciju u obliku regexa, spreciti injection napad
	if err := u.validator.Struct(newUser); err != nil {
		fmt.Println(err)
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
		fmt.Println("fields are empty or xss")
		//http.Error(rw, "Fields are empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return nil, errors.New("fields are empty or xss happened")
	}

	err := u.service.UserExists(newUser.Username)
	if err == nil {
		fmt.Println(err)
		//http.Error(rw, "User with entered username already exists!", http.StatusConflict) //409
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
		fmt.Println("Password format is not valid!")
		//http.Error(rw, "Password format is not valid! error:"+err.Error(), http.StatusBadRequest) //400
		return nil, errors.New("password format is not valid")
	}
	newUser.Password = hashedSaltedPassword
	registeredUser, er := u.service.CreateRegisteredUser(api.MapDtoToUser(newUser))

	if er != nil {
		fmt.Println(er)
		//http.Error(rw, "Failed creating registered user! error:"+er.Error(), http.StatusExpectationFailed) //
		return nil, er
	}

	return &pb.RegisterUserResponse{RegisteredUser: api.MapUserToPbResponseUser(registeredUser)}, nil
}

func (u UserHandler) SendRequestForPasswordRecovery(ctx context.Context, request *pb.PasswordRecoveryRequest) (*pb.PasswordRecoveryResponse, error) {

	u.l.Println("Handling PASSWORD RECCOVERY ")
	//TODO: injection
	var requestUsername = request.Username.Username
	policy := bluemonday.UGCPolicy()
	requestUsername = strings.TrimSpace(policy.Sanitize(requestUsername))

	if requestUsername == "" {
		fmt.Println("usrnname empty or xss")
		//http.Error(rw, "Field empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return &pb.PasswordRecoveryResponse{CodeSent: false}, errors.New("fields are empty or xss happened")
	}
	existsErr := u.service.UserExists(requestUsername)
	if existsErr != nil {
		u.l.Println(existsErr)
		//http.Error(rw, "User with entered username already exists!", http.StatusConflict) //409
		return &pb.PasswordRecoveryResponse{CodeSent: false}, existsErr
	}

	codeSent, codeErr := u.service.SendCodeToRecoveryMail(requestUsername)
	if codeErr != nil {
		u.l.Println(codeErr)
		//http.Error(rw, e.Error(), http.StatusConflict) //409
		return &pb.PasswordRecoveryResponse{CodeSent: false}, codeErr
	}
	if !codeSent {
		u.l.Println("account activation failed")
		//http.Error(rw, "Account activation failed!", http.StatusConflict) //409
		return &pb.PasswordRecoveryResponse{CodeSent: false}, errors.New("account activation failed")
	}

	return &pb.PasswordRecoveryResponse{CodeSent: true}, nil
}

func (u UserHandler) RecoverPassword(ctx context.Context, request *pb.NewPasswordRequest) (*pb.NewPasswordResponse, error) {

	u.l.Println("Handling RecoverPassword handler ")
	//TODO:mzd dodati provjeru da li se uspelo ok mapirati?
	requestDto := api.MapPbToNewPasswordRequestDto(request)

	err := u.validator.Struct(requestDto)
	if err != nil {
		u.l.Println(err)
		return &pb.NewPasswordResponse{PasswordChanged: false}, err
		//http.Error(rw, "New user dto fields aren't entered in valid format! error:"+err.Error(), http.StatusExpectationFailed) //400
	}
	policy := bluemonday.UGCPolicy()
	//sanitize everything
	requestDto.Username = strings.TrimSpace(policy.Sanitize(requestDto.Username))
	requestDto.Code = strings.TrimSpace(policy.Sanitize(requestDto.Code))
	requestDto.NewPassword = strings.TrimSpace(policy.Sanitize(requestDto.NewPassword))
	if requestDto.Username == "" || requestDto.Code == "" || requestDto.NewPassword == "" {
		u.l.Println("fields are empty or xss")
		//http.Error(rw, "Fields are empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return &pb.NewPasswordResponse{PasswordChanged: false}, errors.New("fields are empty or xss happened")
	}

	existsErr := u.service.UserExists(requestDto.Username)
	if existsErr != nil {
		u.l.Println(existsErr)
		//http.Error(rw, "User with entered username already exists!", http.StatusConflict) //409
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

	u.l.Println("Handling PWNED PASSWORD")
	u.l.Println(request.Password.Password)
	pwnedPassword := request.Password.Password
	policy := bluemonday.UGCPolicy()
	//sanitize everything
	pwnedPassword = strings.TrimSpace(policy.Sanitize(pwnedPassword))
	if pwnedPassword == "" {
		u.l.Println("fields are empty or xss")
		//http.Error(rw, "Fields are empty or xss attack happened! error:"+err.Error(), http.StatusExpectationFailed) //400
		return &pb.PwnedResponse{Pwned: true, Message: "fields are empty or xss happened"}, errors.New("fields are empty or xss happened")
	}

	pwned, err := u.pwnedClient.Compromised(pwnedPassword)
	if err != nil {
		fmt.Println(err)
		u.l.Println(pwnedPassword)
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
