// // cmd/main.go
// package main

// import (
// 	"github.com/asifrahaman13/hirego/src/internal/application"
// 	"github.com/asifrahaman13/hirego/src/internal/infrastructure"
// 	"github.com/asifrahaman13/hirego/src/internal/interfaces/route"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	userRepository := infrastructure.NewUserRepository()
// 	userService := application.NewUserService(userRepository)
// 	userHandler := route.NewUserHandler(userService)

// 	// Register routes
// 	userHandler.RegisterRoutes(r.Group("/api"))

// 	r.Run(":8080")
// }
