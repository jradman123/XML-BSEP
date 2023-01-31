package services

import (
	"common/module/logger"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/trycourier/courier-go/v2"
	"math/rand"
	tracer "monitoring/module"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
	"user/module/domain/dto"
	"user/module/domain/model"
	"user/module/domain/repositories"
	"user/module/infrastructure/api"
	"user/module/infrastructure/orchestrators"
)

type UserService struct {
	logInfo        *logger.Logger
	logError       *logger.Logger
	userRepository repositories.UserRepository
	emailRepo      repositories.EmailVerificationRepository
	recoveryRepo   repositories.PasswordRecoveryRequestRepository
	orchestrator   *orchestrators.UserOrchestrator
}

var (
	EmailFormatInvalid     = errors.New("EMAIL FORMAT INVALID")
	EmailDomainInvalid     = errors.New("EMAIL DOMAIN INVALID")
	ErrorEmailVerification = errors.New("ERROR EMAIL VERIFICATION")
	ErrorOrchestrator      = errors.New("ORCHESTRATOR")
	DbError                = errors.New("DB ERROR")
	ErrorCreatingUser      = errors.New("ERROR CREATING USER:check your email, you cant use the same email for 2 accounts")
	subject                = "Activation code"
	body                   = "Welcome to Dislinkt! Here is your activation code:"
)

func NewUserService(logInfo *logger.Logger, logError *logger.Logger, repository repositories.UserRepository, emailRepo repositories.EmailVerificationRepository,
	recoveryRepo repositories.PasswordRecoveryRequestRepository, orchestrator *orchestrators.UserOrchestrator) *UserService {
	return &UserService{logInfo, logError, repository, emailRepo, recoveryRepo, orchestrator}
}

func (u UserService) GetUsers() ([]model.User, error) {

	users, err := u.userRepository.GetUsers()
	if err != nil {
		fmt.Sprintln("evo ovde sam puko - service")
		u.logError.Logger.Errorf("ERR:CANT GET USERS")
		return nil, errors.New("cant get users")
	}

	return users, err

}

func (u UserService) GetByUsername(username string) (*model.User, error) {
	user, err := u.userRepository.GetByUsername(username)

	if err != nil {
		u.logError.Logger.Errorf("ERR:INVALID USERNAME:" + username)
		return nil, err
	}

	return user, nil
}

func (u UserService) GetUserSalt(username string) (string, error) {

	salt, err := u.userRepository.GetUserSalt(username)

	if err != nil {
		return "", err
	}
	return salt, nil
}

func (u UserService) UserExists(username string) error {
	err := u.userRepository.UserExists(username)

	if err != nil {
		//u.logError.Logger.Errorf("ERR:USER DOES NOT EXIST:" + username)
		return err
	}
	return nil
}

func (u UserService) GetUserRole(username string, ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "GetUserRoleService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	role, err := u.userRepository.GetUserRole(username, ctx)

	if err != nil {
		return "", err
	}
	return role, nil
}

func (u UserService) CreateRegisteredUser(user *model.User, ctx context.Context) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "WriteNewUserInDB")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	var er = checkEmailValid(user.Email, ctx)
	if er != nil {
		u.logError.Logger.Println(EmailFormatInvalid)
		return nil, EmailFormatInvalid
	}
	var domEr = checkEmailDomain(user.Email, ctx)
	if domEr != nil {
		u.logError.Logger.Println(EmailDomainInvalid)
		return nil, EmailDomainInvalid
	}

	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(100000)

	regUser, err := u.userRepository.CreateRegisteredUser(user, ctx)
	if err != nil {
		u.logError.Logger.Println(DbError)
		return regUser, ErrorCreatingUser
	}
	emailVerification := model.EmailVerification{
		ID:       uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		VerCode:  rn,
		Time:     time.Now(),
	}

	_, e := u.emailRepo.CreateEmailVerification(&emailVerification, ctx)
	fmt.Println(e)
	if e != nil {
		u.logError.Logger.Println(ErrorEmailVerification)
		return nil, ErrorEmailVerification
	}
	sendMailWithCourier(user.Email, strconv.Itoa(rn), subject, body, u.logError, ctx)

	err = u.orchestrator.CreateUser(user)
	if err != nil {
		u.logError.Logger.Println(ErrorOrchestrator)
		return regUser, err
	}

	err = u.orchestrator.CreateConnectionUser(user)
	if err != nil {
		u.logError.Logger.Println(ErrorOrchestrator)
		return regUser, err
	}

	return regUser, nil
}

