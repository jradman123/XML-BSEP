package model

type NotificationSettings struct {
	Posts       bool `bson:"posts"`
	Messages    bool `bson:"messages"`
	Connections bool `bson:"connections"`
}
