package services

import (
	"common/module/logger"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
	"fmt"
)

type JobOfferService struct {
	jobRepo  repositories.JobOfferRepository
	logInfo  *logger.Logger
	logError *logger.Logger
}

func NewJobOfferService(jobRepo repositories.JobOfferRepository, logInfo *logger.Logger, logError *logger.Logger) *JobOfferService {
	return &JobOfferService{jobRepo, logInfo, logError}
}

func (s JobOfferService) DeleteJob(job model.JobOffer) error {
	return nil
}

func (s JobOfferService) CreateJob(job model.JobOffer) error {
	fmt.Println("create job at job service, connection_service")
	_, err := s.jobRepo.Create(&job)
	if err != nil {
		return err
	}
	return nil
}
