package config

import "os"

type Config struct {
	Port               string
	PostDBHost         string
	PostDBPort         string
	PublicKey          string
	UserCommandSubject string
	UserReplySubject   string
	NatsHost           string
	NatsUser           string
	NatsPort           string
	NatsPass           string
	JobCommandSubject  string
	JobReplySubject    string
}

func NewConfig() *Config {

	return &Config{
		Port:               os.Getenv("POST_SERVICE_PORT"),
		PostDBHost:         os.Getenv("POST_DB_HOST"),
		PostDBPort:         os.Getenv("POST_DB_PORT"),
		PublicKey:          "PostService",
		UserCommandSubject: os.Getenv("USER_COMMAND_SUBJECT"),
		UserReplySubject:   os.Getenv("USER_REPLY_SUBJECT"),
		NatsPort:           os.Getenv("NATS_PORT"),
		NatsHost:           os.Getenv("NATS_HOST"),
		NatsPass:           os.Getenv("NATS_PASS"),
		NatsUser:           os.Getenv("NATS_USER"),
		JobCommandSubject:  os.Getenv("JOB_COMMAND_SUBJECT"),
		JobReplySubject:    os.Getenv("JOB_REPLY_SUBJECT"),
	}
}
