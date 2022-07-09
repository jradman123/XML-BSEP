package config

import "os"

type Config struct {
	Port                                 string
	MessageDBHost                        string
	MessageDBPort                        string
	PublicKey                            string
	UserCommandSubject                   string
	UserReplySubject                     string
	NatsHost                             string
	NatsUser                             string
	NatsPort                             string
	NatsPass                             string
	NotificationAppID                    string
	NotificationKey                      string
	NotificationSecret                   string
	NotificationCluster                  string
	NotificationSecure                   bool
	PostNotificationCommandSubject       string
	PostNotificationReplySubject         string
	ConnectionNotificationCommandSubject string
	ConnectionNotificationReplySubject   string
	MessageAppID                         string
	MessageKey                           string
	MessageSecret                        string
	MessageCluster                       string
	MessageSecure                        bool
}

func NewConfig() *Config {

	return &Config{
		Port:                                 os.Getenv("MESSAGE_SERVICE_PORT"),
		MessageDBHost:                        os.Getenv("MESSAGE_DB_HOST"),
		MessageDBPort:                        os.Getenv("MESSAGE_DB_PORT"),
		PublicKey:                            "PostService",
		UserCommandSubject:                   os.Getenv("USER_COMMAND_SUBJECT"),
		UserReplySubject:                     os.Getenv("USER_REPLY_SUBJECT"),
		NatsPort:                             os.Getenv("NATS_PORT"),
		NatsHost:                             os.Getenv("NATS_HOST"),
		NatsPass:                             os.Getenv("NATS_PASS"),
		NatsUser:                             os.Getenv("NATS_USER"),
		PostNotificationCommandSubject:       os.Getenv("POST_NOTIFICATION_COMMAND_SUBJECT"),
		PostNotificationReplySubject:         os.Getenv("POST_NOTIFICATION_REPLY_SUBJECT"),
		ConnectionNotificationCommandSubject: os.Getenv("CONNECTION_NOTIFICATION_COMMAND_SUBJECT"),
		ConnectionNotificationReplySubject:   os.Getenv("CONNECTION_NOTIFICATION_REPLY_SUBJECT"),
		NotificationAppID:                    "1435187",
		NotificationKey:                      "e92e3e6334c6de83b489",
		NotificationSecret:                   "da4e7a87d5c3f8dd99d0",
		NotificationCluster:                  "eu",
		NotificationSecure:                   true,
		MessageAppID:                         "1435191",
		MessageKey:                           "e49d7a86a937f12da028",
		MessageSecret:                        "cdd023488f9e1881f4be",
		MessageCluster:                       "eu",
		MessageSecure:                        true,
	}
}
