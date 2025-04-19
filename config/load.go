package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	godotenv.Load()

	AppEnv = getEnvOr("APP_ENV", EnvDev)
	DatabaseURL = getEnvOr("DATABASE_URL", fmt.Sprintf("test_structure_%s.db", AppEnv))
}

func getEnvOr(key string, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return val
}
