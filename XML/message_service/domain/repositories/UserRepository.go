package repositories

import (
	"context"
	"github.com/google/uuid"
	"message/module/domain/model"
)

type UserRepository interface {
	CreateUser(user *model.User, ctx context.Context) (*model.User, error)
	UpdateUser(requestUser *model.User, ctx context.Context) (user *model.User, err error)
	DeleteUser(userId uuid.UUID) (err error)
	GetByUsername(username string) (user []*model.User, err error)
	GetSettingsForUser(username string) (*model.NotificationSettings, error)
	ChangeSettingsForUser(username string, newSettings *model.NotificationSettings) (*model.NotificationSettings, error)
	GetById(userId uuid.UUID) ([]*model.User, error)
}
