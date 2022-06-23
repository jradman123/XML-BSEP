package auth

import (
	"common/module/interceptor"
	myerr "gateway/module/application/errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateToken(claims *interceptor.JwtClaims) (string, time.Time, error) {

	var tokenCreationTime = time.Now().UTC()
	var tokenExpirationTime = tokenCreationTime.Add(time.Duration(5) * time.Minute)

	claims.ExpiresAt = tokenExpirationTime.Unix()
	claims.IssuedAt = tokenCreationTime.Unix()
	claims.Issuer = os.Getenv("IP")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := []byte(os.Getenv("SECRET"))
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", tokenExpirationTime, &myerr.AuthenticationError{StatusCode: 404, Err: err, Message: "Error generating token"}
	}
	return tokenString, tokenExpirationTime, nil
}
