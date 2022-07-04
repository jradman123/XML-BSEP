package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationSettings struct {
	Id          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username"`
	Posts       bool               `bson:"posts"`
	Messages    bool               `bson:"posts"`
	Connections bool               `bson:"posts"`
}
