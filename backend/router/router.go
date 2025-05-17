package router

import (
	"net/http"

	"github.com/Nina-99/TripSpotter/backend/controller"
	"github.com/Nina-99/TripSpotter/backend/middleware"
	"github.com/Nina-99/TripSpotter/backend/repository"
	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UsersRepository, authenticationController *controller.AuthenticationController, usersController *controller.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	authenticationRouter := router.Group("/authentication")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	usersRouter := router.Group("/users")
	usersRouter.GET("", middleware.DeserializeUser(userRepository), usersController.GetUsers)

	return service
}
