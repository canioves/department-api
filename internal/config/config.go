package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	AppPort    string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Config error:", err)
	}

	config := Config{}
	envFlag := os.Getenv("GO_ENV")

	if envFlag == "docker" {
		config.DBHost = os.Getenv("DB_HOST_DOCKER")
	} else {
		config.DBHost = os.Getenv("DB_HOST_LOCAL")
	}

	config.DBName = os.Getenv("DB_NAME")
	config.DBUser = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBPort = os.Getenv("DB_PORT")
	config.AppPort = os.Getenv("APP_PORT")

	return &config
}
