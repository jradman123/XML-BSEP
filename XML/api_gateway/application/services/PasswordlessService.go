package services

import (
	"common/module/logger"
	"context"
	"errors"
	"fmt"
	modelGateway "gateway/module/domain/model"
	"gateway/module/domain/repositories"
	"github.com/google/uuid"
	courier "github.com/trycourier/courier-go/v2"
	"math/rand"
	"regexp"
	"time"
)

type PasswordLessService struct {
	logInfo  *logger.Logger
	logError *logger.Logger
	repo     repositories.LoginVerificationRepository
}

var (
	ErrUnauthenticated    = errors.New("unauthenticated")
	ErrInvalidMail        = errors.New("invalid mail")
	ErrInvalidRedirectUri = errors.New("invalid redirect uri")
	ErrInvalidVerCode     = errors.New("invalid verification code")
)

const tokenLifeSpan = time.Hour * 24 * 14

func NewPasswordLessService(logInfo *logger.Logger, logError *logger.Logger, repo repositories.LoginVerificationRepository) *PasswordLessService {

	return &PasswordLessService{logInfo, logError, repo}
}

func (s *PasswordLessService) GetUsernameByCode(code string) (*modelGateway.LoginVerification, error) {
	ver, er := s.repo.GetVerificationByCode(code)
	return ver, er
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

func BadEmail(input string) bool {
	justMail, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, input)
	return !justMail
}
func (s *PasswordLessService) SendLink(ctx context.Context, redirectURI, origin string, user *modelGateway.User) error {
	fmt.Println("send magic link")
	badMail := BadEmail(user.Email)
	if badMail {
		fmt.Println("invalid mail")
		return ErrInvalidMail
	}

	var verificationCode string
	rand.Seed(time.Now().UnixNano())
	rn := rand.Intn(100000)
	verificationCode = fmt.Sprintf("%d", rn)

	loginVerification := modelGateway.LoginVerification{
		ID:       uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		VerCode:  verificationCode,
		Time:     time.Now(),
		Used:     false,
	}

	ver, err := s.repo.CreateEmailVerification(&loginVerification)
	fmt.Println(ver)
	fmt.Println(err)

	defer sendMailWithCourier(user.Email, loginVerification.VerCode, "Passwordless login", "Here is your code :")

	return nil
}

func (s *PasswordLessService) PasswordlesLogin(ver *modelGateway.LoginVerification) (bool, error) {

	if ver.Time.Add(time.Minute * 3).After(time.Now()) {
		if !ver.Used {
			fmt.Println("vreme se uklapa")

			changePassErr := s.repo.UsedCode(ver)
			if changePassErr != nil {
				fmt.Println("error pri izmeni koda")
				return false, errors.New("error while setind code flag used")
			}

		} else {
			fmt.Println("kod iskoristen")
			return false, errors.New("code used")
		}
	} else {
		fmt.Println("istekao kod")
		return false, errors.New("code expired")
	}

	return true, nil
}
