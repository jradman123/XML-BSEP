package repositories

import (
	"context"
	"user/module/domain/model"
)

type EmailVerificationRepository interface {
	CreateEmailVerification(ver *model.EmailVerification, ctx context.Context) (*model.EmailVerification, error)
	GetVerificationByUsername(username string, ctx context.Context) ([]model.EmailVerification, error)
}
