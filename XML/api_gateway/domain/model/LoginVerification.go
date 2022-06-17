package model

import (
	"github.com/google/uuid"
	"time"
)

type LoginVerification struct {
	ID       uuid.UUID `json:"id" gorm:"index:idx_name,unique"`
	Username string    `json:"username" gorm:"not null"`
	Email    string    `json:"email" gorm:"type-varchar(100);not null"`
	VerCode  string    `json:"ver_code" gorm:"not null"`
	Time     time.Time `json:"time" gorm:"not null"`
	Used     bool      `json:"used" gorm:"not null"`
}
