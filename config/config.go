package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func GetRabbitDSN() string {
	dsn := os.Getenv("RABBITMQ_DSN")
	if dsn == "" {
		log.Fatal("RABBITMQ_DSN is not set")
	}
	return dsn
}
