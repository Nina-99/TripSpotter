package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Nina-99/TripSpotter/backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Falied to connected of DataBase", err)
	}

	fmt.Println(" Conneected Successfully to the DataBase")
	db.AutoMigrate(&models.User{})
	return db
}
