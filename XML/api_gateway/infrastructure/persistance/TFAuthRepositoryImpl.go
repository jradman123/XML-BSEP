package persistance

import (
	"errors"
	"fmt"
	"gateway/module/domain/model"
	"gateway/module/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TFAuthRepositoryImpl struct {
	db *gorm.DB
}

func NewTFAuthRepositoryImpl(db *gorm.DB) repositories.TFAuthRepository {
	return &TFAuthRepositoryImpl{db: db}
}

func (t TFAuthRepositoryImpl) Check2FaForUser(username string) (bool, error) {
	result := t.db.First(&model.QrCode{}, "username = ? AND is_valid = ?", username, true)
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (t TFAuthRepositoryImpl) Enable2FaForUser(username string, secret string) (bool, error) {

	check, _ := t.Check2FaForUser(username)
	if check == true {
		return false, errors.New("two factor authentication already enabled ")
	}

	qr := model.QrCode{
		ID:       uuid.New(),
		Secret:   secret,
		Username: username,
		IsValid:  true,
	}
	result := t.db.Create(&qr)
	fmt.Print(result)
	return true, nil
}

func (t TFAuthRepositoryImpl) GetUserSecret(username string) (string, error) {
	var result string = ""
	t.db.Table("qr_codes").Select("secret").Where("username = ? AND is_valid = ?", username, true).Scan(&result)
	if result == "" {
		return "", errors.New("user secret not found")
	}
	return result, nil
}

func (t TFAuthRepositoryImpl) GetUserQr(username string) (model.QrCode, error) {
	var result model.QrCode
	t.db.Table("qr_codes").Select("*").Where("username = ? AND is_valid = ?", username, true).Scan(&result)

	return result, nil
}

func (t TFAuthRepositoryImpl) Disable2FaForUser(username string) (bool, error) {
	qr, _ := t.GetUserQr(username)
	result := t.db.Model(&qr).Update("is_valid", false)
	fmt.Print(result)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
