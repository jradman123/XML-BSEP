package handlers

import (
	"encoding/base32"
	dgoogauth "github.com/dgryski/dgoogauth"
)

func NewOTPConfig(userSecret string) *dgoogauth.OTPConfig {
	secret := []byte{'H', 'e', 'l', 'l', 'o', '!', 0xDE, 0xAD, 0xBE, 0xEF}
	secretBase32 := base32.StdEncoding.EncodeToString(secret)
	return &dgoogauth.OTPConfig{
		Secret:      secretBase32,
		WindowSize:  1,
		HotpCounter: 0,
		// UTC:         true,
		DisallowReuse: make([]int, 0),
		ScratchCodes:  make([]int, 0),
	}
}
