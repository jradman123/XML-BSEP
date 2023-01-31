package repositories

import (
	"gateway/module/domain/model"
)

type TFAuthRepository interface {
	Check2FaForUser(username string) (bool, error)
	Enable2FaForUser(username string, secret string) (bool, error)
	Disable2FaForUser(username string) (bool, error)
	GetUserSecret(username string) (string, error)
	GetUserQr(username string) (model.QrCode, error)
}
