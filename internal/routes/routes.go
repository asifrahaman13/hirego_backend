package routes

import (
    "github.com/asifrahaman13/hirego/internal/handlers"
    "github.com/gin-gonic/gin"
)

func SetupV1Routes(router *gin.Engine) {
    v1 := router.Group("/auth")
    {
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
    SetupV1Routes(router)
    SetupV2Routes(router)
}
