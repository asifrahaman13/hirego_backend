// internal/interfaces/http/routes.go
package route

import (
    "github.com/asifrahaman13/clean/src/internal/application"
    "github.com/asifrahaman13/clean/src/internal/interfaces"
)

func SetupRoutes(userService *application.UserService) *interfaces.Routes {
    userHandler := NewUserHandler(userService)

    routes := &interfaces.Routes{
        Routes: []*interfaces.Route{
            {
                Path:    "/user/",
                Handler: userHandler.GetUserByID,
                Method:  "GET",
            },
            // Add more routes here
        },
    }

    return routes
}
