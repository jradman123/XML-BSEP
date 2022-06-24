package config

type Config struct {
	Port          string
	Neo4jUri      string
	Neo4jUsername string
	Neo4jPassword string
	PublicKey     string
}

func NewConfig() *Config {
	return &Config{
		Port:          "8084",
		Neo4jUsername: "neo4j",
		Neo4jPassword: "90210",
		Neo4jUri:      "bolt://localhost:7687",
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
