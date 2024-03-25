package main

import (
	"fmt"
	"github.com/asifrahaman13/hirego/internal/core/services"
	"github.com/asifrahaman13/hirego/internal/handlers"
	"github.com/asifrahaman13/hirego/internal/repository"
	"github.com/asifrahaman13/hirego/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	parent_route := gin.Default()

	routes.InitializeRoutes(parent_route)

	log.Fatal(parent_route.Run())
}

func run() error {
	// handlers.Base.Initialize()

	db, err := repository.InitializeDB()

	if err != nil {
		panic(err)
	}

	fmt.Println(db)
	docRep := repository.UserRepo.Initialize(db)

	users := service.InitializeUserService(docRep)

	handlers.UserHandler.Initialize(users)

	return nil
}
