package main

import (
	"fmt"
	"log"
	"time"

	"github.com/asifrahaman13/hirego/internal/core/services"
	"github.com/asifrahaman13/hirego/internal/handlers"
	"github.com/asifrahaman13/hirego/internal/repository"
	"github.com/asifrahaman13/hirego/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    if err := run(); err != nil {
        log.Fatal(err)
    }

    parent_route := gin.Default()

    // Allow all origins and proxy all requests.
    parent_route.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

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
