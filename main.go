package main

import (
	"fmt"
	"log"

	"github.com/asifrahaman13/hirego/internal/core/services"
	"github.com/asifrahaman13/hirego/internal/handlers"
	"github.com/asifrahaman13/hirego/internal/repository"
	"github.com/asifrahaman13/hirego/internal/routes"
	"github.com/gin-gonic/gin"
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

	msg, err := repository.InitializeDB("sadfsdf")

	if err != nil {
		panic(err)
	}

	fmt.Println(msg)
	docRep := repository.UserRepo.Initialize(msg)

	users := service.InitializeUserService(docRep)

	handlers.UserHandler.Initialize(users)

	return nil
}
