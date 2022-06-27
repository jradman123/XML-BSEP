package config

import "os"

type Config struct {
	Port            string
	UserHost        string
	UserPort        string
	PostsHost       string
	PostsPort       string
	UserDBHost      string
	UserDBPort      string
	UserDBName      string
	UserDBUser      string
	UserDBPass      string
	ConnectionsHost string
	ConnectionsPort string
}

func NewConfig() *Config {
	return &Config{
		Port:            os.Getenv("GATEWAY_PORT"),
		UserHost:        os.Getenv("USER_SERVICE_HOST"),
		UserPort:        os.Getenv("USER_SERVICE_PORT"),
		PostsHost:       os.Getenv("POST_SERVICE_HOST"),
		PostsPort:       os.Getenv("POST_SERVICE_PORT"),
		UserDBHost:      os.Getenv("USER_DB_HOST"),
		UserDBPort:      os.Getenv("USER_DB_PORT"),
		UserDBName:      os.Getenv("USER_DB_NAME"),
		UserDBUser:      os.Getenv("USER_DB_USER"),
		UserDBPass:      os.Getenv("USER_DB_PASS"),
		ConnectionsPort: "8084",
		ConnectionsHost: "connection_service",
	}
}
