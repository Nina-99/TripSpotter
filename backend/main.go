package main

import (
	"github.com/Nina-99/TripSpotter/backend/config"
	"github.com/Nina-99/TripSpotter/backend/controller"
	"github.com/Nina-99/TripSpotter/backend/models"
	"github.com/Nina-99/TripSpotter/backend/repository"
	"github.com/Nina-99/TripSpotter/backend/router"
	"github.com/Nina-99/TripSpotter/backend/service"
	"github.com/go-playground/validator/v10"
)

func main() {

	//Database
	db := config.ConnectDB()
	validate := validator.New()

	db.Table("users").AutoMigrate(&models.User{})

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)

	//Init Service
	userService := service.NewUserServiceImpl(userRepository, validate)

	//Init controller
	userController := controller.NewUserController(userService)

	routes := router.NewRouter(userController)

	routes.Run()
}
