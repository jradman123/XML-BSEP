package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/trycourier/courier-go/v2"
	"log"
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
	l              *log.Logger
	userRepository repositories.UserRepository
	emailRepo      repositories.EmailVerificationRepository
}

func NewUserService(l *log.Logger, repository repositories.UserRepository, emailRepo repositories.EmailVerificationRepository) *UserService {
	return &UserService{l, repository, emailRepo}
}

func (u UserService) GetUsers() ([]model.User, error) {

	users, err := u.userRepository.GetUsers()
	if err != nil {
		return nil, errors.New("Cant get users")
	}

	return users, err

}

func (u UserService) GetByUsername(ctx context.Context, username string) (*model.User, error) {

	user, err := u.userRepository.GetByUsername(ctx, username)

	if err != nil {
		u.l.Println("Invalid username")
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
		return nil, errors.New("email format invalid")
	}
	var domEr = checkEmailDomain(user.Email)
	if domEr != nil {
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
		return nil, errors.New("error saving emailVerification")
	}

	sendMailWithCourier(user.Email, strconv.Itoa(rn), "Activation code", "Welcome to Dislinkt! Here is your activation code:")

	regUser, err := u.userRepository.CreateRegisteredUser(user)
	if err != nil {
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
		fmt.Println("FAK MAJ LAJF 1")
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
				return false, actErr
			}
			if !activated {
				fmt.Println("user activation failed")
				return false, errors.New("user not activated")
			}
			return true, nil

		} else {
			fmt.Println("istekao kod")
			return false, errors.New("code expired")
		}

	} else {
		fmt.Println("ne valjda kod")
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
func sendMailWithCourier(email string, code string, subject string, body string) {
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
		fmt.Println(err)
	}
	fmt.Println(requestID)
}
