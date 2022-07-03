package persistance

import (
	"common/module/logger"
	"connection/module/domain/dto"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
	"errors"
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
	exitingConnection, _ := r.ConnectionStatusForUsers(connection.UserOneUID, connection.UserTwoUID)
	if exitingConnection != nil {
		if exitingConnection.ConnectionStatus == "CONNECTED" || exitingConnection.ConnectionStatus == "REQUEST_SENT" {
			return &dto.ConnectionResponse{
				UserOneUID:       connection.UserOneUID,
				UserTwoUID:       connection.UserTwoUID,
				ConnectionStatus: "connection already exists",
			}, nil
		}
	}
	result, resultErr := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(connection.UserOneUID, tx) && checkIfUserExist(connection.UserTwoUID, tx) {

			records, err := tx.Run("MATCH (n:UserNode { uid: $uid}) RETURN n.status", map[string]interface{}{
				"uid": connection.UserTwoUID,
			})
			connectionStatus := ""
			if err != nil {
				return &dto.ConnectionResponse{
					UserOneUID:       "",
					UserTwoUID:       "",
					ConnectionStatus: "",
				}, err
			}
			record, err1 := records.Single()
			if err1 != nil {
				return &dto.ConnectionResponse{
					UserOneUID:       "",
					UserTwoUID:       "",
					ConnectionStatus: "",
				}, err1
			}

			status := record.Values[0].(string)
			dateNow := time.Now().Local().Unix()
			fmt.Println(status)
			if status == "PRIVATE" {
				connectionStatus = "REQUEST_SENT"
				_, err2 := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $connect_user_uuid  MATCH (u2:UserNode) WHERE u2.uid = $base_user_uuid CREATE (u1)-[r1:CONNECTION {status: $status, date: $date}]->(u2)", map[string]interface{}{
					"connect_user_uuid": connection.UserTwoUID,
					"base_user_uuid":    connection.UserOneUID,
					"status":            "REQUEST_SENT",
					"date":              dateNow,
				})

				if err2 != nil {
					return &dto.ConnectionResponse{
						UserOneUID:       "",
						UserTwoUID:       "",
						ConnectionStatus: "",
					}, err2
				}

			} else {
				connectionStatus = "CONNECTED"
				_, err3 := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $connect_user_uuid  MATCH (u2:UserNode) WHERE u2.uid = $base_user_uuid CREATE (u1)-[r2:CONNECTION {status: $status, date: $date}]->(u2) CREATE (u2)-[r1:CONNECTION {status: $status, date: $date}]->(u1)", map[string]interface{}{
					"connect_user_uuid": connection.UserTwoUID,
					"base_user_uuid":    connection.UserOneUID,
					"status":            "CONNECTED",
					"date":              dateNow,
				})

				if err3 != nil {
					return &dto.ConnectionResponse{
						UserOneUID:       "",
						UserTwoUID:       "",
						ConnectionStatus: "",
					}, err3
				}
			}
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

	if resultErr != nil {
		return &dto.ConnectionResponse{
			UserOneUID:       "",
			UserTwoUID:       "",
			ConnectionStatus: "",
		}, resultErr
	}

	return result.(*dto.ConnectionResponse), resultErr
}

