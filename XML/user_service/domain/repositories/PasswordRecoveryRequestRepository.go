package repositories

import (
	"user/module/domain/model"
)

type PasswordRecoveryRequestRepository interface {
	CreatePasswordRecoveryRequest(ver *model.PasswordRecoveryRequest) (*model.PasswordRecoveryRequest, error)
	GetPasswordRecoveryRequestByUsername(username string) (*model.PasswordRecoveryRequest, error)
	ClearOutRequestsForUsername(username string) error
}
