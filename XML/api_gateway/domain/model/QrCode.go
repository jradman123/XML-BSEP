package model

import "github.com/google/uuid"

type QrCode struct {
	ID       uuid.UUID `json:"id" gorm:"index:idx_name,unique"`
	Secret   string    `json:"secret" gorm:"unique;not null"`
	Username string    `json:"username" gorm:"not null"`
	IsValid  bool      `json:"is_valid" gorm:"not null"`
}
