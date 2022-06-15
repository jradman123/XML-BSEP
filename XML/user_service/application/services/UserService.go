package services

import (
	"common/module/logger"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/trycourier/courier-go/v2"
	"math/rand"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
	"user/module/domain/model"
	"user/module/domain/repositories"
)

type UserService struct {
	logInfo        *logger.Logger
	logError       *logger.Logger
	userRepository repositories.UserRepository
	emailRepo      repositories.EmailVerificationRepository
	recoveryRepo   repositories.PasswordRecoveryRequestRepository
}

func NewUserService(logInfo *logger.Logger, logError *logger.Logger, repository repositories.UserRepository, emailRepo repositories.EmailVerificationRepository, recoveryRepo repositories.PasswordRecoveryRequestRepository) *UserService {
	return &UserService{logInfo, logError, repository, emailRepo, recoveryRepo}
}

func (u UserService) GetUsers() ([]model.User, error) {

	users, err := u.userRepository.GetUsers()
	if err != nil {
		u.logError.Logger.Errorf("ERR:CANT GET USERS")
		return nil, errors.New("cant get users")
	}

	return users, err

}

func (u UserService) GetByUsername(ctx context.Context, username string) (*model.User, error) {

	user, err := u.userRepository.GetByUsername(ctx, username)

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

func (u UserService) GetUserRole(username string) (string, error) {

	role, err := u.userRepository.GetUserRole(username)

	if err != nil {
		return "", err
	}
	return role, nil
}
func (u UserService) CreateRegisteredUser(user *model.User) (*model.User, error) {

	var er = checkEmailValid(user.Email)
	if er != nil {
		u.logError.Logger.Println("ERR:EMAIL FORMAT INVALID")
		return nil, errors.New("email format invalid")
	}
	var domEr = checkEmailDomain(user.Email)
	if domEr != nil {
		u.logError.Logger.Println("ERR:EMAIL DOMAIN INVALID")
		return nil, errors.New("email domain invalid")
	}
	//TODO:ZAMENITI TOKENOM
	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(100000)
	emailVerification := model.EmailVerification{
		ID:       uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		VerCode:  rn,
		Time:     time.Now(),
	}
	fmt.Println(emailVerification)

	_, e := u.emailRepo.CreateEmailVerification(&emailVerification)
	fmt.Println(e)
	if e != nil {
		u.logError.Logger.Println("ERR:SAVING EMAIL VERIFICATION")
		return nil, errors.New("error saving emailVerification")
	}

	sendMailWithCourier(user.Email, strconv.Itoa(rn), "Activation code", "Welcome to Dislinkt! Here is your activation code:", u.logError)

	regUser, err := u.userRepository.CreateRegisteredUser(user)
	if err != nil {
		u.logError.Logger.Println("ERR:DB:CREATE USER")
		return regUser, err
	}
	return regUser, nil
}

func (u UserService) ActivateUserAccount(username string, verCode int) (bool, error) {

	var codeInfoForUsername *model.EmailVerification
	var dbEr error
	codeInfoForUsername, dbEr = u.emailRepo.GetVerificationByUsername(username)

	if dbEr != nil {
		fmt.Println(dbEr)
		fmt.Println("MAJ LAJF 1")
		u.logError.Logger.Errorf("ERR:DB:CODE DOES NOT EXIST FOR USER")
		return false, dbEr
	}
	fmt.Println("verCode:", codeInfoForUsername.VerCode)

	if codeInfoForUsername.VerCode == verCode {
		//kao dala sam kodu trajanje od 1h
		fmt.Println("kod se poklapa")
		if codeInfoForUsername.Time.Add(time.Hour).After(time.Now()) {
			fmt.Println("vreme se uklapa")
			//ako je kod ok i ako je u okviru vremena trajanja mjenjamo mu status
			user, err := u.userRepository.GetByUsername(context.TODO(), username)
			if err != nil {
				fmt.Println(err)
				fmt.Println("error u get by username kod ucitavanja usera")
				u.logError.Logger.Errorf("ERR:DB")
				return false, err
			}
			user.IsConfirmed = true
			var help string
			if user.IsConfirmed {
				help = "true"
			} else {
				help = "false"
			}
			fmt.Println("novo stanje isConfirmed : " + help)
			activated, actErr := u.userRepository.ActivateUserAccount(user)
			//editedUser, er := u.userRepository.GetByUsername(context.TODO(), username)
			if actErr != nil {
				fmt.Println(actErr)
				fmt.Println("error while activating user(repo)")
				u.logError.Logger.Errorf("ERR:WHILE ACTIVATING USER")
				return false, actErr
			}
			if !activated {
				fmt.Println("user activation failed")
				u.logError.Logger.Errorf("ERR:ACTIVATION FAILED")
				return false, errors.New("user not activated")
			}
			return true, nil

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

func (u UserService) SendCodeToRecoveryMail(username string) (bool, error) {

	user, err := u.userRepository.GetByUsername(context.TODO(), username)

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
	deleteErr := u.recoveryRepo.ClearOutRequestsForUsername(username)
	if deleteErr != nil {
		return false, deleteErr
	}
	_, e := u.recoveryRepo.CreatePasswordRecoveryRequest(&recovery)

	fmt.Println(e)
	if e != nil {
		u.logError.Logger.Println("ERR:PASS RECOVERY REQ")
		return false, e
	} else {
		u.logInfo.Logger.Infof("INFO:CREATED PASS RECOVERY")
	}

	//mzd staviti da ovo vraca bool i da ima parametar poruku i zaglavlje
	sendMailWithCourier(user.RecoveryEmail, strconv.Itoa(rn), "Password recovery code", "Here is your code:", u.logInfo)
	return true, nil
}

func (u UserService) CreateNewPassword(username string, newHashedPassword string, code string) (bool, error) {

	var passwordRecoveryRequest *model.PasswordRecoveryRequest
	var dbEr error
	passwordRecoveryRequest, dbEr = u.recoveryRepo.GetPasswordRecoveryRequestByUsername(username)

	if dbEr != nil {
		u.logError.Logger.Errorf("ERR:THERE IS NOT A PASS RECOVERY REQUEST IN DATABASE FOR USER:" + username)
		fmt.Println(dbEr)

		return false, dbEr
	}
	fmt.Println("verCode:", passwordRecoveryRequest.RecoveryCode)
	///////////////

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
				user, err := u.userRepository.GetByUsername(context.TODO(), username)
				if err != nil {
					fmt.Println(err)
					u.logError.Logger.Errorf("ERR:NO USER")
					fmt.Println("error u get by username kod ucitavanja usera")
					return false, err
				}

				fmt.Println(user.Username)

				//sacuvati izmjene korisnika,tj izmjenjen password
				changePassErr := u.userRepository.ChangePassword(user, newHashedPassword)
				if changePassErr != nil {
					fmt.Println("error pri cuvanju novog pass")
					u.logError.Logger.Errorf("ERR:SAVING NEW PASSWORD")
					return false, changePassErr
				}
				//staviti iskoristen kod na true

				//service.Repo.ActivateUserAccount(user)
				_, er := u.userRepository.GetByUsername(context.TODO(), username)
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

func checkEmailValid(email string) error {
	// check email syntax is valid
	//func MustCompile(str string) *Regexp
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
func checkEmailDomain(email string) error {
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
func sendMailWithCourier(email string, code string, subject string, body string, logErr *logger.Logger) {
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
