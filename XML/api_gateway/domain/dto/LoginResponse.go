package dto

import "time"

type LogInResponseDto struct {
	Token          string    `json:"token"`
	Role           string    `json:"role"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	ExpirationTime time.Time `json:"expirationTime"`
}
