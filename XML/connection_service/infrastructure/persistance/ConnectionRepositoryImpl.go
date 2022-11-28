package persistance

import (
	"common/module/logger"
	"connection/module/domain/dto"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
	"context"
	"errors"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go.mongodb.org/mongo-driver/bson/primitive"
	tracer "monitoring/module"
	"strings"
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
func (r ConnectionRepositoryImpl) CreateConnection(connection *model.Connection, ctx context.Context) (*dto.ConnectionResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CreateConnectionRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)
	exitingConnection, _ := r.ConnectionStatusForUsers(connection.UserOneUID, connection.UserTwoUID, ctx)
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

		if checkIfUserExist(connection.UserOneUID, tx, ctx) && checkIfUserExist(connection.UserTwoUID, tx, ctx) {

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

func (r ConnectionRepositoryImpl) AcceptConnection(connection *model.Connection, ctx context.Context) (*dto.ConnectionResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "AcceptConnectionRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	result, resultErr := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(connection.UserOneUID, tx, ctx) && checkIfUserExist(connection.UserTwoUID, tx, ctx) {
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

func (r ConnectionRepositoryImpl) GetAllConnectionForUser(userUid string, ctx context.Context) (userNodes []*model.User, error1 error) {
	span := tracer.StartSpanFromContext(ctx, "GetAllConnectionForUserRepository")
	defer span.Finish()
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if !checkIfUserExist(userUid, tx, ctx) {
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

func (r ConnectionRepositoryImpl) GetAllConnectionRequestsForUser(userUid string, ctx context.Context) (userNodes []*model.User, error1 error) {
	span := tracer.StartSpanFromContext(ctx, "GetAllConnectionRequestsForUserRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if !checkIfUserExist(userUid, tx, ctx) {
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

func (r ConnectionRepositoryImpl) ConnectionStatusForUsers(senderId string, receiverId string, ctx context.Context) (*dto.ConnectionResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "ConnectionStatusForUsersService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if checkIfUserExist(senderId, tx, ctx) && checkIfUserExist(receiverId, tx, ctx) {
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

func (r ConnectionRepositoryImpl) BlockUser(con *model.Connection, ctx context.Context) (*dto.ConnectionResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "BlockUserRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	result, resultErr := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		replayStatus := ""
		if checkIfUserExist(con.UserOneUID, tx, ctx) && checkIfUserExist(con.UserTwoUID, tx, ctx) {

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

func (r ConnectionRepositoryImpl) GetRecommendedNewConnections(userId string, ctx context.Context) (userNodes []*model.User, error1 error) {
	span := tracer.StartSpanFromContext(ctx, "GetRecommendedNewConnectionsRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if !checkIfUserExist(userId, tx, ctx) {
			return &model.User{
				UserUID: "",
				Status:  "",
			}, nil
		}

		records, err := tx.Run("MATCH (user:UserNode {uid:$userUid})<-[:CONNECTION {status:$status}]->(friend:UserNode)-[:CONNECTION {status:$status}]->(newFriend:UserNode) WHERE user <> newFriend AND NOT (newFriend)<-[:CONNECTION {status:$status}]->(user) RETURN newFriend.uid, newFriend.status, newFriend.username, newFriend.firstName, newFriend.lastName,  count(newFriend) as frequency ORDER BY frequency DESC LIMIT 20", map[string]interface{}{
			"userUid": userId,
			"status":  "CONNECTED",
		})

		for records.Next() {
			node := model.User{UserUID: records.Record().Values[0].(string), Status: model.ProfileStatus(records.Record().Values[1].(string)), Username: records.Record().Values[2].(string), FirstName: records.Record().Values[3].(string), LastName: records.Record().Values[4].(string)}
			userNodes = append(userNodes, &node)
		}
		if userNodes == nil {
			fmt.Println("checkpoint 1")
			recordsFamily, er := tx.Run("match (n:UserNode {uid:$userUid}) match (u:UserNode {lastName:n.lastName}) where n <> u return u.uid, u.status, u.username, u.firstName, u.lastName LIMIT 20", map[string]interface{}{
				"userUid": userId,
			})
			for recordsFamily.Next() {
				fmt.Println("checkpoint 2")
				node := model.User{UserUID: recordsFamily.Record().Values[0].(string), Status: model.ProfileStatus(recordsFamily.Record().Values[1].(string)), Username: recordsFamily.Record().Values[2].(string), FirstName: recordsFamily.Record().Values[3].(string), LastName: recordsFamily.Record().Values[4].(string)}
				userNodes = append(userNodes, &node)
			}
			if er != nil {
				fmt.Println("checkpoint 3")
				return nil, er
			}
			fmt.Println("checkpoint 4")
			return userNodes, nil
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

func (r ConnectionRepositoryImpl) GetRecommendedJobOffers(userId string, ctx context.Context) (jobNodes []*model.JobOffer, error1 error) {
	span := tracer.StartSpanFromContext(ctx, "GetRecommendedJobOffersRepository")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	session := (*r.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			r.logError.Logger.Errorf(neo4jSessionError)
		}
	}(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		if !checkIfUserExist(userId, tx, ctx) {
			return nil, nil
		}

		records, err := tx.Run("MATCH (user:UserNode {uid:$userUid})-[:HAS]->(e:SkillNode)<-[:REQUIRES]->(job:JobNode)  RETURN job.jobId, job.publisher, job.position, job.jobDescription, job.datePosted, job.duration, job.requirements, count(job) as frequency ORDER BY frequency DESC LIMIT 20", map[string]interface{}{
			"userUid": userId,
		})

		for records.Next() {
			id, err := primitive.ObjectIDFromHex(records.Record().Values[0].(string))
			if err != nil {
				return nil, nil
			}
			var listt []string
			str := fmt.Sprintf("%v", records.Record().Values[6])
			listt = strings.Fields(str)

			firstEl := listt[0]
			listt[0] = firstEl[1:]
			lastEl := listt[len(listt)-1]
			listt[len(listt)-1] = lastEl[:len(lastEl)-1]

			node := model.JobOffer{JobId: id, Publisher: records.Record().Values[1].(string),
				Position: records.Record().Values[2].(string), JobDescription: records.Record().Values[3].(string),
				DatePosted: records.Record().Values[4].(string), Duration: records.Record().Values[5].(string),
				Requirements: listt}
			jobNodes = append(jobNodes, &node)
		}
		if jobNodes == nil {
			return nil, nil
		}

		if err != nil {
			return nil, err
		}
		return jobNodes, nil
	})
	if err != nil {
		return nil, err
	}
	return jobNodes, nil
}
