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
	span1 := tracer.StartSpanFromContext(ctx, "ReadVerification")
	ver, er := s.repo.GetVerificationByCode(code)
	span1.Finish()
	return ver, er

}

func sendMailWithCourier(ctx context.Context, email string, code string, subject string, body string) {
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
		tracer.LogError(span, errors.New(err.Error()))
		fmt.Println(err)
	}
	fmt.Println(requestID)
}

func BadEmail(input string) bool {
	justMail, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, input)
	return !justMail
}
func (s *PasswordLessService) SendLink(ctx context.Context, redirectURI, origin string, user *modelGateway.User) error {
	span := tracer.StartSpanFromContext(ctx, "SendLink-Service")
	defer span.Finish()

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
	ctx = tracer.ContextWithSpan(context.Background(), span)
	span1 := tracer.StartSpanFromContext(ctx, "WriteEmailVerification")
	ver, err := s.repo.CreateEmailVerification(&loginVerification)
	span1.Finish()

	fmt.Println(ver)
	fmt.Println(err)

	defer sendMailWithCourier(ctx, user.Email, loginVerification.VerCode, "Passwordless login", "Here is your code :")

	return nil
}

func (s *PasswordLessService) PasswordlessLogin(ver *modelGateway.LoginVerification, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "PasswordlessLogin-Service")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	if ver.Time.Add(time.Minute * 3).After(time.Now()) {
		if !ver.Used {
			span1 := tracer.StartSpanFromContext(ctx, "WriteInDBThatCodeIsUsed")
			changePassErr := s.repo.UsedCode(ver)
			span1.Finish()
			if changePassErr != nil {
				tracer.LogError(span1, errors.New(changePassErr.Error()))
				return false, ErrCodeFlagUsed
			}

		} else {
			tracer.LogError(span, errors.New(ErrCodeUsed.Error()))
			return false, ErrCodeUsed
		}
	} else {
		tracer.LogError(span, errors.New(ErrCodeExpired.Error()))
		return false, ErrCodeExpired
	}
	return true, nil
}
