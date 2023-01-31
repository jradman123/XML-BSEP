package repositories

import (
	"context"
	"github.com/google/uuid"
	"user/module/domain/model"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetByUsername(username string) (*model.User, error)
	CreateRegisteredUser(user *model.User, ctx context.Context) (*model.User, error)
	UserExists(username string) error
	GetUserSalt(username string) (string, error)
	GetUserRole(username string, ctx context.Context) (string, error)
	ActivateUserAccount(user *model.User) (bool, error)
	ChangePassword(user *model.User, password string, ctx context.Context) error
	EditUserDetails(user *model.User) (bool, error)
	ChangeProfileStatus(user *model.User) (bool, error)
	GetById(id uuid.UUID) (*model.User, error)
	UpdateEmail(user *model.User) (bool, error)
	UpdateUsername(ctx context.Context, user *model.User) (bool, error)
}
