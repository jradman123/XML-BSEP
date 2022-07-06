package orchestrators

import (
	events "common/module/saga/job_events"
	saga "common/module/saga/messaging"
	"fmt"
	"post/module/domain/model"
)

type JobOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewJobOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*JobOrchestrator, error) {
	o := &JobOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle) //slusa odgovore
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *JobOrchestrator) handle(reply *events.JobOfferReply) events.JobOfferReplyType {
	if reply.Type == events.JobRolledBack {
		fmt.Println("UserRolledBack")
	}
	if reply.Type == events.JobOfferCreated {
		fmt.Println("UserCreated")
	}
	if reply.Type == events.JobOfferDeleted {
		fmt.Println("UserUpdated")
	}
	fmt.Println("BAAAAAAACKS")
	return events.UnknownReply

}

func (o *JobOrchestrator) CreateJobOffer(jobOffer model.JobOffer) error {
	event := &events.JobOfferCommand{
		Type: events.CreateJobOffer,
		JobOffer: events.JobOffer{
			JobId:          jobOffer.Id,
			Duration:       jobOffer.Duration,
			JobDescription: jobOffer.JobDescription,
			DatePosted:     jobOffer.DatePosted,
			Position:       jobOffer.Position,
			Publisher:      jobOffer.Publisher,
			Requirements:   jobOffer.Requirements,
		},
	}
	return o.commandPublisher.Publish(event)
}
