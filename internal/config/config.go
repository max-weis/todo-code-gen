package config

import (
	"github.com/labstack/gommon/log"
	"os"
)

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func NewDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		User:     readEnvVar("MYSQL_USER", "mysql"),
		Password: readEnvVar("MYSQL_PASSWORD", "mysql"),
		Host:     readEnvVar("MYSQL_HOST", "localhost"),
		Port:     readEnvVar("MYSQL_PORT", "3306"),
		DbName:   readEnvVar("MYSQL_DB", "todo"),
	}
}

func readEnvVar(key, def string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Printf("Could not read env var: %s, using default value: %s", key, def)
		return def
	}

	return value
}
