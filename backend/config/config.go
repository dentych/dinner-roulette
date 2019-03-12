package config

func ConfigFromEnv() *Configuration {
	return &Configuration{
		DbConfig: DatabaseConfig{
			Hostname: GetenvOrDefault("DINNERDASH_DB_HOST", "localhost"),
			Username: GetenvOrDefault("DINNERDASH_DB_USER", "postgres"),
			Password: GetenvOrDefault("DINNERDASH_DB_PASS", "password"),
			Database: GetenvOrDefault("DINNERDASH_DB_DB", "dinnerdash"),
		},
		CookieHost: "localhost",
	}
}

type Configuration struct {
	DbConfig   DatabaseConfig
	CookieHost string
}

type DatabaseConfig struct {
	Hostname string
	Username string
	Password string
	Database string
}
