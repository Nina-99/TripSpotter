package router

import (
	"net/http"

	"github.com/Nina-99/TripSpotter/backend/controller"
	"github.com/Nina-99/TripSpotter/backend/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	api := service.Group("/api")
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)

		usersRouter := api.Group("/users")
		usersRouter.GET("/", userController.FindAll)
		usersRouter.Use(middleware.JWTAuthMiddleware())
		{
			// usersRouter.GET("/", userController.FindAll)
			usersRouter.DELETE("/:id", userController.Delete)
			usersRouter.PUT("/:id", userController.Update)
		}
	}

	return service
}
