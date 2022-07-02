package model

import (
	uuid "github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID     `json:"id" gorm:"index:idx_name,unique"`
	Username      string        `json:"username" gorm:"unique;not null"`
	Password      string        `json:"password" gorm:"not null"`
	Email         string        `json:"email" gorm:"type-varchar(100);unique;not null"`
	PhoneNumber   string        `json:"phoneNumber" gorm:"not null"`
	FirstName     string        `json:"firstName" gorm:"not null"`
	LastName      string        `json:"lastName" gorm:"not null"`
	Gender        Gender        `json:"gender" gorm:"not null"`
	Role          Role          `json:"role" gorm:"not null"`
	IsConfirmed   bool          `json:"is_confirmed" gorm:"not null"`
	DateOfBirth   time.Time     `json:"dateOfBirth" gorm:"not null"`
	RecoveryEmail string        `json:"recovery_email" gorm:"typevarchar(100);unique;not null"`
	Biography     string        `json:"biography"`
	Interests     []Interest    `json:"interests"`
	Skills        []Skill       `json:"skills"`
	Educations    []Education   `json:"educations"`
	Experiences   []Experience  `json:"experiences"`
	ProfileStatus ProfileStatus `json:"profileStatus"`
}
type Skill struct {
	ID     int
	UserID uuid.UUID
	Skill  string
}
type Interest struct {
	ID       int
	UserID   uuid.UUID
	Interest string
}
type Education struct {
	ID        int
	UserID    uuid.UUID
	Education string
}
type Experience struct {
	ID         int
	UserID     uuid.UUID
	Experience string
}
type ProfileStatus string

const (
	Private ProfileStatus = "PRIVATE"
	Public  ProfileStatus = "PUBLIC"
)
