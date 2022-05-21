package repository

import (
	"fmt"
	"user/module/model"

	"gorm.io/gorm"
)

type EmailVerificationRepository struct {
	DB *gorm.DB
}

func (r *EmailVerificationRepository) Create(EmailVerification *model.EmailVerification) error {
	result := r.DB.Create(&EmailVerification)
	fmt.Print(result)
	return result.Error
}
