package repositories

import (
	"context"
	"user/module/domain/model"
)

type PasswordRecoveryRequestRepository interface {
	CreatePasswordRecoveryRequest(ver *model.PasswordRecoveryRequest) (*model.PasswordRecoveryRequest, error)
	GetPasswordRecoveryRequestByUsername(username string, ctx context.Context) (*model.PasswordRecoveryRequest, error)
	ClearOutRequestsForUsername(username string) error
}
