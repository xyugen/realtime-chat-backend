package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TursoDatabaseUrl string
	TursoAuthToken   string
}

var Envs = InitConfig()

func InitConfig() Config {
	godotenv.Load()

	return Config{
		TursoDatabaseUrl: getEnv("TURSO_DATABASE_URL", ""),
		TursoAuthToken:   getEnv("TURSO_AUTH_TOKEN", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
