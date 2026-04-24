package config

import (
	"os"
	"strconv"
)

type Config struct {
	Env              string
	Port             string
	JWTSecret        string
	DatabaseURL      string
	AuthRateLimitRPS int
}

func Load() Config {
	rps, err := strconv.Atoi(getEnv("RATE_LIMIT_AUTH_RPS", "5"))
	if err != nil || rps <= 0 {
		rps = 5
	}
	return Config{
		Env:              getEnv("APP_ENV", "development"),
		Port:             getEnv("APP_PORT", "8080"),
		JWTSecret:        getEnv("JWT_SECRET", "change-me"),
		DatabaseURL:      getEnv("DATABASE_URL", ""),
		AuthRateLimitRPS: rps,
	}
}

func getEnv(k, fallback string) string {
	v := os.Getenv(k)
	if v == "" {
		return fallback
	}
	return v
}
