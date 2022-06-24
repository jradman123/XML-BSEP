package persistance

import (
	"common/module/logger"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"time"
)

type ConnectionRepositoryImpl struct {
	db       *neo4j.Driver
	logInfo  *logger.Logger
	logError *logger.Logger
}

func NewConnectionRepositoryImpl(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger) repositories.ConnectionRepository {
	return &ConnectionRepositoryImpl{
		db:       client,
		logInfo:  logInfo,
		logError: logError,
	}
}
func (r ConnectionRepositoryImpl) CreateConnection(connection *model.Connection) (interface{}, error) {
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(connection.UserOneUID, tx) && checkIfUserExist(connection.UserTwoUID, tx) {

			records, err := tx.Run("MATCH (n:UserNode { uid: $uid}) RETURN n.status", map[string]interface{}{
				"uid": connection.UserOneUID,
			})
			connectionStatus := ""
			if err != nil {
				return nil, err
			}
			record, err := records.Single()
			if err != nil {
				return nil, err
			}

			status := record.Values[0].(string)
			dateNow := time.Now().Local().Unix()
			fmt.Println(status)
			if status == "PRIVATE" {
				connectionStatus = "REQUEST_SENT"
				_, err := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $connect_user_uuid  MATCH (u2:UserNode) WHERE u2.uid = $base_user_uuid CREATE (u1)-[r1:CONNECTION {status: $status, date: $date}]->(u2)", map[string]interface{}{
					"connect_user_uuid": connection.UserTwoUID,
					"base_user_uuid":    connection.UserOneUID,
					"status":            "REQUEST_SENT",
					"date":              dateNow,
				})

				if err != nil {
					return nil, err
				}

			} else {
				connectionStatus = "CONNECTED"
				_, err := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $connect_user_uuid  MATCH (u2:UserNode) WHERE u2.uid = $base_user_uuid CREATE (u1)-[r2:CONNECTION {status: $status, date: $date}]->(u2) CREATE (u2)-[r1:CONNECTION {status: $status, date: $date}]->(u1)", map[string]interface{}{
					"connect_user_uuid": connection.UserTwoUID,
					"base_user_uuid":    connection.UserOneUID,
					"status":            "CONNECTED",
					"date":              dateNow,
				})

				if err != nil {
					return nil, err
				}
			}

			fmt.Println(status)
			fmt.Println(connectionStatus)
			// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
			//return &pb.NewConnectionResponse{
			//	BaseUserUUID:       baseUserUuid,
			//	ConnectUserUUID:    connectUserUuid,
			//	ConnectionResponse: connectionStatus,
			//}, nil
			return nil, nil
		} else {
			//return &pb.NewConnectionResponse{
			//	BaseUserUUID:       baseUserUuid,
			//	ConnectUserUUID:    connectUserUuid,
			//	ConnectionResponse: "Connection refused",
			//}, nil
			return nil, nil
		}

	})

	if err != nil {
		return nil, err
	}

	fmt.Println(result)
	//return result.(*pb.NewConnectionResponse), nil
	return nil, nil
}
