package persistance

import (
	"common/module/logger"
	connectionModel "connection/module/domain/model"
	"connection/module/domain/repositories"
	"context"
	"errors"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	tracer "monitoring/module"
	"time"
)

type UserRepositoryImpl struct {
	db             *neo4j.Driver
	logInfo        *logger.Logger
	logError       *logger.Logger
	connectionRepo repositories.ConnectionRepository
}

func NewUserRepositoryImpl(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger, connectionRepo repositories.ConnectionRepository) repositories.UserRepository {
	return &UserRepositoryImpl{
		db:             client,
		logInfo:        logInfo,
		logError:       logError,
		connectionRepo: connectionRepo,
	}
}

func (u UserRepositoryImpl) Register(userNode *connectionModel.User, ctx context.Context) (*connectionModel.User, error) {
	span := tracer.StartSpanFromContext(ctx, "registerUser")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
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
		if checkIfUserExist(userNode.UserUID, tx, ctx) {
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
func checkIfUserExist(uid string, transaction neo4j.Transaction, ctx context.Context) bool {
	span := tracer.StartSpanFromContext(ctx, "checkIfUserExist")
	defer span.Finish()
	result, _ := transaction.Run(
		"MATCH (n:UserNode { uid: $uid }) RETURN n.uid",
		map[string]interface{}{"uid": uid})

	if result != nil && result.Next() && result.Record().Values[0] == uid {
		return true
	}
	return false
}
func (u UserRepositoryImpl) UpdateUser(userNode *connectionModel.User, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "updateUserRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		fmt.Println("UUID " + userNode.UserUID)
		if checkIfUserExist(userNode.UserUID, tx, ctx) {
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
			fmt.Println("NE POSTOJI")
			return nil, nil
		}

	})
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepositoryImpl) GetUserId(username string, ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "getUserIdRepository")
	defer span.Finish()
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

