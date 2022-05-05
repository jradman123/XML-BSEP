package model

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username" gorm:"unique;not null"`
	Password    string    `json:"password" gorm:"not null"`
	Email       string    `json:"email" gorm:"typevarchar(100);unique;not null"`
	PhoneNumber string    `json:"phoneNumber" gorm:"not null"`
	FirstName   string    `json:"firstName" gorm:"not null"`
	LastName    string    `json:"lastName" gorm:"not null"`
	Gender      Gender    `json:"gender" gorm:"not null"`
}
