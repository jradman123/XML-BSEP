package model

import (
	"github.com/google/uuid"
	"time"
)

type EmailVerification struct {
	ID       uuid.UUID `json:"id" gorm:"index:idx_name,unique"`
	Username string    `json:"username" gorm:"not null"`
	Email    string    `json:"email" gorm:"type-varchar(100);not null"`
	VerCode  int       `json:"ver_code" gorm:"not null"`
	Time     time.Time `json:"time" gorm:"not null"`
}
