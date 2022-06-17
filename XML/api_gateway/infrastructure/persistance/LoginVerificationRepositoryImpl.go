package persistance

import (
	"errors"
	"fmt"
	modelGateway "gateway/module/domain/model"
	"gateway/module/domain/repositories"
	"gorm.io/gorm"
)

type LoginVerificationRepositoryImpl struct {
	db *gorm.DB
}

func (l LoginVerificationRepositoryImpl) UsedCode(ver *modelGateway.LoginVerification) error {
	result := l.db.Model(&ver).Update("used", true)
	fmt.Print(result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (l LoginVerificationRepositoryImpl) GetVerificationByCode(code string) (*modelGateway.LoginVerification, error) {
	ver := &modelGateway.LoginVerification{}
	if l.db.First(&ver, "ver_code = ?", code).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return ver, nil
}

func NewLoginVerificationRepositoryImpl(db *gorm.DB) repositories.LoginVerificationRepository {
	return &LoginVerificationRepositoryImpl{db: db}
}

func (l LoginVerificationRepositoryImpl) CreateEmailVerification(ver *modelGateway.LoginVerification) (*modelGateway.LoginVerification, error) {
	result := l.db.Create(&ver)
	fmt.Print(result)
	return ver, result.Error
}

func (l LoginVerificationRepositoryImpl) GetVerificationByUsername(username string) (*modelGateway.LoginVerification, error) {
	verification := &modelGateway.LoginVerification{}
	if l.db.First(&verification, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return verification, nil
}
