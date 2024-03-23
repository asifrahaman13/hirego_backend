// internal/interfaces/http/routes.go
package route

import (
	"github.com/asifrahaman13/clean/src/internal/application"
	"github.com/asifrahaman13/clean/src/internal/interfaces"
)

func SetupRoutes(userService *application.UserService) *interfaces.Routes {

	// Define the user handler and the associated routes. In the next step combine them to the parent routes. 
	userHandler := NewUserHandler(userService)
	userRoutes := setupUserRoutes(userHandler)

	allRoutes := userRoutes

	routes := &interfaces.Routes{
		Routes: allRoutes,
	}

	return routes
}
