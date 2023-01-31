package repositories

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"post/module/domain/model"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(requestUser *model.User) (user *model.User, err error)
	DeleteUser(userId uuid.UUID) (err error)
	Get(id primitive.ObjectID) (user *model.User, err error)
	GetByUserId(id uuid.UUID) (user []*model.User, err error)
	GetByUsername(username string) (user []*model.User, err error)
	ActivateUserAccount(userId uuid.UUID) (err error)
}