func (u UserRepositoryImpl) ChangeProfileStatus(m *connectionModel.User, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "changeProfileStatusRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		fmt.Println("UUID " + m.UserUID)
		if checkIfUserExist(m.UserUID, tx, ctx) {

			var status string
			if m.Status == connectionModel.Private {
				status = "PRIVATE"
			} else if m.Status == connectionModel.Public {
				status = "PUBLIC"
			}
			fmt.Println(status)

			results, err := tx.Run("MATCH (n:UserNode { uid: $uid}) set n.status = $status return n.status",
				map[string]interface{}{
					"uid":    m.UserUID,
					"status": m.Status,
				})
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			if results.Next() {
				if results.Record().Values[0].(string) == "PUBLIC" {
					dateNow := time.Now().Local().Unix()
					_, goldErr := tx.Run("MATCH (u1:UserNode)   MATCH (u2:UserNode) WHERE u1.uid =$oldPrivateUser match (u1)-[r2:CONNECTION {status:$oldStatus}]->(u2) SET r2.status=$newStatus CREATE (u2)-[r1:CONNECTION {status:$newStatus, date: $date}]->(u1)",
						map[string]interface{}{
							"oldPrivateUser": m.UserUID,
							"oldStatus":      "REQUEST_SENT",
							"newStatus":      "CONNECTED",
							"date":           dateNow,
						})
					if goldErr != nil {
						return nil, goldErr
					}

				}
			}
		} else {
			fmt.Println("NE POSTOJI")
			return nil, nil
		}

		return true, nil
	})
	fmt.Println(result)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepositoryImpl) UpdateUserProfessionalDetails(user *connectionModel.User, details *connectionModel.UserDetails, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "updateUserProfessionalDetailsRepository")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		fmt.Println("UUID " + user.UserUID)
		if !checkIfUserExist(user.UserUID, tx, ctx) {
			fmt.Println("NE POSTOJI")
			return nil, nil
		}
		return user.UserUID, nil
	})
	if result == nil {
		return errors.New("user doesn't exist")
	}
	err = u.updateSkills(user.UserUID, details.Skills, ctx)
	if err != nil {
		return err
	}
	err = u.updateInterests(user.UserUID, details.Interests, ctx)
	if err != nil {
		return err
	}
	err = u.updateEducations(user.UserUID, details.Educations, ctx)
	if err != nil {
		return err
	}
	err = u.updateExperiences(user.UserUID, details.Experiences, ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepositoryImpl) updateExperiences(uid string, experiences []string, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "updateExperiences")
	defer span.Finish()
	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		// TODO: UKLONITI VEZE KOJE VISE NE VAZE
		// brisanje
		//oldExperiences, oldErr := tx.Run("MATCH (u:UserNode {uid:$uid}) MATCH (u)-[r:HAS]->(e:ExperienceNode) RETURN e.name",
		//	map[string]interface{}{
		//		"uid": uid,
		//	})
		//if oldErr != nil {
		//	fmt.Println(oldErr)
		//	return nil, oldErr
		//}
		//var oldExperiencesStrings []string
		//for oldExperiences.Next() {
		//	oldExperiencesStrings := append(oldExperiencesStrings, oldExperiences.Record().Values[0].(string))
		//	fmt.Println(oldExperiencesStrings)
		//}
		//deleted := compareExperiences(oldExperiencesStrings, experiences)
		//for _, d := range deleted {
		//	_, err := tx.Run("MATCH (u:UserNode {uid:$uid} MATCH (n:ExperienceNode { name: $name}) MATCH (u)-[r:HAS]->(e) DELETE r",
		//		map[string]interface{}{
		//			"name": d,
		//			"uid":  uid,
		//		})
		//	if err != nil {
		//		fmt.Println(err)
		//		return nil, err
		//	}
		//}
		// brisanje
		for _, e := range experiences {
			fmt.Println("za svaaki skill gledam dal postoji ")
			results, err := tx.Run("MATCH (n:SkillNode { name: $name}) return n.name",
				map[string]interface{}{
					"name": e,
				})
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			if results.Next() {
				fmt.Println("ako skill postoji gledal dal vec postoji veza od usera ka tom skilu")
				connectionExists, errCon := tx.Run("MATCH (u:UserNode {uid:$uid}) MATCH (e:SkillNode {name:$name}) MATCH (u)-[r:HAS]->(e) RETURN r",
					map[string]interface{}{
						"uid":  uid,
						"name": e,
					})
				if errCon != nil {
					return nil, errCon
				}
				if !connectionExists.Next() {
					fmt.Println("ako posoji node i ne postoji konekcija kreiraj konekciju od usera ka nodu")
					_, errCreate := tx.Run("MATCH (u:UserNode {uid:$uid}) MATCH (e:SkillNode {name:$name}) CREATE (u)-[r:HAS]->(e) RETURN r",
						map[string]interface{}{
							"uid":  uid,
							"name": e,
						})
					if errCreate != nil {
						return nil, errCreate
					}
				}
				fmt.Println("ako postoji konekcija ka tom nodu nista ne radi")

			} else {
				fmt.Println("ako ne postoji node kreiraj node i vezu ka tom nodu ")
				_, errCreate := tx.Run("MATCH (u:UserNode {uid:$uid}) CREATE (e:SkillNode {name:$name}) CREATE (u)-[r:HAS]->(e) RETURN r",
					map[string]interface{}{
						"uid":  uid,
						"name": e,
					})
				if errCreate != nil {
					return nil, errCreate
				}
			}

		}
		fmt.Println("kraj")
		return nil, nil
	})
	return err
}

func compareExperiences(oldExperiences []string, experiences []string) []string {
	var deleted []string
	// TODO : ODRADITI LOGIKU OVOG
	return deleted
}

func (u UserRepositoryImpl) updateEducations(uid string, educations []string, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "updateEducations")
	defer span.Finish()
	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		// TODO: UKLONITI VEZE KOJE VISE NE VAZE
		for _, e := range educations {
			fmt.Println("za svaaki skill gledam dal postoji ")
			results, err := tx.Run("MATCH (n:SkillNode { name: $name}) return n.name",
				map[string]interface{}{
					"name": e,
				})
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			if results.Next() {
				fmt.Println("ako skill postoji gledal dal vec postoji veza od usera ka tom skilu")
				connectionExists, errCon := tx.Run("MATCH (u:UserNode {uid:$uid}) MATCH (e:SkillNode {name:$name}) MATCH (u)-[r:HAS]->(e) RETURN r",
					map[string]interface{}{
						"uid":  uid,
						"name": e,
					})
				if errCon != nil {
					return nil, errCon
				}
				if !connectionExists.Next() {
					fmt.Println("ako posoji node i ne postoji konekcija kreiraj konekciju od usera ka nodu")
					_, errCreate := tx.Run("MATCH (u:UserNode {uid:$uid}) MATCH (e:SkillNode {name:$name}) CREATE (u)-[r:HAS]->(e) RETURN r",
						map[string]interface{}{
							"uid":  uid,
							"name": e,
						})
					if errCreate != nil {
						return nil, errCreate
					}
				}
				fmt.Println("ako postoji konekcija ka tom nodu nista ne radi")

			} else {
				fmt.Println("ako ne postoji node kreiraj node i vezu ka tom nodu ")
				_, errCreate := tx.Run("MATCH (u:UserNode {uid:$uid}) CREATE (e:SkillNode {name:$name}) CREATE (u)-[r:HAS]->(e) RETURN r",
					map[string]interface{}{
						"uid":  uid,
						"name": e,
					})
				if errCreate != nil {
					return nil, errCreate
				}
			}

		}
		fmt.Println("kraj")
		return nil, nil
	})
	return err

}

