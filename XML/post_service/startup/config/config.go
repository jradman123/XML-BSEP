package config

import "os"

type Config struct {
	Port       string
	PostDBHost string
	PostDBPort string
	PublicKey  string
}

func NewConfig() *Config {

	return &Config{
		Port:       os.Getenv("POST_SERVICE_PORT"),
		PostDBHost: os.Getenv("POST_DB_HOST"),
		PostDBPort: os.Getenv("POST_DB_PORT"),
		PublicKey:  "PostService",
	}
}
