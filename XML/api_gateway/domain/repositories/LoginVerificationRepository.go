package repositories

import "gateway/module/domain/model"

type LoginVerificationRepository interface {
	CreateEmailVerification(ver *model.LoginVerification) (*model.LoginVerification, error)
	GetVerificationByUsername(username string) (*model.LoginVerification, error)
	GetVerificationByCode(code string) (*model.LoginVerification, error)
	UsedCode(ver *model.LoginVerification) error
}
