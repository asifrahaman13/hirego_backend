// internal/interfaces/http/routes.go
package route

import (
	"github.com/asifrahaman13/hirego/src/internal/application"
	"github.com/asifrahaman13/hirego/src/internal/interfaces"
)

func SetupRoutes(userService *application.UserService, authSerive *application.AuthService) *interfaces.Routes {

	// Define the user handler and the associated routes. In the next step combine them to the parent routes. 
	userHandler := NewUserHandler(userService)
	userRoutes := setupUserRoutes(userHandler)


	authHandler:= NewAuthHandler(authSerive)
	authRoutes := AuthRoute(authHandler)

	allRoutes := userRoutes
	allRoutes = append(allRoutes, authRoutes...)

	routes := &interfaces.Routes{
		Routes: allRoutes,
	}

	return routes
}
