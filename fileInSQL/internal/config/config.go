package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHOST     string
	DBUSER     string
	DBPassword string
	DBName     string
	DBPort     string
}

func LoadEnv() Config {
	godotenv.Load()
	return Config{
		DBHOST:     getenv("DB_HOST", "localhost"),
		DBUSER:     getenv("DB_USER", "postgres"),
		DBPassword: getenv("DB_PASSWORD", "1234"),
		DBName:     getenv("DB_NAME", "gindb"),
		DBPort:     getenv("DB_PORT", "5432"),
	}
}

func getenv(key, defaults string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaults
	}
	return val
}
