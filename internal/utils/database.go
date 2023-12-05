package utils

import (
	"e-commerce/config"
	"e-commerce/pkg/gorm"

	"github.com/gofiber/fiber/v2/log"
)

func ConnectToDB() {
	dbConfig := gorm.DatabaseConfig{
		Host:     config.GetString("database.host"),
		Port:     config.GetInt("database.port"),
		Database: config.GetString("database.dbname"),
		Username: config.GetString("database.username"),
		Password: config.GetString("database.password"),
		Driver:   config.GetString("database.driver"),
	}

	if err := gorm.Connect(dbConfig); err != nil {
		log.Fatalf("failed to connect to db:%s", err)
	}
}
