package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	my_err "user/module/errors"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(claims *JwtClaims, expirationTime time.Time) (string, error) {

	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = os.Getenv("IP")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	mySigningKey := []byte(os.Getenv("SECRET"))
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", &my_err.RequestError{StatusCode: 404, Err: err, Message: "Error generating token"}
	}
	return tokenString, nil
}
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func TokenIsValid(tokenString string) (error, string) {
	claims, err := VerifyToken(tokenString)

	if err != nil {
		return &my_err.RequestError{
			StatusCode: http.StatusUnauthorized,
			Err:        err,
			Message:    "Token wasn't verified",
		}, ""
	}
	err = claims.Valid()
	if err != nil {
		return &my_err.RequestError{
			StatusCode: http.StatusUnauthorized,
			Err:        err,
			Message:    "Calims are not valid",
		}, ""
	}
	return nil, claims.Username
}

func VerifyToken(tokenString string) (*JwtClaims, error) {
	claims := &JwtClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}
