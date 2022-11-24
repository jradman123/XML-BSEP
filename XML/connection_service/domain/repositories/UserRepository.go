package repositories

import (
	connectionModel "connection/module/domain/model"
	"context"
)

type UserRepository interface {
	Register(userNode *connectionModel.User, ctx context.Context) (*connectionModel.User, error)
	UpdateUser(userNode *connectionModel.User, ctx context.Context) error
	GetUserId(username string, ctx context.Context) (string, error)
	ChangeProfileStatus(m *connectionModel.User, ctx context.Context) error
	UpdateUserProfessionalDetails(user *connectionModel.User, details *connectionModel.UserDetails, ctx context.Context) error
}
