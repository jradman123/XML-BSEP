package repository

import (
	"errors"
	"fmt"
	"user/module/model"

	"gorm.io/gorm"
)

type PasswordRecoveryRepository struct {
	DB *gorm.DB
}

func NewPasswordRecoveryRepository(db *gorm.DB) PasswordRecoveryRepository {
	return PasswordRecoveryRepository{DB: db}
}

func (r *PasswordRecoveryRepository) Create(PasswordRecoveryRequest *model.PasswordRecoveryRequest) error {
	result := r.DB.Create(&PasswordRecoveryRequest)
	fmt.Print(result)
	return result.Error
}

func (r *PasswordRecoveryRepository) GetRequestByUsername(username string) (*model.PasswordRecoveryRequest, error) {
	recovery := &model.PasswordRecoveryRequest{}
	if r.DB.First(&recovery, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return recovery, nil
}
