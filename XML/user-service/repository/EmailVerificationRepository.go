package repository

import (
	"errors"
	"fmt"
	"user/module/model"

	"gorm.io/gorm"
)

type EmailVerificationRepository struct {
	DB *gorm.DB
}

func NewEmailVerificationRepository(db *gorm.DB) EmailVerificationRepository {
	return EmailVerificationRepository{DB: db}
}

func (r *EmailVerificationRepository) Create(EmailVerification *model.EmailVerification) error {
	result := r.DB.Create(&EmailVerification)
	fmt.Print(result)
	return result.Error
}

func (r *EmailVerificationRepository) GetVerificationByUsername(username string) (*model.EmailVerification, error) {
	verifiction := &model.EmailVerification{}
	if r.DB.First(&verifiction, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return verifiction, nil
}
