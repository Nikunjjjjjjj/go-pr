package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHOST              string
	DBUSER              string
	DBPassword          string
	DBName              string
	DBPort              string
	JWTSecret           string
	Bucket              string
	SMTP_EMAIL          string
	SMTP_EMAIL_PASSWORD string
}

func LoadEnv() Config {
	godotenv.Load()
	return Config{
		DBHOST:              getenv("DB_HOST", "localhost"),
		DBUSER:              getenv("DB_USER", "postgres"),
		DBPassword:          getenv("DB_PASSWORD", "1234"),
		DBName:              getenv("DB_NAME", "gindb"),
		DBPort:              getenv("DB_PORT", "5432"),
		JWTSecret:           getenv("JWT_SECRET", "key1"),
		Bucket:              getenv("BUCKET", "bucket0710"),
		SMTP_EMAIL:          getenv("SMTP_EMAIL", "no env"),
		SMTP_EMAIL_PASSWORD: getenv("SMTP_EMAIL_PASSWORD", "nope"),
	}
}

func getenv(key, defaults string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaults
	}
	return val
}