func (u UserService) ActivateUserAccount(username string, verCode int, ctx context.Context) (bool, error) {

	var allVerForUsername []model.EmailVerification
	var dbEr error
	span1 := tracer.StartSpanFromContext(ctx, "ReadVerificationForUser")
	allVerForUsername, dbEr = u.emailRepo.GetVerificationByUsername(username)
	span1.Finish()
	if dbEr != nil {
		u.logError.Logger.Errorf("ERR:DB:CODE DOES NOT EXIST FOR USER")
		return false, dbEr
	}
	codeInfoForUsername, err := findMostRecent(allVerForUsername)
	if err != nil {
		u.logError.Logger.Errorf("ERR:NO VERIFICATION FOR USER")
		return false, err
	}

	if codeInfoForUsername.VerCode == verCode {

		if codeInfoForUsername.Time.Add(time.Hour).After(time.Now()) {
			user, err := u.userRepository.GetByUsername(username)
			if err != nil {
				fmt.Println(err)
				u.logError.Logger.Errorf("ERR:DB")
				return false, err
			}
			user.IsConfirmed = true
			activated, actErr := u.userRepository.ActivateUserAccount(user)
			err = u.orchestrator.ActivateUserAccount(user)
			if err != nil {
				return false, err
			}

			if actErr != nil {
				u.logError.Logger.Errorf("ERR:WHILE ACTIVATING USER")
				return false, actErr
			}
			if !activated {
				u.logError.Logger.Errorf("ERR:ACTIVATION FAILED")
				return false, errors.New("user not activated")
			}
			fmt.Println("uspeoo sam da ga aktiviram")
			return true, nil

		} else {
			u.logError.Logger.Errorf("ERR:CODE EXPIRED")
			return false, errors.New("code expired")
		}
	} else {
		u.logError.Logger.Errorf("ERR:WRONG CODE")
		return false, errors.New("wrong code")
	}
}

func findMostRecent(verifications []model.EmailVerification) (*model.EmailVerification, error) {

	if len(verifications) > 1 {
		latest := verifications[0]
		latestIdx := 0
		fmt.Println(latest)
		fmt.Println(latestIdx)
		for i, ver := range verifications {
			if ver.Time.After(latest.Time) {
				latest = ver
				latestIdx = i
			}
		}
		return &latest, nil
	} else {
		if len(verifications) > 0 {
			return &verifications[0], nil
		} else {
			return nil, errors.New("verifications array empty ")
		}
	}

}

func (u UserService) SendCodeToRecoveryMail(username string, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "SendCodeToRecoveryMail")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err := u.userRepository.GetByUsername(username)

	if err != nil {
		u.logError.Logger.Println("ERR:USER DOES NOT EXIST")
		return false, err
	}

	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(100000)
	recovery := model.PasswordRecoveryRequest{
		ID:            uuid.New(),
		Username:      username,
		Email:         user.Email,
		RecoveryEmail: user.RecoveryEmail,
		IsUsed:        false,
		Time:          time.Now(),
		RecoveryCode:  rn,
	}

	fmt.Println(recovery)

	//obrisi prethodni zahtev ako postoji jer eto da uvijek samo poslednji ima u bazi
	span1 := tracer.StartSpanFromContext(ctx, "DeleteAllRequestsForUser")
	deleteErr := u.recoveryRepo.ClearOutRequestsForUsername(username)
	span1.Finish()

	if deleteErr != nil {
		return false, deleteErr
	}

	span2 := tracer.StartSpanFromContext(ctx, "WriteNewRequestForPasswordRecovery")
	_, e := u.recoveryRepo.CreatePasswordRecoveryRequest(&recovery)
	span2.Finish()

	fmt.Println(e)
	if e != nil {
		u.logError.Logger.Println("ERR:PASS RECOVERY REQ")
		return false, e
	} else {
		u.logInfo.Logger.Infof("INFO:CREATED PASS RECOVERY")
	}

	//mzd staviti da ovo vraca bool i da ima parametar poruku i zaglavlje
	sendMailWithCourier(user.RecoveryEmail, strconv.Itoa(rn), "Password recovery code", "Here is your code:", u.logInfo, context.TODO())
	return true, nil
}

