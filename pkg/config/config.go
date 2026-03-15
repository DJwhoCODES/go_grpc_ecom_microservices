package config

import "os"

type Config struct {
	ServiceName string
	ServicePort string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string

	RedisHost string
	RedisPort string
}

func Load() *Config {
	return &Config{
		ServiceName: getEnv("SERVICE_NAME", "unknown-service"),
		ServicePort: getEnv("SERVICE_PORT", "50051"),

		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:     getEnv("POSTGRES_USER", "postgres"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "postgres"),
		PostgresDB:       getEnv("POSTGRES_DB", "ecommerce"),

		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
	}
}

func getEnv(key string, fallback string) string {

	value, exists := os.LookupEnv(key)

	if !exists {
		return fallback
	}

	return value
}
