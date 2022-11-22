package repositories

import (
	"context"
	"github.com/google/uuid"
	"user/module/domain/model"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	CreateRegisteredUser(user *model.User, ctx context.Context) (*model.User, error)
	UserExists(username string, ctx context.Context) error
	GetUserSalt(username string) (string, error)
	GetUserRole(username string, ctx context.Context) (string, error)
	ActivateUserAccount(user *model.User, ctx context.Context) (bool, error)
	ChangePassword(user *model.User, password string, ctx context.Context) error
	EditUserDetails(user *model.User, ctx context.Context) (bool, error)
	ChangeProfileStatus(user *model.User) (bool, error)
	GetById(ctx context.Context, id uuid.UUID) (*model.User, error)
	UpdateEmail(ctx context.Context, user *model.User) (bool, error)
	UpdateUsername(ctx context.Context, user *model.User) (bool, error)
}
