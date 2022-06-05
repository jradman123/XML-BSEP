package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type JobOffer struct {
	Id             primitive.ObjectID `bson:"_id"`
	Publisher      string             `bson:"publisher"`
	Position       string             `bson:"position"`
	JobDescription string             `bson:"description"`
	Requirements   []string           `bson:"requirements"`
	DatePosted     time.Time          `bson:"date_posted"`
	Duration       time.Time          `bson:"duration"`
}
