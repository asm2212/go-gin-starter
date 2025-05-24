package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort  string
	AppEnv   string
	DBHost   string
	DBPort   int
	DBUser   string
	DBPass   string
	DBName   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading from environment")
	}

	dbPort, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))

	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "development"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  dbPort,
		DBUser:  getEnv("DB_USER", "postgres"),
		DBPass:  getEnv("DB_PASSWORD", "postgres"),
		DBName:  getEnv("DB_NAME", "starterdb"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}