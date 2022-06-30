package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	Id          primitive.ObjectID `bson:"_id"`
	SenderId    string             `bson:"sender_one_id"`
	ReceiverId  string             `bson:"receiver_two_id"`
	MessageText string             `bson:"message_text"`
	TimeSent    time.Time          `bson:"date_posted"`
}
