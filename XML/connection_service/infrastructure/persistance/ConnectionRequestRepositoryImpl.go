package persistance

import (
	"common/module/logger"
	"connection/module/domain/repositories"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ConnectionRequestRepositoryImpl struct {
	db       *neo4j.Driver
	logInfo  *logger.Logger
	logError *logger.Logger
}

func NewConnectionRequestRepositoryImpl(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger) repositories.ConnectionRequestRepository {
	return &ConnectionRequestRepositoryImpl{
		db:       client,
		logInfo:  logInfo,
		logError: logError,
	}
}
