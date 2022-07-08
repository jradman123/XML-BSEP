package job_events

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type JobOffer struct {
	JobId          primitive.ObjectID
	Publisher      string
	Position       string
	JobDescription string
	Requirements   []string
	DatePosted     time.Time
	Duration       time.Time
}

type JobCommandType int8

const (
	CreateJobOffer JobCommandType = iota
	DeleteJobOffer
	RollbackJobOffer
	UnknownCommand
)

type JobOfferReplyType int8

const (
	JobOfferCreated JobOfferReplyType = iota
	JobOfferDeleted
	JobRolledBack
	UnknownReply
)

type JobOfferCommand struct {
	JobOffer JobOffer
	Type     JobCommandType
}

type JobOfferReply struct {
	Type JobOfferReplyType
}
