package persistance

import (
	"common/module/logger"
	connectionModel "connection/module/domain/model"
	"connection/module/domain/repositories"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UserRepositoryImpl struct {
	db       *neo4j.Driver
	logInfo  *logger.Logger
	logError *logger.Logger
}

func NewUserRepositoryImpl(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger) repositories.UserRepository {
	return &UserRepositoryImpl{
		db:       client,
		logInfo:  logInfo,
		logError: logError,
	}
}

func (u UserRepositoryImpl) Register(userNode *connectionModel.User) (*connectionModel.User, error) {
	fmt.Println("[ConnectionDBStore Register]")
	fmt.Println(userNode)
	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	fmt.Println(session)
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		fmt.Println("linija5")
		if checkIfUserExist(userNode.UserUID, tx) {
			fmt.Println("linija36")
			return &connectionModel.User{
				UserUID:   "",
				Status:    "",
				FirstName: "",
				LastName:  "",
				Username:  "",
			}, nil
		}

		fmt.Println("[ConnectionDBStore Register1]")
		fmt.Println(userNode)
		records, err := tx.Run("CREATE (n:UserNode { uid: $uid, status: $status, username: $username, firstName: $firstName, lastName: $lastName  }) RETURN n.uid, n.status", map[string]interface{}{
			"uid":       userNode.UserUID,
			"status":    userNode.Status,
			"username":  userNode.Username,
			"firstName": userNode.FirstName,
			"lastName":  userNode.LastName,
		})
		fmt.Println("TU SAM")
		if err != nil {
			return nil, err
		}
		record, err := records.Single()
		if err != nil {
			return nil, err
		}
		// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
		return &connectionModel.User{
			UserUID: record.Values[0].(string),
			Status:  connectionModel.ProfileStatus(record.Values[1].(string)),
		}, nil
	})

	if err != nil {
		return nil, err
	}

	return result.(*connectionModel.User), nil
}
func checkIfUserExist(uid string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run(
		"MATCH (n:UserNode { uid: $uid }) RETURN n.uid",
		map[string]interface{}{"uid": uid})

	if result != nil && result.Next() && result.Record().Values[0] == uid {
		return true
	}
	return false
}
func (u UserRepositoryImpl) UpdateUser(userNode *connectionModel.User) error {
	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		fmt.Println("UUID " + userNode.UserUID)
		if checkIfUserExist(userNode.UserUID, tx) {

			//var status string
			//if userNode.Status == connectionModel.Private {
			//	status = "PRIVATE"
			//} else {
			//	status = "PUBLIC"
			//}
			//fmt.Println("MENJAM U " + status)

			_, err := tx.Run("MATCH (n:UserNode { uid: $uid}) set n.status = $status, n.username = $username, n.firstName = $firstName, n.lastName = $lastName",
				map[string]interface{}{
					"uid":       userNode.UserUID,
					"status":    userNode.Status,
					"username":  userNode.Username,
					"firstName": userNode.FirstName,
					"lastName":  userNode.LastName,
				})

			if err != nil {
				return nil, err
			}
			return nil, nil
		} else {
			fmt.Println("NEPOSTOJI")
			return nil, nil
		}

	})
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepositoryImpl) GetUserId(username string) (string, error) {

	fmt.Println("[ConnectionDBStore GetUserId]")
	fmt.Println(username)
	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	fmt.Println(session)
	result, err := session.Run(
		"MATCH (n:UserNode { username: $username }) RETURN n.uid",
		map[string]interface{}{"username": username})

	if err != nil {
		return "", err
	}

	if result != nil && result.Next() {
		return result.Record().Values[0].(string), nil
	}

	return "", nil
}
