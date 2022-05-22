package model

import (
	"time"

	"github.com/google/uuid"
)

type EmailVerification struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username" gorm:"unique;not null"`
	Email    string    `json:"email" gorm:"typevarchar(100);not null"`
	VerCode  int       `json:"ver_code" gorm:"not null"`
	Time     time.Time `json:"time" gorm:"not null"`
}
