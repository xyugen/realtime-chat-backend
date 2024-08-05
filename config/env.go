package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret              string
	JWTExpirationInSeconds int64
	TursoDatabaseUrl       string
	TursoAuthToken         string
	Port                   string
}

var Envs = InitConfig()

func InitConfig() Config {
	godotenv.Load()

	return Config{
		JWTSecret:              getEnv("JWT_SECRET", "EggVV#rg%#2A#MgCP^9xE3"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		TursoDatabaseUrl:       getEnv("TURSO_DATABASE_URL", ""),
		TursoAuthToken:         getEnv("TURSO_AUTH_TOKEN", ""),
		Port:                   getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
