package persistance

import (
	"common/module/logger"
	"connection/module/domain/dto"
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

const neo4jSessionError = "Neo4j error on session.Close()"

func NewConnectionRepositoryImpl(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger) repositories.ConnectionRepository {
	return &ConnectionRepositoryImpl{
		db:       client,
		logInfo:  logInfo,
		logError: logError,
	}
}
func (r ConnectionRepositoryImpl) CreateConnection(connection *model.Connection) (*dto.ConnectionResponse, error) {
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(connection.UserOneUID, tx) && checkIfUserExist(connection.UserTwoUID, tx) {

			records, err := tx.Run("MATCH (n:UserNode { uid: $uid}) RETURN n.status", map[string]interface{}{
				"uid": connection.UserTwoUID,
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
			return &dto.ConnectionResponse{
				UserOneUID:       connection.UserOneUID,
				UserTwoUID:       connection.UserTwoUID,
				ConnectionStatus: connectionStatus,
			}, nil
		} else {
			return &dto.ConnectionResponse{
				UserOneUID:       connection.UserOneUID,
				UserTwoUID:       connection.UserTwoUID,
				ConnectionStatus: "connection refused",
			}, nil

		}

	})

	if err != nil {
		return nil, err
	}

	fmt.Println(result)
	return result.(*dto.ConnectionResponse), nil
}

func (r ConnectionRepositoryImpl) AcceptConnection(connection *model.Connection) (*dto.ConnectionResponse, error) {
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(connection.UserOneUID, tx) && checkIfUserExist(connection.UserTwoUID, tx) {
			//u ovom slucaju sender je user koji je privaatan i prihvata zahtev requestGet koji mu je prethodno poslao taj zahtjev
			records, err := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet  MATCH (u1)-[r1:CONNECTION]->(u2) return r1.status", map[string]interface{}{
				"requestSender": connection.UserOneUID,
				"requestGet":    connection.UserTwoUID,
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
			if status == "REQUEST_SENT" {
				connectionStatus = "CONNECTED"
				//
				_, err := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet match  (u1)-[r1:CONNECTION {status: $status0 }]->(u2) set r1.status = $status1 , r1.date = $date  CREATE  (u2)-[r2:CONNECTION {status: $status1, date: $date}]->(u1)", map[string]interface{}{
					"requestSender": connection.UserOneUID,
					"requestGet":    connection.UserTwoUID,
					"status0":       "REQUEST_SENT",
					"status1":       "CONNECTED",
					"date":          dateNow,
				})

				if err != nil {
					return nil, err
				}

			} else if status == "CONNECTED" {
				connectionStatus = "CONNECTION EXISTS"
				if err != nil {
					return nil, err
				}
			}

			fmt.Println(connectionStatus)
			// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
			return &dto.ConnectionResponse{
				UserOneUID:       connection.UserOneUID,
				UserTwoUID:       connection.UserTwoUID,
				ConnectionStatus: connectionStatus,
			}, nil
			//return nil, nil
		} else {
			//return &pb.NewConnectionResponse{
			//	BaseUserUUID:       "",
			//	ConnectUserUUID:    "",
			//	ConnectionResponse: "Connection refused - user not found",
			//}, nil
			return nil, nil
		}

	})

	if err != nil {
		return nil, err
	}

	fmt.Println(result)
	return result.(*dto.ConnectionResponse), nil
	//return nil, nil
}

func (r ConnectionRepositoryImpl) GetAllConnectionForUser(userUid string) (userNodes []*model.User, error1 error) {

	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if !checkIfUserExist(userUid, tx) {
			return &model.User{
				UserUID:   "",
				Status:    "",
				FirstName: "",
				LastName:  "",
				Username:  "",
			}, nil
		}

		records, err := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $userUid MATCH (u2:UserNode) WHERE not u2.uid = $userUid match (u2)-[r1:CONNECTION {status: $status}]->(u1) match (u1)-[r2:CONNECTION {status: $status }]->(u2) return u2.uid, u2.status, u2.username, u2.firstName, u2.lastName", map[string]interface{}{
			"userUid": userUid,
			"status":  "CONNECTED",
		})

		for records.Next() {
			node := model.User{UserUID: records.Record().Values[0].(string), Status: model.ProfileStatus(records.Record().Values[1].(string)), Username: records.Record().Values[2].(string), FirstName: records.Record().Values[3].(string), LastName: records.Record().Values[4].(string)}
			userNodes = append(userNodes, &node)
		}

		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		return userNodes, nil
	})
	fmt.Println(result)
	if err != nil {
		return nil, err
	}
	return userNodes, nil
}

func (r ConnectionRepositoryImpl) GetAllConnectionRequestsForUser(userUid string) (userNodes []*model.User, error1 error) {

	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		fmt.Println("UUID ")
		fmt.Println(userUid)
		if !checkIfUserExist(userUid, tx) {
			fmt.Println("NE POSTOJI")
			return &model.User{
				UserUID: "",
				Status:  "",
			}, nil
		}

		records, err := tx.Run("MATCH (u1:UserNode)   MATCH (u2:UserNode) WHERE u2.uid = $userUid match ("+
			"u2)-[r2:CONNECTION {status:$status}]->(u1) return u1.uid, u1.status, u1.username, u1.firstName, u1.lastName", map[string]interface{}{
			"userUid": userUid,
			"status":  "REQUEST_SENT",
		})

		for records.Next() {
			fmt.Println(records.Record())
			node := model.User{UserUID: records.Record().Values[0].(string), Status: model.ProfileStatus(records.Record().Values[1].(string)), Username: records.Record().Values[2].(string), FirstName: records.Record().Values[3].(string), LastName: records.Record().Values[4].(string)}
			userNodes = append(userNodes, &node)
			fmt.Println("USAO")
		}

		if err != nil {
			return nil, err
		}
		return userNodes, nil
	})
	if err != nil {
		return nil, err
	}
	return userNodes, nil
}
