package main

import (
	"fmt"
	"github.com/asifrahaman13/hirego/internal/core/services"
	"github.com/asifrahaman13/hirego/internal/handlers"
	"log"
	// "github.com/asifrahaman13/hirego/internal/middleware"
	"github.com/asifrahaman13/hirego/internal/repository"
	"github.com/asifrahaman13/hirego/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	parent_route := gin.Default()

	// Add the middleware to the parent route.
	// Middleware should be added before initializing the routes.
	// parent_route.Use(middleware.AuthMiddleware())

	// Initialize the routes.
	routes.InitializeRoutes(parent_route)

	log.Fatal(parent_route.Run())
}

func run() error {
	// Initialize the database.
	db, err := repository.InitializeDB()

	if err != nil {
		panic(err)
	}

	fmt.Println(db)

	// Initialize the:
	// 1. User repository.
	// 2. User service.
	// 3. User handler.
	userRep := repository.UserRepo.Initialize(db)
	hrRep := repository.HRRepo.Initialize(db)

	users := service.InitializeUserService(userRep)
	hr_manager := service.InitializeHRService(hrRep)

	handlers.UserHandler.Initialize(users)
	handlers.HRHandler.Initialize(hr_manager)

	return nil
}
