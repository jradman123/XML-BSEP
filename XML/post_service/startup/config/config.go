package config

type Config struct {
	Port       string
	PostDBHost string
	PostDBPort string
	PublicKey  string
}

func NewConfig() *Config {

	return &Config{
		Port:       "8083",
		PostDBHost: "localhost",
		PostDBPort: "27017",
		PublicKey:  "PostService",
	}
}
