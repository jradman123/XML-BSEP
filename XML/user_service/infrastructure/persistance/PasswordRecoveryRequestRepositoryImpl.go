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

func (p PasswordRecoveryRequestRepositoryImpl) ClearOutRequestsForUsername(username string, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "ClearOutRequestsForUsername")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	req, _ := p.GetPasswordRecoveryRequestByUsername(username, ctx)
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

func (p PasswordRecoveryRequestRepositoryImpl) CreatePasswordRecoveryRequest(ver *model.PasswordRecoveryRequest, ctx context.Context) (*model.PasswordRecoveryRequest, error) {
	span := tracer.StartSpanFromContext(ctx, "CreatePasswordRecoveryRequest")
	defer span.Finish()
	result := p.db.Create(&ver)
	fmt.Print(result)
	return ver, result.Error
}

func (p PasswordRecoveryRequestRepositoryImpl) GetPasswordRecoveryRequestByUsername(username string, ctx context.Context) (*model.PasswordRecoveryRequest, error) {
	span := tracer.StartSpanFromContext(ctx, "GetPasswordRecoveryRequestByUsername")
	defer span.Finish()

	recovery := &model.PasswordRecoveryRequest{}
	if p.db.First(&recovery, "username = ?", username).RowsAffected == 0 {
		return nil, errors.New("user not found")

	}
	return recovery, nil
}
