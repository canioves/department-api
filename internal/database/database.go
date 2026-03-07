package database

import (
	"department-api/internal/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(config *config.Config) *gorm.DB {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName,
	)

	pg := postgres.Open(connStr)
	gormCfg := &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}

	db, err := gorm.Open(pg, gormCfg)
	if err != nil {
		log.Fatalln("Can't connect to database:", err)
	} else {
		log.Printf("Succesfully connect to %s!\n", config.DBName)
	}

	return db
}
