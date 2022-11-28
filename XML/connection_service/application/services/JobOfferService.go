package services

import (
	"common/module/logger"
	"connection/module/domain/model"
	"connection/module/domain/repositories"
	"context"
	"fmt"
	tracer "monitoring/module"
)

type JobOfferService struct {
	jobRepo  repositories.JobOfferRepository
	logInfo  *logger.Logger
	logError *logger.Logger
}

func NewJobOfferService(jobRepo repositories.JobOfferRepository, logInfo *logger.Logger, logError *logger.Logger) *JobOfferService {
	return &JobOfferService{jobRepo, logInfo, logError}
}

func (s JobOfferService) DeleteJob(job model.JobOffer, ctx context.Context) error {
	return nil
}

func (s JobOfferService) CreateJob(job model.JobOffer, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "CreateJobService")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)
	fmt.Println("create job at job service, connection_service")
	_, err := s.jobRepo.Create(&job, ctx)
	if err != nil {
		return err
	}
	return nil
}