func (r ConnectionRepositoryImpl) AcceptConnection(connection *model.Connection) (*dto.ConnectionResponse, error) {
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	result, resultErr := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(connection.UserOneUID, tx) && checkIfUserExist(connection.UserTwoUID, tx) {
			records, err := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet  MATCH (u1)-[r1:CONNECTION]->(u2) return r1.status", map[string]interface{}{
				"requestSender": connection.UserOneUID,
				"requestGet":    connection.UserTwoUID,
			})
			connectionStatus := ""
			if err != nil {
				return &dto.ConnectionResponse{
					UserOneUID:       "",
					UserTwoUID:       "",
					ConnectionStatus: "",
				}, err
			}
			record, err1 := records.Single()
			if err1 != nil {
				return &dto.ConnectionResponse{
					UserOneUID:       "",
					UserTwoUID:       "",
					ConnectionStatus: "",
				}, err1
			}

			status := record.Values[0].(string)
			dateNow := time.Now().Local().Unix()
			if status == "REQUEST_SENT" {
				connectionStatus = "CONNECTED"
				_, err2 := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet match  (u1)-[r1:CONNECTION {status: $status0 }]->(u2) set r1.status = $status1 , r1.date = $date  CREATE  (u2)-[r2:CONNECTION {status: $status1, date: $date}]->(u1)", map[string]interface{}{
					"requestSender": connection.UserOneUID,
					"requestGet":    connection.UserTwoUID,
					"status0":       "REQUEST_SENT",
					"status1":       "CONNECTED",
					"date":          dateNow,
				})

				if err2 != nil {
					return &dto.ConnectionResponse{
						UserOneUID:       "",
						UserTwoUID:       "",
						ConnectionStatus: "",
					}, err2
				}

			} else if status == "CONNECTED" {
				connectionStatus = "CONNECTION EXISTS"
			}

			return &dto.ConnectionResponse{
				UserOneUID:       connection.UserOneUID,
				UserTwoUID:       connection.UserTwoUID,
				ConnectionStatus: connectionStatus,
			}, nil
		} else {
			return &dto.ConnectionResponse{
				UserOneUID:       "",
				UserTwoUID:       "",
				ConnectionStatus: "",
			}, errors.New("user does not exist")

		}

	})

	if resultErr != nil {
		return &dto.ConnectionResponse{
			UserOneUID:       "",
			UserTwoUID:       "",
			ConnectionStatus: "",
		}, resultErr
	}

	return result.(*dto.ConnectionResponse), resultErr
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

		if !checkIfUserExist(userUid, tx) {
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
			node := model.User{UserUID: records.Record().Values[0].(string), Status: model.ProfileStatus(records.Record().Values[1].(string)), Username: records.Record().Values[2].(string), FirstName: records.Record().Values[3].(string), LastName: records.Record().Values[4].(string)}
			userNodes = append(userNodes, &node)
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

func (r ConnectionRepositoryImpl) ConnectionStatusForUsers(senderId string, receiverId string) (*dto.ConnectionResponse, error) {
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(senderId, tx) && checkIfUserExist(receiverId, tx) {
			records, err := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet  MATCH (u1)-[r1:CONNECTION]->(u2) return r1.status", map[string]interface{}{
				"requestSender": senderId,
				"requestGet":    receiverId,
			})
			status := ""
			if err != nil {
				return &dto.ConnectionResponse{
					UserOneUID:       "",
					UserTwoUID:       "",
					ConnectionStatus: "NO_CONNECTION :error extracting from database",
				}, err
			}
			if !records.Next() {
				records, err = tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet  MATCH (u1)-[r1:CONNECTION]->(u2) return r1.status", map[string]interface{}{
					"requestSender": receiverId,
					"requestGet":    senderId,
				})

				if !records.Next() {
					status = "NO_CONNECTION"
				} else {
					record := records.Record()
					status = record.Values[0].(string)
				}
				return &dto.ConnectionResponse{
					UserOneUID:       senderId,
					UserTwoUID:       receiverId,
					ConnectionStatus: status,
				}, nil
			} else {
				record := records.Record()
				status = record.Values[0].(string)
				return &dto.ConnectionResponse{
					UserOneUID:       senderId,
					UserTwoUID:       receiverId,
					ConnectionStatus: status,
				}, nil
			}

		} else {
			return &dto.ConnectionResponse{
				UserOneUID:       "",
				UserTwoUID:       "",
				ConnectionStatus: "",
			}, errors.New("user does not exist")
		}

	})

	if err != nil {
		return &dto.ConnectionResponse{
			UserOneUID:       "",
			UserTwoUID:       "",
			ConnectionStatus: "",
		}, errors.New("error extracting connection status")
	}

	return result.(*dto.ConnectionResponse), err
}

