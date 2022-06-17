package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID `json:"id" gorm:"index:idx_name,unique"`
	Username      string    `json:"username" gorm:"unique;not null"`
	Password      string    `json:"password" gorm:"not null"`
	Email         string    `json:"email" gorm:"type-varchar(100);unique;not null"`
	PhoneNumber   string    `json:"phoneNumber" gorm:"not null"`
	FirstName     string    `json:"firstName" gorm:"not null"`
	LastName      string    `json:"lastName" gorm:"not null"`
	Gender        Gender    `json:"gender" gorm:"not null"`
	Role          Role      `json:"role" gorm:"not null"`
	IsConfirmed   bool      `json:"is_confirmed" gorm:"not null"`
	DateOfBirth   time.Time `json:"dateOfBirth" gorm:"not null"`
	RecoveryEmail string    `json:"recovery_email" gorm:"type-varchar(100);unique;not null"`
}
