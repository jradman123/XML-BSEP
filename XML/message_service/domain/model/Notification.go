package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type NotificationType int

const (
	PROFILE = iota
	POST
	MESSAGE
)

type Notification struct {
	Id               primitive.ObjectID `bson:"_id"`
	Timestamp        time.Time          `bson:"timestamp"`
	Content          string             `bson:"content"`
	RedirectPath     string             `bson:"redirect_path"`
	Read             bool               `bson:"read"`
	Type             NotificationType   `bson:"type"`
	NotificationFrom string             `bson:"notification_from"` //username
	NotificationTo   string             `bson:"notification_to"`
}