func (r ConnectionRepositoryImpl) BlockUser(con *model.Connection) (*dto.ConnectionResponse, error) {
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	result, resultErr := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		replayStatus := ""
		if checkIfUserExist(con.UserOneUID, tx) && checkIfUserExist(con.UserTwoUID, tx) {

			records, err := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet  MATCH (u1)-[r1:CONNECTION]->(u2) return r1.status", map[string]interface{}{
				"requestSender": con.UserOneUID,
				"requestGet":    con.UserTwoUID,
			})
			if err != nil {
				return &dto.ConnectionResponse{
					UserOneUID:       "",
					UserTwoUID:       "",
					ConnectionStatus: "",
				}, err
			}

			if records.Next() {
				status := records.Record().Values[0].(string)
				if status == "CONNECTED" {
					_, err1 := tx.Run("MATCH (u1:UserNode {uid:$requestSender})  MATCH (u2:UserNode {uid:$requestGet}) match (u1)-[r2:CONNECTION ]->(u2) match (u2)-[r1:CONNECTION ]->(u1) SET r2.status=$blockStatus , r1.status=$blockStatus", map[string]interface{}{
						"requestSender": con.UserOneUID,
						"requestGet":    con.UserTwoUID,
						"blockStatus":   "BLOCKED",
					})

					if err1 != nil {
						replayStatus = "ERROR"
						return &dto.ConnectionResponse{
							UserOneUID:       "",
							UserTwoUID:       "",
							ConnectionStatus: "",
						}, err1
					}
					replayStatus = "BLOCKED"
					return &dto.ConnectionResponse{
						UserOneUID:       con.UserOneUID,
						UserTwoUID:       con.UserTwoUID,
						ConnectionStatus: "BLOCKED",
					}, nil
				} else if status == "REQUEST_SENT" {
					dateNow := time.Now().Local().Unix()
					_, err1 := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet match (u1)-[r2:CONNECTION]->(u2) SET r2.status=$blockStatus CREATE (u2)-[r1:CONNECTION {status:$blockStatus, date: $date}]->(u1)", map[string]interface{}{
						"requestSender": con.UserOneUID,
						"requestGet":    con.UserTwoUID,
						"blockStatus":   "BLOCKED",
						"date":          dateNow,
					})

					if err1 != nil {
						replayStatus = "ERROR"
						return &dto.ConnectionResponse{
							UserOneUID:       "",
							UserTwoUID:       "",
							ConnectionStatus: "",
						}, err1
					}
					replayStatus = "BLOCKED"
					return &dto.ConnectionResponse{
						UserOneUID:       con.UserOneUID,
						UserTwoUID:       con.UserTwoUID,
						ConnectionStatus: "BLOCKED",
					}, nil
				}

			} else {
				recordsNew, errr := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet  MATCH (u2)-[r1:CONNECTION]->(u1) return r1.status", map[string]interface{}{
					"requestSender": con.UserOneUID,
					"requestGet":    con.UserTwoUID,
				})
				if errr != nil {
					replayStatus = "ERROR"
					return &dto.ConnectionResponse{
						UserOneUID:       "",
						UserTwoUID:       "",
						ConnectionStatus: "",
					}, errr
				}

				if recordsNew.Next() {
					statusNew := recordsNew.Record().Values[0].(string)
					if statusNew == "REQUEST_SENT" {
						dateNow := time.Now().Local().Unix()
						_, err1 := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet match (u2)-[r2:CONNECTION]->(u1) SET r2.status=$blockStatus CREATE (u1)-[r1:CONNECTION {status:$blockStatus, date: $date}]->(u2)", map[string]interface{}{
							"requestSender": con.UserOneUID,
							"requestGet":    con.UserTwoUID,
							"blockStatus":   "BLOCKED",
							"date":          dateNow,
						})

						if err1 != nil {
							replayStatus = "ERROR"
							return &dto.ConnectionResponse{
								UserOneUID:       "",
								UserTwoUID:       "",
								ConnectionStatus: "",
							}, err1
						}
						replayStatus = "BLOCKED"
						return &dto.ConnectionResponse{
							UserOneUID:       con.UserOneUID,
							UserTwoUID:       con.UserTwoUID,
							ConnectionStatus: "BLOCKED",
						}, nil
					}
				} else {
					dateNow := time.Now().Local().Unix()
					_, err1 := tx.Run("MATCH (u1:UserNode) WHERE u1.uid = $requestSender MATCH (u2:UserNode) WHERE u2.uid = $requestGet CREATE (u1)-[r1:CONNECTION {status:$blockStatus, date: $date}]->(u2) CREATE (u2)-[r2:CONNECTION {status:$blockStatus, date: $date}]->(u1)", map[string]interface{}{
						"requestSender": con.UserOneUID,
						"requestGet":    con.UserTwoUID,
						"blockStatus":   "BLOCKED",
						"date":          dateNow,
					})

					if err1 != nil {
						replayStatus = "ERROR"
						return &dto.ConnectionResponse{
							UserOneUID:       "",
							UserTwoUID:       "",
							ConnectionStatus: "",
						}, err1
					}
					replayStatus = "BLOCKED"
					return &dto.ConnectionResponse{
						UserOneUID:       con.UserOneUID,
						UserTwoUID:       con.UserTwoUID,
						ConnectionStatus: "BLOCKED",
					}, nil
				}
			}

		} else {
			return &dto.ConnectionResponse{
				UserOneUID:       "",
				UserTwoUID:       "",
				ConnectionStatus: "",
			}, errors.New("user does not exist")

		}
		return &dto.ConnectionResponse{
			UserOneUID:       con.UserOneUID,
			UserTwoUID:       con.UserTwoUID,
			ConnectionStatus: replayStatus,
		}, nil

	})

	if resultErr != nil {
		return &dto.ConnectionResponse{
			UserOneUID:       "",
			UserTwoUID:       "",
			ConnectionStatus: "",
		}, resultErr
	}

	return result.(*dto.ConnectionResponse), resultErr
}
