package services

import (
	"context"
	"crypto/rand"
	"encoding/base32"
	"errors"
	"gateway/module/domain/repositories"
	"github.com/dgryski/dgoogauth"
	"log"
	tracer "monitoring/module"
)

type TFAuthService struct {
	l          *log.Logger
	repository repositories.TFAuthRepository
}

func NewTFAuthService(l *log.Logger, repository repositories.TFAuthRepository) *TFAuthService {
	return &TFAuthService{l, repository}
}
func NewOTPConfig(secretBase32 string) *dgoogauth.OTPConfig {
	return &dgoogauth.OTPConfig{
		Secret:      secretBase32,
		WindowSize:  1,
		HotpCounter: 0,
		// UTC:         true,
		DisallowReuse: make([]int, 0),
		ScratchCodes:  make([]int, 0),
	}
}

var (
	TwoFactorEnabled = errors.New("two factor authentication already enabled ")
)

func GenerateNewUserSecret() []byte {
	secret := make([]byte, 10)
	_, err := rand.Read(secret)
	if err != nil {
		panic(err)
	}

	return secret
}

func (u TFAuthService) Check2FaForUser(username string, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "Check2FaForUserService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	res, err := u.repository.Check2FaForUser(username, ctx)

	if err != nil {
		return false, err
	}
	return res, nil
}

func (u TFAuthService) Enable2FaForUser(username string, ctx context.Context) (bool, string, error) {
	span := tracer.StartSpanFromContext(ctx, "Enable2FaForUserService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	secrets := GenerateNewUserSecret()
	secret := base32.StdEncoding.EncodeToString(secrets)
	check, _ := u.repository.Check2FaForUser(username, ctx)
	if check == true {
		return false, "", TwoFactorEnabled
	}

	res, err := u.repository.Enable2FaForUser(username, secret, ctx)
	if err != nil {
		return false, "", err
	}

	twofa := NewOTPConfig(secret)
	uri := twofa.ProvisionURI(username)
	//log.Println("This is URI: " + uri)
	// No more writing to file
	//err = qrcode.WriteFile(uri, qrcode.Medium, 256, "qr2.png")
	//if err != nil {
	//	return false, "", err
	//}

	if err != nil {
		return false, "", err
	}
	return res, uri, nil

}

func (u TFAuthService) Disable2FaForUser(username string, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "Disable2FaForUserService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	res, err := u.repository.Disable2FaForUser(username, ctx)

	if err != nil {
		return false, err
	}
	return res, nil

}

func (u TFAuthService) GetUserSecret(username string, ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "GetUserSecretService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	res, err := u.repository.GetUserSecret(username, ctx)

	if err != nil {
		return "", err
	}
	return res, nil

}
