package config

import (
	"github.com/joho/godotenv"
	"os"
	"sync"
)

var (
	envInstance *Env
	once        sync.Once
	initError   error
)

type Env struct {
	AppPort string

	DbHost     string
	DbPort     string
	DbDatabase string
	DbUsername string
	DbPassword string

	JwtSecret string
}

func GetInstanceEnv() (*Env, error) {
	once.Do(func() {
		initError = godotenv.Load()

		envInstance = &Env{
			AppPort:    getEnv("APP_PORT", ""),
			DbHost:     getEnv("DB_HOST", ""),
			DbPort:     getEnv("DB_PORT", ""),
			DbDatabase: getEnv("DB_DATABASE", ""),
			DbUsername: getEnv("DB_USERNAME", ""),
			DbPassword: getEnv("DB_PASSWORD", ""),
			JwtSecret:  getEnv("JWT_SECRET", ""),
		}
	})

	return envInstance, initError
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
