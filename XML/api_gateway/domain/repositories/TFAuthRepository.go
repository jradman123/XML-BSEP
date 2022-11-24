package repositories

import (
	"context"
	"gateway/module/domain/model"
)

type TFAuthRepository interface {
	Check2FaForUser(username string, ctx context.Context) (bool, error)
	Enable2FaForUser(username string, secret string, ctx context.Context) (bool, error)
	Disable2FaForUser(username string, ctx context.Context) (bool, error)
	GetUserSecret(username string, ctx context.Context) (string, error)
	GetUserQr(username string, ctx context.Context) (model.QrCode, error)
}