func (u UserRepositoryImpl) updateInterests(uid string, interests []string, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "updateInterests")
	defer span.Finish()
	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		// TODO: UKLONITI VEZE KOJE VISE NE VAZE

		for _, e := range interests {
			fmt.Println("za svaaki skill gledam dal postoji ")
			results, err := tx.Run("MATCH (n:SkillNode { name: $name}) return n.name",
				map[string]interface{}{
					"name": e,
				})
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			if results.Next() {
				fmt.Println("ako skill postoji gledal dal vec postoji veza od usera ka tom skilu")
				connectionExists, errCon := tx.Run("MATCH (u:UserNode {uid:$uid}) MATCH (e:SkillNode {name:$name}) MATCH (u)-[r:HAS]->(e) RETURN r",
					map[string]interface{}{
						"uid":  uid,
						"name": e,
					})
				if errCon != nil {
					return nil, errCon
				}
				if !connectionExists.Next() {
					fmt.Println("ako posoji node i ne postoji konekcija kreiraj konekciju od usera ka nodu")
					_, errCreate := tx.Run("MATCH (u:UserNode {uid:$uid}) MATCH (e:SkillNode {name:$name}) CREATE (u)-[r:HAS]->(e) RETURN r",
						map[string]interface{}{
							"uid":  uid,
							"name": e,
						})
					if errCreate != nil {
						return nil, errCreate
					}
				}
				fmt.Println("ako postoji konekcija ka tom nodu nista ne radi")

			} else {
				fmt.Println("ako ne postoji node kreiraj node i vezu ka tom nodu ")
				_, errCreate := tx.Run("MATCH (u:UserNode {uid:$uid}) CREATE (e:SkillNode {name:$name}) CREATE (u)-[r:HAS]->(e) RETURN r",
					map[string]interface{}{
						"uid":  uid,
						"name": e,
					})
				if errCreate != nil {
					return nil, errCreate
				}
			}

		}
		fmt.Println("kraj")
		return nil, nil
	})
	return err
}

func (u UserRepositoryImpl) updateSkills(uid string, skills []string, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "updateSkills")
	defer span.Finish()

	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		// TODO: UKLONITI VEZE KOJE VISE NE VAZE
		for _, e := range skills {
			fmt.Println("za svaaki skill gledam dal postoji ")
			results, err := tx.Run("MATCH (n:SkillNode { name: $name}) return n.name",
				map[string]interface{}{
					"name": e,
				})
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			if results.Next() {
				fmt.Println("ako skill postoji gledal dal vec postoji veza od usera ka tom skilu")
				connectionExists, errCon := tx.Run("MATCH (u:UserNode {uid:$uid}) MATCH (e:SkillNode {name:$name}) MATCH (u)-[r:HAS]->(e) RETURN r",
					map[string]interface{}{
						"uid":  uid,
						"name": e,
					})
				if errCon != nil {
					return nil, errCon
				}
				if !connectionExists.Next() {
					fmt.Println("ako posoji node i ne postoji konekcija kreiraj konekciju od usera ka nodu")
					_, errCreate := tx.Run("MATCH (u:UserNode {uid:$uid}) MATCH (e:SkillNode {name:$name}) CREATE (u)-[r:HAS]->(e) RETURN r",
						map[string]interface{}{
							"uid":  uid,
							"name": e,
						})
					if errCreate != nil {
						return nil, errCreate
					}
				}
				fmt.Println("ako postoji konekcija ka tom nodu nista ne radi")

			} else {
				fmt.Println("ako ne postoji node kreiraj node i vezu ka tom nodu ")
				_, errCreate := tx.Run("MATCH (u:UserNode {uid:$uid}) CREATE (e:SkillNode {name:$name}) CREATE (u)-[r:HAS]->(e) RETURN r",
					map[string]interface{}{
						"uid":  uid,
						"name": e,
					})
				if errCreate != nil {
					return nil, errCreate
				}
			}

		}
		fmt.Println("kraj")
		return nil, nil
	})
	return err
}
