package main

import (
	"department-api/internal/config"
	"department-api/internal/database"
	"department-api/internal/models"
)

func main() {
	config := config.LoadConfig()
	database := database.Connect(config)
}
