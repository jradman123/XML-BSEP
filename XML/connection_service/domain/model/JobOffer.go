package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOffer struct {
	JobId          primitive.ObjectID
	Publisher      string
	Position       string
	JobDescription string
	Requirements   []string
	DatePosted     string
	Duration       string
}
