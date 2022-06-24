package repositories

import connectionModel "connection/module/domain/model"

type ConnectionRepository interface {
	CreateConnection(connection *connectionModel.Connection) (interface{}, error)
}
