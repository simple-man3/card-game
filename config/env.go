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
}

func NewEnv() (*Env, error) {
	once.Do(func() {
		initError = godotenv.Load()

		envInstance = &Env{
			AppPort: getEnv("APP_PORT", ""),
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
