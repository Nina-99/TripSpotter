package router

import (
	"net/http"

	"github.com/Nina-99/TripSpotter/backend/config"
	"github.com/Nina-99/TripSpotter/backend/controller"
	"github.com/Nina-99/TripSpotter/backend/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	service := gin.Default()

	config.SetupCORS(service)

	service.Static("/uploads/images", "./uploads/images")

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
		api.GET("/role/:email", userController.GetUserByEmail)

		api.GET("/weather/forecast", controller.GetForecast)

		reviewRouter := api.Group("/reviews")
		reviewRouter.Use(middleware.JWTAuthMiddleware())
		{
			reviewRouter.POST("/upload", controller.SubmitReview)
			reviewRouter.POST("uploadImg", controller.UploadImage)
		}

		layersRouter := api.Group("/layers")
		layersRouter.Use(middleware.JWTAuthMiddleware())
		{
			layersRouter.POST("/upload", controller.UploadShapefile)
			layersRouter.GET("/", controller.GetAllLayers)
		}

		usersRouter := api.Group("/users")
		usersRouter.Use(middleware.JWTAuthMiddleware())
		{
			usersRouter.GET("/", userController.FindAll)
			usersRouter.DELETE("/:id", userController.Delete)
			usersRouter.PUT("/:id", userController.Update)
		}
	}

	return service
}
