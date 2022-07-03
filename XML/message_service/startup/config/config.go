package config

import "os"

type Config struct {
	Port                string
	MessageDBHost       string
	MessageDBPort       string
	PublicKey           string
	UserCommandSubject  string
	UserReplySubject    string
	NatsHost            string
	NatsUser            string
	NatsPort            string
	NatsPass            string
	NotificationAppID   string
	NotificationKey     string
	NotificationSecret  string
	NotificationCluster string
	NotificationSecure  bool
}

func NewConfig() *Config {

	return &Config{
		Port:                os.Getenv("POST_SERVICE_PORT"),
		MessageDBHost:       os.Getenv("MESSAGE_DB_HOST"),
		MessageDBPort:       os.Getenv("MESSAGE_DB_PORT"),
		PublicKey:           "PostService",
		UserCommandSubject:  os.Getenv("USER_COMMAND_SUBJECT"),
		UserReplySubject:    os.Getenv("USER_REPLY_SUBJECT"),
		NatsPort:            os.Getenv("NATS_PORT"),
		NatsHost:            os.Getenv("NATS_HOST"),
		NatsPass:            os.Getenv("NATS_PASS"),
		NatsUser:            os.Getenv("NATS_USER"),
		NotificationAppID:   "1432217",
		NotificationKey:     "6753c65c92c944af3a85",
		NotificationSecret:  "b49ce7b031eba94ac0a2",
		NotificationCluster: "eu",
		NotificationSecure:  true,
	}
}
