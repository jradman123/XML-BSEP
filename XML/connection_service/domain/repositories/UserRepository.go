package repositories

import connectionModel "connection/module/domain/model"

type UserRepository interface {
	Register(userNode *connectionModel.User) (*connectionModel.User, error)
	UpdateUser(userUUID string, private bool) error
}
