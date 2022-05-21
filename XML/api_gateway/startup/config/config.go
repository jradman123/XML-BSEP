package config

type Config struct {
	Port     string
	UserHost string
	UserPort string
}

func NewConfig() *Config {
	return &Config{
		Port:     "9090",
		UserHost: "localhost",
		UserPort: "8082",
		//TODO:
		/*
			lokalno => localhost
			preko dockera => ime kontainera
			Posto svaki kontainer ima svoje adrese,kada samo prosledis ime kontainera to je kao da si
			rekla localhost unutar tog mog kontainera
		*/
	}
}
