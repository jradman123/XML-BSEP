package repositories

import "user/module/domain/model"

type EmailVerificationRepository interface {
	CreateEmailVerification(ver *model.EmailVerification) (*model.EmailVerification, error)
	GetVerificationByUsername(username string) ([]model.EmailVerification, error)
}
