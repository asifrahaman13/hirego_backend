package routes

import (
	"github.com/asifrahaman13/hirego/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
   router.GET("/users", handlers.UserHandler.GetUsers)
   router.POST("/signup", handlers.UserHandler.Signup)
   router.POST("/login", handlers.UserHandler.Login)
   router.GET("/userdata", handlers.UserHandler.ProtectedRoute)
}