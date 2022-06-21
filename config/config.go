package config

import (
	"fmt"
	"os"
)

type Config struct {
	Host string
	Port string

	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPassword string
	PostgresSSLMode string
	PostgresDB string
}


func Load() Config {
	config := Config{}


	config.Host = anyToString(getEnvOrDefault("HOST", "localhost"))
	config.Port = anyToString(getEnvOrDefault("PORT", 8080))

	config.PostgresHost = anyToString(getEnvOrDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = anyToString(getEnvOrDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = anyToString(getEnvOrDefault("POSTGRES_USER", "postgres"))
	config.PostgresPassword = anyToString(getEnvOrDefault("POSTGRES_PASSWORD", "postgres"))
	config.PostgresSSLMode = anyToString(getEnvOrDefault("POSTGRES_SSL_MODE", "disable"))
	config.PostgresDB = anyToString(getEnvOrDefault("POSTGRES_DB", "simple_bank"))

	return config
}


func anyToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func getEnvOrDefault(key string, default_value interface{}) interface{} {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}

	return default_value
}
