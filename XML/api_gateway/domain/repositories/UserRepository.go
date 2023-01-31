package repositories

import (
	"gateway/module/domain/model"
)

type UserRepository interface {
	GetByUsername(username string) (*model.User, error)
	UserExists(username string) error
	GetUserSalt(username string) (string, error)
	GetUserRole(username string) (string, error)
}
