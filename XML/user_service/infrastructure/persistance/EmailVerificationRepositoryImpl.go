package persistance

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"user/module/domain/model"
	"user/module/domain/repositories"
)

type EmailVerificationRepositoryImpl struct {
	db *gorm.DB
}

func NewEmailVerificationRepositoryImpl(db *gorm.DB) repositories.EmailVerificationRepository {
	return &EmailVerificationRepositoryImpl{db: db}
}

func (e EmailVerificationRepositoryImpl) CreateEmailVerification(ver *model.EmailVerification) (*model.EmailVerification, error) {
	result := e.db.Create(&ver)
	fmt.Print(result)
	return ver, result.Error
}
func (e EmailVerificationRepositoryImpl) GetVerificationByUsername(username string) (*model.EmailVerification, error) {
	verification := &model.EmailVerification{}
	if e.db.First(&verification, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return verification, nil
}
