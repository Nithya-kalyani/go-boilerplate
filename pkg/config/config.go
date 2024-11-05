package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	Port          string
	RedisHost     string
	RedisPort     string
	RedisPassword string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	// Map env
	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		Port:          os.Getenv("PORT"),
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}
}
