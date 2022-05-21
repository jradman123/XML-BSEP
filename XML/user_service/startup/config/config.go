package config

import "os"

type Config struct {
	Port       string
	UserDBHost string
	UserDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:       "8082",
		UserDBHost: os.Getenv("HOST"),
		UserDBPort: os.Getenv("PG_DBPORT"),
	}
}
