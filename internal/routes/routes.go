package routes

import (
	"github.com/asifrahaman13/hirego/internal/handlers"
	"github.com/asifrahaman13/hirego/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupV1Routes(router *gin.Engine) {

	v1 := router.Group("/auth")
	{

		v1.POST("/signup", handlers.UserHandler.Signup)
		v1.POST("/login", handlers.UserHandler.Login)

	}
}

func SetupV2Routes(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	v2 := router.Group("/user")
	{ // Add all the available middlewares which needs to be attached to the router.
		for _, middleware := range middlewares {
			v2.Use(middleware)
		}
		v2.POST("/userinformation", handlers.UserHandler.UserInformation)
		v2.GET("/account-information", handlers.UserHandler.GetUserInformation)
		v2.POST("/userworkinformation", handlers.UserHandler.SetUserWrorkInformation)
		// More routes to be added here
	}
}

func InitializeRoutes(router *gin.Engine) {

	// Add the middleware to the parent route.
	SetupV1Routes(router)
	SetupV2Routes(router, middleware.AuthMiddleware())
}
