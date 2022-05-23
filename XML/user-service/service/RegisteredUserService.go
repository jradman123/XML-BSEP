package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
	"user/module/model"
	"user/module/repository"

	"github.com/google/uuid"
	"github.com/trycourier/courier-go/v2"
)

type RegisteredUserService struct {
	Repo         *repository.RegisteredUserRepository
	EmailRepo    *repository.EmailVerificationRepository
	RecoveryRepo *repository.PasswordRecoveryRepository
}

func (service *RegisteredUserService) CreateNewPassword(username string, newPassword string, code string) (bool, error) {

	var passwordRecoveryRequest *model.PasswordRecoveryRequest
	var dbEr error
	passwordRecoveryRequest, dbEr = service.RecoveryRepo.GetRequestByUsername(username)

	if dbEr != nil {
		fmt.Println(dbEr)

		return false, dbEr
	}
	fmt.Println("verCode:", passwordRecoveryRequest.RecoveryCode)
	///////////////

	var codeInt, convErr = strconv.Atoi(code)
	if convErr != nil {
		return false, errors.New("error converting code to int")
	}
	if passwordRecoveryRequest.RecoveryCode == codeInt {
		//kao dala sam kodu trajanje od 1h
		fmt.Println("kod se poklapa")
		if passwordRecoveryRequest.Time.Add(time.Minute * 3).After(time.Now()) {
			fmt.Println("vreme se uklapa")
			//ako je kod ok i ako je u okviru vremena trajanja mjenjamo mu status
			user, err := service.Repo.GetByUsername(username)
			if err != nil {
				fmt.Println(err)
				fmt.Println("error u get by username kod ucitavanja usera")
				return false, err
			}

			fmt.Println(user.Username)
			//CHANGE PASSWORD

			// var hashedSaltedPassword = ""
			// validPassword := u.passwordUtil.IsValidPassword(password)

			// if validPassword {
			// 	//PASSWORD SALT
			// 	//salt, password = u.passwordUtil.GeneratePasswordWithSalt(newUser.Password)
			// 	//cuvamo password kao hash neki
			// 	pass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
			// 	if err != nil {
			// 		fmt.Println(err)
			// 		err := ErrorResponse{
			// 			Err: "Password Encryption  failed",
			// 		}
			// 		json.NewEncoder(rw).Encode(err)
			// 	}

			// 	hashedSaltedPassword = string(pass)

			// } else {
			// 	fmt.Println("Password format is not valid!")
			// 	http.Error(rw, "Password format is not valid! error:"+err.Error(), http.StatusBadRequest) //400
			// 	return
			// }

			/////////

			//service.Repo.ActivateUserAccount(user)
			_, er := service.Repo.GetByUsername(username)
			if er != nil {
				fmt.Println(er)
				fmt.Println("FAK MAJ LAJF 2")
				return false, er
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
	////////////////////
	return true, nil
}

func (service *RegisteredUserService) SendCodeToRecoveryMail(username string) bool {
	user, err := service.Repo.GetByUsername(username)

	if err != nil {
		return false
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

	e := service.RecoveryRepo.Create(&recovery)

	fmt.Println(e)
	if e != nil {
		return false
	}

	//mzd staviti da ovo vraca bool i da ima parametar poruku i zaglavlje
	sendMailWithCourier(user.RecoveryEmail, strconv.Itoa(rn), "Password recovery code", "Here is your code:")
	return true

}

func (service *RegisteredUserService) ActivateUserAccount(username string, verCode int) (bool, error) {

	var codeInfoForUsername *model.EmailVerification
	var dbEr error
	codeInfoForUsername, dbEr = service.EmailRepo.GetVerificationByUsername(username)

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
			user, err := service.Repo.GetByUsername(username)
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
			service.Repo.ActivateUserAccount(user)
			editedUser, er := service.Repo.GetByUsername(username)
			if er != nil {
				fmt.Println(er)
				fmt.Println("FAK MAJ LAJF 2")
				return false, er
			}
			if !editedUser.IsConfirmed {
				fmt.Println("FAK MAJ LAJF 3")
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

func (service *RegisteredUserService) CreateRegisteredUser(username string, password string, email string, phone string, firstName string, lastName string, gender model.Gender, role model.UserType, dateOfBirth time.Time, recoveryMail string) (string, error) {
	user := model.User{
		ID:            uuid.New(),
		Username:      username,
		Password:      password,
		Email:         email,
		PhoneNumber:   phone,
		FirstName:     firstName,
		LastName:      lastName,
		Gender:        gender,
		Role:          role,
		IsConfirmed:   false,
		DateOfBirth:   dateOfBirth,
		RecoveryEmail: recoveryMail,
	}
	mail, err := service.Repo.CreateRegisteredUser(&user)
	if err != nil {
		return mail, err
	}

	//TODO: send confirmation mail
	//	sendConfirmationMail(user.Email, user.Username, service)

	var er = checkEmailValid(email)
	if er != nil {
		return mail, errors.New("email format invalid")
	}
	var domEr = checkEmailDomain(email)
	if domEr != nil {
		return mail, errors.New("email domain invalid")
	}
	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(100000)
	emailVerification := model.EmailVerification{
		Username: username,
		Email:    email,
		VerCode:  rn,
		Time:     time.Now(),
	}
	fmt.Println(emailVerification)

	e := service.EmailRepo.Create(&emailVerification)
	fmt.Println(e)
	if e != nil {
		return mail, errors.New("error saving emailVerification")
	}

	//emailVerCode(rn, email) //valjda ne cekamo da se zavrsi
	// mailError := sendCodeWithMail(rn, email)
	// fmt.Println(mailError)
	// fmt.Println("UBICU SE ")
	// println(mailError)

	sendMailWithCourier(email, strconv.Itoa(rn), "Activation code", "Welcome to Dislinkt! Here is your activation code:")

	//TODO: send confirmation mail
	//koriste redis bazu gdje privremeno cuvaju zahteve za registraciju
	//a ovo sto serijalizuju mzd kasnije tek upisuju u bazu
	//expiration  := 1000000000 * 3600 * 2 //2h
	//serializedUser, err := serialize(user)
	// err = s.RedisUsecase.AddKeyValueSet(context, redisKey, serializedUser, time.Duration(expiration));
	// if err != nil {
	// 	return err
	// }
	// confirmationCode := helpers.RandomStringGenerator(8)
	// hashedConfirmationCode, err := helpers.Hash(confirmationCode)
	// user.hashedConfirmationCode = hashedConfirmationCode
	// if err != nil {
	// 	s.logger.Logger.Errorf("error while registering user, error %v\n", err)
	// 	return err
	// }
	return mail, nil
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

func (u *RegisteredUserService) UsernameExists(username string) bool {

	return u.Repo.UsernameExists(username)
}
