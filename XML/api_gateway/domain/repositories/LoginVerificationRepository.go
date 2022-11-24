package repositories

import (
	"context"
	"gateway/module/domain/model"
)

type LoginVerificationRepository interface {
	CreateEmailVerification(ver *model.LoginVerification, ctx context.Context) (*model.LoginVerification, error)
	GetVerificationByUsername(username string, ctx context.Context) (*model.LoginVerification, error)
	GetVerificationByCode(code string, ctx context.Context) (*model.LoginVerification, error)
	UsedCode(ver *model.LoginVerification, ctx context.Context) error
}
