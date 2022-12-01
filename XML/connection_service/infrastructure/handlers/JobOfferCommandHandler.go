package handlers

import (
	events "common/module/saga/job_events"
	saga "common/module/saga/messaging"
	"connection/module/application/services"
	"connection/module/domain/model"
	"context"
	"fmt"
)

type JobOfferCommandHandler struct {
	service           *services.JobOfferService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewJobOfferCommandHandler(service *services.JobOfferService, publisher saga.Publisher, subscriber saga.Subscriber) (*JobOfferCommandHandler, error) {
	o := &JobOfferCommandHandler{
		service:           service,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *JobOfferCommandHandler) handle(command *events.JobOfferCommand) {

	fmt.Println("usao u user command handler connection servisa")

	job := model.JobOffer{
		Requirements:   command.JobOffer.Requirements,
		Position:       command.JobOffer.Position,
		DatePosted:     command.JobOffer.DatePosted.String(),
		Publisher:      command.JobOffer.Publisher,
		JobDescription: command.JobOffer.JobDescription,
		Duration:       command.JobOffer.Duration.String(),
		JobId:          command.JobOffer.JobId,
	}

	var reply = events.JobOfferReply{}
	switch command.Type {
	case events.CreateJobOffer:
		fmt.Println("events.DeleteJobOffer")
		err := handler.service.CreateJob(job, context.TODO())
		if err != nil {
			reply.Type = events.JobRolledBack
		}
		reply.Type = events.JobOfferCreated

		// TODO:Cannot update users' username
	case events.DeleteJobOffer:
		fmt.Println("events.DeleteJobOffer")
		err := handler.service.DeleteJob(job, context.TODO())
		if err != nil {
			reply.Type = events.JobRolledBack
		}
		reply.Type = events.JobOfferDeleted
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}

	fmt.Println("dosao do kraja user command handler-a connection servisa")

}
