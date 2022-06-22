package config

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
}

func NewConfig() *Config {

	return &Config{
		Port:               "8083",
		PostDBHost:         "localhost",
		PostDBPort:         "27017",
		PublicKey:          "PostService",
		UserCommandSubject: "user.command",
		UserReplySubject:   "user.reply",
	}
}
