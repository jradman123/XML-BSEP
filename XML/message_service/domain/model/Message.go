package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	Id          primitive.ObjectID `bson:"_id"`
	SenderId    uuid.UUID          `bson:"sender_id"`
	ReceiverId  uuid.UUID          `bson:"receiver_id"`
	MessageText string             `bson:"message_text"`
	TimeSent    time.Time          `bson:"time_sent"`
}
