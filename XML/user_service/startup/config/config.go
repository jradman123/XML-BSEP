package config

import "os"

type Config struct {
	Port               string
	UserDBHost         string
	UserDBPort         string
	UserDBUser         string
	UserDBPass         string
	UserDBName         string
	PublicKey          string
	NatsHost           string
	NatsPort           string
	NatsUser           string
	NatsPass           string
	UserCommandSubject string
	UserReplySubject   string
}

func NewConfig() *Config {
	return &Config{

		Port:               os.Getenv("USER_SERVICE_PORT"),
		UserDBHost:         os.Getenv("USER_DB_HOST"),
		UserDBPort:         os.Getenv("USER_DB_PORT"),
		UserDBUser:         os.Getenv("USER_DB_USER"),
		UserDBPass:         os.Getenv("USER_DB_PASS"),
		UserDBName:         os.Getenv("USER_DB_NAME"),
		NatsPort:           os.Getenv("NATS_PORT"),
		NatsHost:           os.Getenv("NATS_HOST"),
		NatsPass:           os.Getenv("NATS_PASS"),
		NatsUser:           os.Getenv("NATS_USER"),
		UserCommandSubject: os.Getenv("USER_COMMAND_SUBJECT"),
		UserReplySubject:   os.Getenv("USER_REPLY_SUBJECT"),
		PublicKey: "-----BEGIN PUBLIC KEY-----\n" +
			"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0AzWYJTc9jiPn+RMNjMJ\n" +
			"hscn8hg/Mt0U22efM6IvM83CyQCiFHP1Z8rs2HFqRbid/hQxW23HrXQzKx5hGPdU\n" +
			"14ncF8oN7utDQxdq6ivTsF1tMQtHWb2jnYmpKwTyelbMMGKLHj3yy2j59Y/X94EX\n" +
			"PNtQtgAO9FF5gKzjkaBu6KzLU2RJC9bADVd5sotM/JP/Ce5D/97XV7i1KStTUDiV\n" +
			"fDBWCkDylBTQTmI1rO9MdayVduuAzNdWXRfyqKcWI2i4pA1aaskiaViVsIhF3ksm\n" +
			"YW4Bu0RxK5SP2byHj7pv93XsabA+QXZ37QRhYzBxx6nS0x/dNtAxIltIBZaeSTN0\n" +
			"gQIDAQAB\n" +
			"-----END PUBLIC KEY-----",
	}
}
