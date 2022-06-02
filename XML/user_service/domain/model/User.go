package model

import (
	uuid "github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID    `json:"id" gorm:"index:idx_name,unique"`
	Username      string       `json:"username" gorm:"unique;not null"`
	Password      string       `json:"password" gorm:"not null"`
	Email         string       `json:"email" gorm:"type-varchar(100);unique;not null"`
	PhoneNumber   string       `json:"phoneNumber" gorm:"not null"`
	FirstName     string       `json:"firstName" gorm:"not null"`
	LastName      string       `json:"lastName" gorm:"not null"`
	Gender        Gender       `json:"gender" gorm:"not null"`
	Role          Role         `json:"role" gorm:"not null"`
	IsConfirmed   bool         `json:"is_confirmed" gorm:"not null"`
	DateOfBirth   time.Time    `json:"dateOfBirth" gorm:"not null"`
	RecoveryEmail string       `json:"recovery_email" gorm:"typevarchar(100);unique;not null"`
	Biography     string       `json:"biography"`
	Interests     []Interest   `json:"interests" gorm:"foreignKey:UserId"`
	Skills        []Skill      `json:"skills" gorm:"foreignKey:UserId"`
	Educations    []Education  `json:"educations" gorm:"foreignKey:UserId"`
	Experiences   []Experience `json:"experiences" gorm:"foreignKey:UserId"`
}
type Skill struct {
	UserId uuid.UUID
	Skill  string
}
type Interest struct {
	UserId   uuid.UUID
	Interest string
}
type Education struct {
	UserId    uuid.UUID
	Education string
}
type Experience struct {
	UserId     uuid.UUID
	Experience string
}
