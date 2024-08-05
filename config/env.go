package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TursoDatabaseUrl string
	TursoAuthToken   string
	Port             string
}

var Envs = InitConfig()

func InitConfig() Config {
	godotenv.Load()

	return Config{
		TursoDatabaseUrl: getEnv("TURSO_DATABASE_URL", ""),
		TursoAuthToken:   getEnv("TURSO_AUTH_TOKEN", ""),
		Port:             getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
