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
	tracer "monitoring/module"
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
	ErrCodeFlagUsed       = errors.New("code flag used")
	ErrCodeUsed           = errors.New("code used")
	ErrCodeExpired        = errors.New("code expired")
)

const authToken = "pk_prod_0FQXVBPMDHMZ3VJ3WN6CYC12KNMH"

func NewPasswordLessService(logInfo *logger.Logger, logError *logger.Logger, repo repositories.LoginVerificationRepository) *PasswordLessService {
	return &PasswordLessService{logInfo, logError, repo}

}

func (s *PasswordLessService) GetUsernameByCode(code string, ctx context.Context) (*modelGateway.LoginVerification, error) {
	span := tracer.StartSpanFromContext(ctx, "GetUsernameByCode")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	ver, er := s.repo.GetVerificationByCode(code, ctx)
	return ver, er

}

func sendMailWithCourier(email string, code string, subject string, body string, ctx context.Context) {
	span := tracer.StartSpanFromContext(ctx, "sendMailWithCourier")
	defer span.Finish()

	client := courier.CreateClient(authToken, nil)
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
	span := tracer.StartSpanFromContext(ctx, "SendLink")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	badMail := BadEmail(user.Email)
	if badMail {

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

	ver, err := s.repo.CreateEmailVerification(&loginVerification, ctx)
	fmt.Println(ver)
	fmt.Println(err)

	defer sendMailWithCourier(user.Email, loginVerification.VerCode, "Passwordless login", "Here is your code :", ctx)

	return nil
}

func (s *PasswordLessService) PasswordlessLogin(ver *modelGateway.LoginVerification, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "PasswordlessLogin")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	if ver.Time.Add(time.Minute * 3).After(time.Now()) {
		if !ver.Used {
			changePassErr := s.repo.UsedCode(ver, ctx)
			if changePassErr != nil {
				return false, ErrCodeFlagUsed
			}

		} else {
			return false, ErrCodeUsed
		}
	} else {
		return false, ErrCodeExpired
	}
	return true, nil
}
