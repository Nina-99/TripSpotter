package config

import (
	"fmt"

	"github.com/Nina-99/TripSpotter/backend/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *Config) *gorm.DB {
	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName,
	)

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println(" Conneected Successfully to the DataBase")
	return db
}
