package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Nina-99/TripSpotter/backend/config"
	"github.com/Nina-99/TripSpotter/backend/controller"
	"github.com/Nina-99/TripSpotter/backend/helper"
	"github.com/Nina-99/TripSpotter/backend/models"
	"github.com/Nina-99/TripSpotter/backend/repository"
	"github.com/Nina-99/TripSpotter/backend/router"
	"github.com/Nina-99/TripSpotter/backend/service"
	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := config.ConnectDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&models.User{})

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)

	//Init Service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)

	//Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUsersController(userRepository)

	routes := router.NewRouter(userRepository, authenticationController, usersController)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
