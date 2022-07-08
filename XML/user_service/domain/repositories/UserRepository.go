package repositories

import (
	"context"
	"user/module/domain/model"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	CreateRegisteredUser(user *model.User) (*model.User, error)
	UserExists(username string) error
	GetUserSalt(username string) (string, error)
	GetUserRole(username string) (string, error)
	ActivateUserAccount(user *model.User) (bool, error)
	ChangePassword(user *model.User, password string) error
	EditUserDetails(user *model.User) (bool, error)
	ChangeProfileStatus(user *model.User) (bool, error)
}
