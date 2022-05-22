package service

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/smtp"
	"regexp"
	"strconv"
	"strings"
	"time"
	"user/module/model"
	"user/module/repository"

	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

type RegisteredUserService struct {
	Repo      *repository.RegisteredUserRepository
	EmailRepo *repository.EmailVerificationRepository
}

func (service *RegisteredUserService) ActivateUserAccount(username string, verCode int) (bool, error) {

	var codeInfoForUsername *model.EmailVerification
	var dbEr error
	codeInfoForUsername, dbEr = service.EmailRepo.GetVerificationByUsername(username)

	if dbEr != nil {
		return false, dbEr
	}
	fmt.Println("verCode:", codeInfoForUsername.VerCode)

	if codeInfoForUsername.VerCode == verCode {
		//kao dala sam kodu trajanje od 1h
		if codeInfoForUsername.Time.Add(time.Hour).After(time.Now()) {
			//ako je kod ok i ako je u okviru vremena trajanja mjenjamo mu status
			user, err := service.Repo.GetByUsername(username)
			if err != nil {
				return false, err
			}
			user.IsConfirmed = true
			service.Repo.ActivateUserAccount(user)
			editedUser, er := service.Repo.GetByUsername(username)
			if er != nil {
				return false, er
			}
			if !editedUser.IsConfirmed {
				return false, errors.New("user not activated")
			}
			return true, nil

		} else {
			return false, errors.New("code expired")
		}

	} else {
		return false, errors.New("wrong code")
	}

	//OVO TEK NAKON IZVRSI PROVJERU DA LI SE KODOVI POKLAPAJU

}

func (service *RegisteredUserService) CreateRegisteredUser(username string, password string, email string, phone string, firstName string, lastName string, gender model.Gender, role model.UserType, dateOfBirth time.Time, question string, answer string) (string, error) {
	user := model.User{
		ID:           uuid.New(),
		Username:     username,
		Password:     password,
		Email:        email,
		PhoneNumber:  phone,
		FirstName:    firstName,
		LastName:     lastName,
		Gender:       gender,
		Role:         role,
		IsConfirmed:  false,
		DateOfBirth:  dateOfBirth,
		Question:     question,
		HashedAnswer: answer,
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
	mailError := sendCodeWithMail(rn, email)
	fmt.Println(mailError)
	fmt.Println("UBICU SE ")
	println(mailError)

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

// func sendConfirmationMail(email string, username string, service *RegisteredUserService) error {
// 	var err = checkEmailValid(email)
// 	if err != nil {
// 		return errors.New("email format invalid")
// 	}
// 	var er = checkEmailDomain(email)
// 	if er != nil {
// 		return errors.New("email domain invalid")
// 	}
// 	rand.Seed(time.Now().UnixNano())
// 	rn := rand.Intn(100000)
// 	emailVerification := model.EmailVerification{
// 		Username: username,
// 		Email:    email,
// 		VerCode:  rn,
// 		Time:     time.Now(),
// 	}

// 	e := service.EmailRepo.Create(&emailVerification)
// 	if e != nil {
// 		return errors.New("error saving emailVerification")
// 	}

// 	go emailVerCode(rn, email) //valjda ne cekamo da se zavrsi

// 	return nil
// }

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

func sendCodeWithMail(code int, toEmail string) error {

	msg := gomail.NewMessage()
	msg.SetHeader("From", "bespxml@gmail.com")
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", "Verification code:"+string(code))
	msg.SetBody("text/html", "<b>Welcome to DISLINKT</b>")
	//	msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer("smtp.gmail.com", 587, "bespxml@gmail.com", "ohdearLord!")

	fmt.Println("SENDING MAIL SECOND WAY HEHE ")
	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}

func emailVerCode(rn int, toEmail string) error {
	// sender data
	//from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	//password := os.Getenv("SMTPpwd")   // ex: "ieiemcjdkejspqz"
	fmt.Println("SALJEMOOOOO MEEEEEEJJJJLLLLL")
	from := "bespxml@gmail.com"
	password := "ohdearLord!"
	// receiver address privided through toEmail argument
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: Email Verification Code\r\n\r\n"
	verCode := strconv.Itoa(rn)
	fmt.Println("verCode:", verCode)
	body := "verification code: " + verCode
	fmt.Println("body:", body)
	message := []byte(subject + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	fmt.Println("message:", string(message))
	err := smtp.SendMail(address, auth, from, to, message)
	fmt.Println(err)
	fmt.Println("MEJLLL POSLAAAAT")
	return err
}

func (u *RegisteredUserService) UsernameExists(username string) bool {

	return u.Repo.UsernameExists(username)
}