func (u UserService) CreateNewPassword(username string, newHashedPassword string, code string, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateNewPassword")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	var passwordRecoveryRequest *model.PasswordRecoveryRequest
	var dbEr error

	span1 := tracer.StartSpanFromContext(ctx, "ReadPasswordRecoveryRequestByUsername")
	passwordRecoveryRequest, dbEr = u.recoveryRepo.GetPasswordRecoveryRequestByUsername(username)
	span1.Finish()

	if dbEr != nil {
		u.logError.Logger.Errorf("ERR:THERE IS NOT A PASS RECOVERY REQUEST IN DATABASE FOR USER:" + username)
		fmt.Println(dbEr)

		return false, dbEr
	}
	fmt.Println("verCode:", passwordRecoveryRequest.RecoveryCode)

	var codeInt, convErr = strconv.Atoi(code)
	if convErr != nil {
		u.logError.Logger.Errorf("ERR:CONVERTING CODE TO INT")
		return false, errors.New("error converting code to int")
	}
	if passwordRecoveryRequest.RecoveryCode == codeInt {
		//kao dala sam kodu trajanje od 1h
		fmt.Println("kod se poklapa")
		if passwordRecoveryRequest.Time.Add(time.Minute * 3).After(time.Now()) {
			if !passwordRecoveryRequest.IsUsed {
				fmt.Println("vreme se uklapa")
				//ako je kod ok i ako je u okviru vremena trajanja mjenjamo mu status
				user, err := u.userRepository.GetByUsername(username)
				if err != nil {
					fmt.Println(err)
					u.logError.Logger.Errorf("ERR:NO USER")
					fmt.Println("error u get by username kod ucitavanja usera")
					return false, err
				}

				fmt.Println(user.Username)

				//sacuvati izmjene korisnika,tj izmjenjen password
				span2 := tracer.StartSpanFromContext(ctx, "WriteNewPasswordInDB")
				changePassErr := u.userRepository.ChangePassword(user, newHashedPassword, ctx)
				span2.Finish()

				if changePassErr != nil {
					fmt.Println("error pri cuvanju novog pass")
					u.logError.Logger.Errorf("ERR:SAVING NEW PASSWORD")
					return false, changePassErr
				}
				//staviti iskoristen kod na true

				//service.Repo.ActivateUserAccount(user)
				_, er := u.userRepository.GetByUsername(username)
				if er != nil {
					fmt.Println(er)

					u.logError.Logger.Errorf("ERR:NO USER")
					fmt.Println("FAK MAJ LAJF 2")
					return false, er
				}
				return true, nil
			} else {
				fmt.Println("kod iskoristen")
				u.logError.Logger.Errorf("ERR:CODE USED")
				return false, errors.New("code used")
			}

		} else {
			fmt.Println("istekao kod")
			u.logError.Logger.Errorf("ERR:CODE EXPIRED")
			return false, errors.New("code expired")
		}

	} else {
		fmt.Println("ne valjda kod")
		u.logError.Logger.Errorf("ERR:WRONG CODE")
		return false, errors.New("wrong code")
	}
}

func checkEmailValid(email string, ctx context.Context) error {
	// check email syntax is valid
	//func MustCompile(str string) *Regexp
	span := tracer.StartSpanFromContext(ctx, "checkEmailValid")
	defer span.Finish()

	emailRegex, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		fmt.Println(err)
		return errors.New("sorry, something went wrong")
	}
	rg := emailRegex.MatchString(email)
	if !rg {
		return errors.New("email address is not a valid syntax, please check again")
	}
	// check email length
	if len(email) < 4 {
		return errors.New("email length is too short")
	}
	if len(email) > 253 {
		return errors.New("email length is too long")
	}
	return nil
}

func checkEmailDomain(email string, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "checkEmailDomain")
	defer span.Finish()

	i := strings.Index(email, "@")
	host := email[i+1:]
	// func LookupMX(name string) ([]*MX, error)
	_, err := net.LookupMX(host)
	if err != nil {
		err = errors.New("eould not find email's domain server, please chack and try again")
		return err
	}
	return nil
}

func sendMailWithCourier(email string, code string, subject string, body string, logErr *logger.Logger, ctx context.Context) {
	span := tracer.StartSpanFromContext(ctx, "sendMailWithCourier")
	defer span.Finish()

	client := courier.CreateClient("pk_prod_0FQXVBPMDHMZ3VJ3WN6CYC12KNMH", nil)
	fmt.Println(code)
	requestID, err := client.SendMessage(
		context.Background(),
		courier.SendMessageRequestBody{
			Message: map[string]interface{}{
				"to": map[string]string{
					"email": email,
				},
				"content": map[string]string{
					"title": subject,
					"body":  body + code,
				},
				"data": map[string]string{
					"joke": "What did C++ say to C? You have no class.",
					"code": code,
				},
			},
		})

	if err != nil {
		logErr.Logger.Println("ERR:SENDING MAIL")
		fmt.Println(err)
	}
	fmt.Println(requestID)
}

