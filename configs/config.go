package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName       string
	DBUser       string
	DBPassword   string
	DBPort       string
	JWTSecretKey string
	Port         string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := &Config{
		DBName:       getEnv("DB_NAME", ""),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBUser:       getEnv("DB_USER", ""),
		DBPort:       getEnv("DB_PORT", ""),
		JWTSecretKey: getEnv("JWT_SECRET_KEY", ""),
		Port:         getEnv("PORT", "8080"),
	}

	return config
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
