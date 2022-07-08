package repositories

import (
	"context"
	"github.com/google/uuid"
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
	GetById(ctx context.Context, id uuid.UUID) (*model.User, error)
	UpdateEmail(ctx context.Context, user *model.User) (bool, error)
	UpdateUsername(ctx context.Context, user *model.User) (bool, error)
}
