package repositories

import (
	"context"
	"gateway/module/domain/model"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	UserExists(username string) error
	GetUserSalt(username string) (string, error)
	GetUserRole(username string) (string, error)
}
