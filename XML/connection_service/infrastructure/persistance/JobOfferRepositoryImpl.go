package persistance

import (
	"common/module/logger"
	connectionModel "connection/module/domain/model"
	"connection/module/domain/repositories"
	"errors"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type JobOfferRepositoryImpl struct {
	db             *neo4j.Driver
	logInfo        *logger.Logger
	logError       *logger.Logger
	connectionRepo repositories.ConnectionRepository
}

func NewJobOfferRepositoryImpl(client *neo4j.Driver, logInfo *logger.Logger, logError *logger.Logger, connectionRepo repositories.ConnectionRepository) repositories.JobOfferRepository {
	return &UserRepositoryImpl{
		db:             client,
		logInfo:        logInfo,
		logError:       logError,
		connectionRepo: connectionRepo,
	}
}
func (u UserRepositoryImpl) Create(job *connectionModel.JobOffer) (connectionModel.JobOffer, error) {
	fmt.Println("[Create job offer connection service]")
	fmt.Println(job)
	session := (*u.db).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)
	fmt.Println(session)
	result, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("CREATE (n:JobNode { jobId:$jobId, publisher:$publisher, position:$position, jobDescription:$jobDescription, datePosted: $datePosted, duration:$duration, requirements:$requirements  }) RETURN n.jobId", map[string]interface{}{
			"jobId":          job.JobId.Hex(),
			"publisher":      job.Publisher,
			"position":       job.Position,
			"jobDescription": job.JobDescription,
			"datePosted":     job.DatePosted,
			"duration":       job.Duration,
			"requirements":   job.Requirements,
		})
		fmt.Println("kreiran offer")
		if err != nil {
			fmt.Println("error nakon kreiranja noda")
			return nil, err
		}
		_, singleErr := records.Single()
		if singleErr != nil {
			return nil, singleErr
		}
		for _, e := range job.Requirements {
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
				fmt.Println("ako skill vec postoji povezem ga na job ")
				_, errCon := tx.Run("MATCH (u:JobNode {jobId:$jobId}) MATCH (e:SkillNode {name:$name}) CREATE (u)-[r:REQUIRES]->(e) RETURN r",
					map[string]interface{}{
						"jobId": job.JobId.Hex(),
						"name":  e,
					})
				if errCon != nil {
					return nil, errCon
				}

			} else {
				fmt.Println("ako ne postoji node kreiraj node i vezu ka tom nodu ")
				_, errCreate := tx.Run("MATCH (u:JobNode {jobId:$jobId}) CREATE (e:SkillNode {name:$name}) CREATE (u)-[r:REQUIRES]->(e) RETURN r",
					map[string]interface{}{
						"jobId": job.JobId.Hex(),
						"name":  e,
					})
				if errCreate != nil {
					return nil, errCreate
				}
			}

		}
		return connectionModel.JobOffer{
			JobId: job.JobId,
		}, nil
	})

	if err != nil {
		return connectionModel.JobOffer{}, err
	}
	if result == nil {
		return connectionModel.JobOffer{}, errors.New("result empty")
	}

	return connectionModel.JobOffer{
		JobId: job.JobId,
	}, nil
}
