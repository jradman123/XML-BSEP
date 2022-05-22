package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	Username string   `json:"username,omitempty"`
	Roles    []string `json:"roles,omitempty"`
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

func (c JwtClaims) Valid() error {

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

func TimeFunc() {
	panic("unimplemented")
}
