package services

import (
	"common/module/logger"
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	tracer "monitoring/module"
	"os"
	"time"
	"user/module/domain/model"
)

type ApiTokenService struct {
	logInfo     *logger.Logger
	logError    *logger.Logger
	userService *UserService
}

func NewApiTokenService(logInfo *logger.Logger, logError *logger.Logger, userService *UserService) *ApiTokenService {
	return &ApiTokenService{logInfo, logError, userService}
}

func (s ApiTokenService) GenerateApiToken(user *model.User, ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "GenerateApiToken-Service")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	var claims = &ApiTokenClaims{}
	claims.Username = user.Username
	claims.Method = "ShareJobOffer"

	userRoles := "Agent"

	claims.Roles = append(claims.Roles, userRoles)
	var tokenCreationTime = time.Now().UTC()
	var tokenExpirationTime = tokenCreationTime.Add(time.Duration(720) * time.Hour)

	token, err := generateToken(claims, tokenExpirationTime, ctx)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s ApiTokenService) CheckIfHasAccess(token string) (bool, error) {

	return true, nil
}

func generateToken(claims *ApiTokenClaims, expirationTime time.Time, ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "generateToken")
	defer span.Finish()

	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	//claims.Issuer = os.Getenv("IP")
	claims.Issuer = "Dislinkt"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := []byte(os.Getenv("SECRET"))
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		tracer.LogError(span, err)
		return "", err
	}
	return tokenString, nil
}

//api token claims
type ApiTokenClaims struct {
	Username string   `json:"username,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Method   string   `json:"method,omitempty"`
	jwt.StandardClaims
	/**
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"` //url of my server
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
	*/
}

//Needs to override
type Claims interface {
	Valid() error
}

func (c ApiTokenClaims) valid() error {

	now := time.Now().UTC().Unix()

	if c.VerifyExpiresAt(now, false) == false {
		delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		return fmt.Errorf("token is expired by %v", delta)
	}
	if c.VerifyIssuedAt(now, false) == false {
		return fmt.Errorf("Token used before issued")

	}

	if c.VerifyNotBefore(now, false) == false {
		return fmt.Errorf("token is not valid yet")

	}
	return nil

}

func authorize(method string, token string) (bool, error) {

	err, claimsRoles, methodName := tokenIsValid(token)
	if err != nil {
		fmt.Println("TOKEN NOT VALID")
		return false, err
	}

	for _, claimsRole := range claimsRoles {
		if "Agent" == claimsRole {
			fmt.Println("Agent")
			return true, nil
		}
	}
	if "ShareJobOffer" == methodName {
		fmt.Println("Agent")
		return true, nil
	}
	return false, errors.New("authorization for api token failed")
}

func tokenIsValid(token string) (error, []string, string) {
	claims, err := verifyToken(token)

	if err != nil {
		return errors.New("unauthorized"), nil, ""
	}
	err = claims.valid()
	if err != nil {
		fmt.Println("CLAIMS NOT VALID")
		return errors.New("unauthorized"), nil, ""
	}
	return nil, claims.Roles, claims.Method
}

func verifyToken(token string) (*ApiTokenClaims, error) {
	claims := &ApiTokenClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		fmt.Println("Error parsing claims")
		return nil, err
	}

	return claims, nil
}
