package model

import (
	"time"

	"github.com/google/uuid"
)

//rola za sad samo string
type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username" gorm:"unique;not null"`
	Password     string    `json:"password" gorm:"not null"`
	Email        string    `json:"email" gorm:"typevarchar(100);unique;not null"`
	PhoneNumber  string    `json:"phoneNumber" gorm:"not null"`
	FirstName    string    `json:"firstName" gorm:"not null"`
	LastName     string    `json:"lastName" gorm:"not null"`
	Gender       Gender    `json:"gender" gorm:"not null"`
	Role         UserType  `json:"role" gorm:"not null"`
	IsConfirmed  bool      `json:"is_confirmed" gorm:"not null"`
	DateOfBirth  time.Time `json:"dateOfBirth" gorm:"not null"`
	Question     string    `json:"question" gorm:"not null"`
	HashedAnswer string    `json:"hashedAnswer" gorm:"not null"`
}
