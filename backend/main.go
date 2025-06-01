package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Nina-99/TripSpotter/backend/config"
	"github.com/Nina-99/TripSpotter/backend/controller"
	"github.com/Nina-99/TripSpotter/backend/repository"
	"github.com/Nina-99/TripSpotter/backend/router"
	"github.com/Nina-99/TripSpotter/backend/service"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	// db := config.DB
	validate := validator.New()
	config.ConnectDB()

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(config.DB)

	//Init Service
	userService := service.NewUserServiceImpl(userRepository, validate)

	//Init controller
	userController := controller.NewUserController(userService)

	routes := router.NewRouter(userController)

	routes.Run()

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error to Create Listener", err)
	}
	defer listener.Close()
}
