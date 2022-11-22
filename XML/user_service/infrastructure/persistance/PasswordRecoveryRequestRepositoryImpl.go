package persistance

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	tracer "monitoring/module"
	"user/module/domain/model"
	"user/module/domain/repositories"
)

type PasswordRecoveryRequestRepositoryImpl struct {
	db *gorm.DB
}

func (p PasswordRecoveryRequestRepositoryImpl) ClearOutRequestsForUsername(username string) error {
	req, _ := p.GetPasswordRecoveryRequestByUsername(username, context.TODO())
	if req != nil {
		result := p.db.Delete(&model.PasswordRecoveryRequest{}, req.ID)
		if result.Error != nil {
			return result.Error
		}
		return nil
	}
	return nil
}

func NewPasswordRecoveryRequestRepositoryImpl(db *gorm.DB) repositories.PasswordRecoveryRequestRepository {
	return &PasswordRecoveryRequestRepositoryImpl{db: db}
}

func (p PasswordRecoveryRequestRepositoryImpl) CreatePasswordRecoveryRequest(ver *model.PasswordRecoveryRequest) (*model.PasswordRecoveryRequest, error) {
	result := p.db.Create(&ver)
	fmt.Print(result)
	return ver, result.Error
}

func (p PasswordRecoveryRequestRepositoryImpl) GetPasswordRecoveryRequestByUsername(username string, ctx context.Context) (*model.PasswordRecoveryRequest, error) {
	span := tracer.StartSpanFromContext(ctx, "getPasswordRecoveryRequestByUsername")
	defer span.Finish()

	recovery := &model.PasswordRecoveryRequest{}
	if p.db.First(&recovery, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return recovery, nil
}
