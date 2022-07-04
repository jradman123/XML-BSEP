package repositories

import (
	"github.com/google/uuid"
	"message/module/domain/model"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(requestUser *model.User) (user *model.User, err error)
	DeleteUser(userId uuid.UUID) (err error)
	GetByUsername(username string) (user []*model.User, err error)
	GetSettingsForUser(username string) (*model.NotificationSettings, error)
	ChangeSettingsForUser(username string, newSettings *model.NotificationSettings) (*model.NotificationSettings, error)
}
