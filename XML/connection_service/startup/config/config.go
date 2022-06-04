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
		Port:          "8083",
		Neo4jUsername: "neo4j",
		Neo4jPassword: "90210",
		Neo4jUri:      "bolt://localhost:7687",
		PublicKey:     "-----BEGIN PUBLIC KEY-----\\n\" +\n\t\t\t\"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0AzWYJTc9jiPn+RMNjMJ\\n\" +\n\t\t\t\"hscn8hg/Mt0U22efM6IvM83CyQCiFHP1Z8rs2HFqRbid/hQxW23HrXQzKx5hGPdU\\n\" +\n\t\t\t\"14ncF8oN7utDQxdq6ivTsF1tMQtHWb2jnYmpKwTyelbMMGKLHj3yy2j59Y/X94EX\\n\" +\n\t\t\t\"PNtQtgAO9FF5gKzjkaBu6KzLU2RJC9bADVd5sotM/JP/Ce5D/97XV7i1KStTUDiV\\n\" +\n\t\t\t\"fDBWCkDylBTQTmI1rO9MdayVduuAzNdWXRfyqKcWI2i4pA1aaskiaViVsIhF3ksm\\n\" +\n\t\t\t\"YW4Bu0RxK5SP2byHj7pv93XsabA+QXZ37QRhYzBxx6nS0x/dNtAxIltIBZaeSTN0\\n\" +\n\t\t\t\"gQIDAQAB\\n\" +\n\t\t\t\"-----END PUBLIC KEY-----",
	}
}
