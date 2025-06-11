package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Nina-99/TripSpotter/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
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
	db.AutoMigrate(&models.Image{})
	db.AutoMigrate(&models.Marker{})
	db.AutoMigrate(&models.Shapefile{})
	db.AutoMigrate(&models.Review{})
	DB = db
}
