package persistance

import (
	"context"
	"errors"
	"fmt"
	"gateway/module/domain/model"
	"gateway/module/domain/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
	tracer "monitoring/module"
)

type TFAuthRepositoryImpl struct {
	db *gorm.DB
}

func NewTFAuthRepositoryImpl(db *gorm.DB) repositories.TFAuthRepository {
	return &TFAuthRepositoryImpl{db: db}
}

func (t TFAuthRepositoryImpl) Check2FaForUser(username string, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "check2FaForUserRepository")
	defer span.Finish()
	result := t.db.First(&model.QrCode{}, "username = ? AND is_valid = ?", username, true)
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func (t TFAuthRepositoryImpl) Enable2FaForUser(username string, secret string, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "enable2FaForUserService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
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

func (t TFAuthRepositoryImpl) GetUserSecret(username string, ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "getUserSecretRepository")
	defer span.Finish()

	var result string = ""
	t.db.Table("qr_codes").Select("secret").Where("username = ? AND is_valid = ?", username, true).Scan(&result)
	if result == "" {
		return "", errors.New("user secret not found")
	}
	return result, nil
}

func (t TFAuthRepositoryImpl) GetUserQr(username string, ctx context.Context) (model.QrCode, error) {
	span := tracer.StartSpanFromContext(ctx, "getUserQr")
	defer span.Finish()

	var result model.QrCode
	t.db.Table("qr_codes").Select("*").Where("username = ? AND is_valid = ?", username, true).Scan(&result)

	return result, nil
}

func (t TFAuthRepositoryImpl) Disable2FaForUser(username string, ctx context.Context) (bool, error) {
	span := tracer.StartSpanFromContext(ctx, "disable2FaForUserRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	qr, _ := t.GetUserQr(username, ctx)
	result := t.db.Model(&qr).Update("is_valid", false)
	fmt.Print(result)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
