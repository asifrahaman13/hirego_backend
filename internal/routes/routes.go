package routes

import (
	"github.com/asifrahaman13/hirego/internal/handlers"
	"github.com/asifrahaman13/hirego/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupV1Routes(router *gin.Engine, middlewares ...gin.HandlerFunc) {

	v1 := router.Group("/auth")
	{
		// Add all the available middlewares which needs to be attached to the router.
		for _, middleware := range middlewares {
			v1.Use(middleware)
		}
		v1.GET("/users", handlers.UserHandler.GetUsers)
		v1.POST("/signup", handlers.UserHandler.Signup)
		v1.POST("/login", handlers.UserHandler.Login)
		v1.GET("/userdata", handlers.UserHandler.ProtectedRoute)
	}
}

func SetupV2Routes(router *gin.Engine) {
	v2 := router.Group("/v2")
	{
		v2.GET("/users", handlers.UserHandler.GetUsers)
		// More routes to be added here
	}
}

func InitializeRoutes(router *gin.Engine) {

	// Add the middleware to the parent route.
	SetupV1Routes(router, middleware.AuthMiddleware())
	SetupV2Routes(router)
}
