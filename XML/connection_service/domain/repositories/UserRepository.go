package repositories

import (
	connectionModel "connection/module/domain/model"
	"context"
)

type UserRepository interface {
	Register(userNode *connectionModel.User) (*connectionModel.User, error)
	UpdateUser(userNode *connectionModel.User) error
	GetUserId(username string, ctx context.Context) (string, error)
	ChangeProfileStatus(m *connectionModel.User) error
	UpdateUserProfessionalDetails(user *connectionModel.User, details *connectionModel.UserDetails) error
}
