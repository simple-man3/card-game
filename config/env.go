package config

import (
	"github.com/joho/godotenv"
	"os"
)

var EnvInstance *env

type env struct {
	AppPort string

	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string

	JwtSecret string
}

func InitEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	EnvInstance = &env{
		AppPort:    getEnv("APP_PORT", ""),
		DbHost:     getEnv("DB_HOST", ""),
		DbPort:     getEnv("DB_PORT", ""),
		DbDatabase: getEnv("DB_DATABASE", ""),
		DbUsername: getEnv("DB_USERNAME", ""),
		DbPassword: getEnv("DB_PASSWORD", ""),
		JwtSecret:  getEnv("JWT_SECRET", ""),
	}

	return nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
