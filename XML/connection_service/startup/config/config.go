package config

import "os"

type Config struct {
	Port               string
	Neo4jUri           string
	Neo4jUsername      string
	Neo4jPassword      string
	UserCommandSubject string
	UserReplySubject   string
	NatsHost           string
	NatsUser           string
	NatsPort           string
	NatsPass           string
	PublicKey          string
}

func NewConfig() *Config {
	return &Config{
		Port:               os.Getenv("CONNECTION_SERVICE_PORT"),
		Neo4jUsername:      os.Getenv("NEO4J_USERNAME"),
		Neo4jPassword:      os.Getenv("NEO4J_PASS"),
		Neo4jUri:           "neo4j+s://525ffd8e.databases.neo4j.io",
		UserCommandSubject: os.Getenv("USER_COMMAND_SUBJECT"),
		UserReplySubject:   os.Getenv("USER_REPLY_SUBJECT"),
		NatsPort:           os.Getenv("NATS_PORT"),
		NatsHost:           os.Getenv("NATS_HOST"),
		NatsPass:           os.Getenv("NATS_PASS"),
		NatsUser:           os.Getenv("NATS_USER"),
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
