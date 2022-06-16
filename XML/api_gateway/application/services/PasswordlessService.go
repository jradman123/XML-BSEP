package services

import (
	"context"
	"errors"
	"fmt"
	modelGateway "gateway/module/domain/model"
	"gateway/module/domain/repositories"
	"github.com/google/uuid"
	"github.com/hako/branca"
	courier "github.com/trycourier/courier-go/v2"
	"html/template"
	"log"
	"regexp"
	"time"
)

type PasswordLessService struct {
	l    *log.Logger
	cdc  *branca.Branca
	repo repositories.LoginVerificationRepository
}

var (
	ErrUnauthenticated    = errors.New("unauthenticated")
	ErrInvalidMail        = errors.New("invalid mail")
	ErrInvalidRedirectUri = errors.New("invalid redirect uri")
)

var magicLinkTmpl *template.Template

const tokenLifeSpan = time.Hour * 24 * 14
const verificationLifeSpan = time.Minute * 3

func NewPasswordLessService(l *log.Logger, repo repositories.LoginVerificationRepository, brancaKey string) *PasswordLessService {
	cdc := branca.NewBranca(brancaKey)
	cdc.SetTTL(uint32(tokenLifeSpan.Seconds()))
	return &PasswordLessService{l, cdc, repo}
}

func (s *PasswordLessService) SendMagicLink(ctx context.Context, redirectURI, origin string, user *modelGateway.User) error {
	fmt.Println("send magic link")
	badMail := BadEmail(user.Email)
	if badMail {
		s.l.Println("invalid mail")
		return ErrInvalidMail
	}

	//uri, err := url.ParseRequestURI(redirectURI)
	//if err != nil {
	//	return ErrInvalidRedirectUri
	//}

	var verificationCode string
	verificationCode = "225883"
	loginVerification := modelGateway.LoginVerification{
		ID:       uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		VerCode:  verificationCode,
		Time:     time.Now(),
	}

	ver, err := s.repo.CreateEmailVerification(&loginVerification)
	fmt.Println(ver)
	fmt.Println(err)
	//
	//magicLink, _ := url.Parse(origin)
	//magicLink.Path = "/api/auth_redirect"
	//q := magicLink.Query()
	//q.Set("verification_code", verificationCode)
	//q.Set("redirect_uri", uri.String())
	//magicLink.RawQuery = q.Encode()
	//
	//if magicLinkTmpl == nil {
	//	magicLinkTmpl, err := template.ParseFiles("template/magic-link.html")
	//	fmt.Println(magicLinkTmpl)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return errors.New("could not parse magic link mail template")
	//	}
	//}
	//
	//var mail bytes.Buffer
	//if err = magicLinkTmpl.Execute(&mail, map[string]interface{}{
	//	"MagicLink": magicLink.String(),
	//	"Minutes":   int(verificationLifeSpan.Minutes()),
	//}); err != nil {
	//	fmt.Println(err.Error())
	//	return errors.New("could not execute magic link mail template")
	//}

	defer sendMailWithCourier(user.Email, "1234", "Passwordless login", "http://localhost:9090/users/login/passwordless-auth/")

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
func BadEmail(input string) bool {
	justMail, _ := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, input)
	return !justMail
}