func (u UserService) EditUser(userDetails *dto.UserDetails, ctx context.Context) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "EditUserService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err := u.GetByUsername(userDetails.Username)
	if err != nil {
		return nil, err
	}
	user = api.MapUserDetailsDtoToUser(userDetails, user)

	span1 := tracer.StartSpanFromContext(ctx, "WriteChangedUserDetailsInDB")
	edited, e := u.userRepository.EditUserDetails(user)
	span1.Finish()

	if e != nil {
		return nil, e
	}
	if !edited {
		return nil, errors.New("user was not edited")
	}
	err = u.orchestrator.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	err = u.orchestrator.EditConnectionUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) ChangeProfileStatus(username string, newStatus string, ctx context.Context) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "ChangeProfileStatusService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err := u.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	user.ProfileStatus = model.ProfileStatus(newStatus)

	span1 := tracer.StartSpanFromContext(ctx, "WriteChangedProfileStatusInDB")
	edited, e := u.userRepository.EditUserDetails(user)
	span1.Finish()

	if e != nil {
		return nil, e
	}
	if !edited {
		return nil, errors.New("user status was not edited")
	}
	err = u.orchestrator.ChangeProfileStatus(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserService) EditUserPersonalDetails(userPersonalDetails *dto.UserPersonalDetails, ctx context.Context) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "EditUserPersonalDetailsService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err := u.GetByUsername(userPersonalDetails.Username)
	if err != nil {
		return nil, err
	}
	user = api.MapUserPersonalDetailsDtoToUser(userPersonalDetails, user)

	span1 := tracer.StartSpanFromContext(ctx, "WriteChangedUserDetailsInDB")
	edited, e := u.userRepository.EditUserDetails(user)
	span1.Finish()

	if e != nil {
		return nil, e
	}
	if !edited {
		return nil, errors.New("user was not edited")
	}
	err = u.orchestrator.EditConnectionUser(user)
	if err != nil {
		return nil, err
	}
	err = u.orchestrator.UpdateUser(user)
	return user, nil
}

func (u UserService) EditUserProfessionalDetails(userProfessionalDetails *dto.UserProfessionalDetails, ctx context.Context) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "EditUserProfessionalDetailsService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	user, err := u.GetByUsername(userProfessionalDetails.Username)
	if err != nil {
		return nil, err
	}
	user = api.MapUserProfessionalDetailsDtoToUser(userProfessionalDetails, user)
	fmt.Println(user)
	fmt.Println("edit professional details user_service user service ")
	err = u.orchestrator.EditConnectionUserProfessionalDetails(user)
	if err != nil {
		return nil, err
	}

	span1 := tracer.StartSpanFromContext(ctx, "WriteChangedProfessionalDetailsInDB")
	edited, e := u.userRepository.EditUserDetails(user)
	span1.Finish()

	if e != nil {
		return nil, e
	}
	if !edited {
		return nil, errors.New("user was not edited")
	}
	//err = u.orchestrator.EditConnectionUser(user)
	//if err != nil {
	//	return nil, err
	//}
	return user, nil
}

func (u UserService) CheckIfEmailExists(id uuid.UUID, email string, ctx context.Context) bool {
	span := tracer.StartSpanFromContext(ctx, "CheckIfEmailExists")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	users, _ := u.userRepository.GetUsers()
	for _, element := range users {
		if element.ID == id {
			continue
		}
		if element.Email == email {
			return true
		}
	}
	return false
}

func (u UserService) CheckIfUsernameExists(id uuid.UUID, username string, ctx context.Context) bool {
	span := tracer.StartSpanFromContext(ctx, "CheckIfUsernameExists")
	defer span.Finish()

	users, _ := u.userRepository.GetUsers()
	for _, element := range users {
		if element.ID == id {
			continue
		}
		if element.Username == username {
			return true
		}
	}
	return false
}

func (u UserService) GetById(id uuid.UUID) (*model.User, error) {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (u UserService) UpdateEmail(user *model.User) (*model.User, error) {

	result, err := u.userRepository.UpdateEmail(user)
	if err != nil {
		return nil, err
	}
	if result {
		user, _ := u.userRepository.GetById(user.ID)
		er := u.orchestrator.ChangeEmail(user)
		if er != nil {
			return nil, er
		}
		return user, nil
	} else {
		return nil, nil
	}
}

func (u UserService) UpdateUsername(ctx context.Context, user *model.User) (*model.User, error) {
	span := tracer.StartSpanFromContext(ctx, "UpdateUsernameService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	result, err := u.userRepository.UpdateUsername(ctx, user)
	if err != nil {
		return nil, err
	}
	if result {
		user, _ := u.userRepository.GetById(user.ID)
		return user, nil
	} else {
		return nil, nil
	}
}
