package main

import (
	"department-api/internal/config"
	"department-api/internal/database"
)

func main() {
	config := config.LoadConfig()
	database.Connect(config)
}
