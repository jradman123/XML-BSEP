package auth

import (
	"common/module/interceptor"
	myerr "gateway/module/application/errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateToken(claims *interceptor.JwtClaims, expirationTime time.Time) (string, error) {

	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = os.Getenv("IP")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := []byte(os.Getenv("SECRET"))
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", &myerr.AuthenticationError{StatusCode: 404, Err: err, Message: "Error generating token"}
	}
	return tokenString, nil
}
