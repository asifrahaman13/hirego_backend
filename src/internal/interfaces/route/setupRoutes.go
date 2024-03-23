package route

import "github.com/asifrahaman13/hirego/src/internal/interfaces"

// The routes here is dedicated to the users component only. So all the routes are now isolated.
func setupUserRoutes(userHandler *UserHandler) []*interfaces.Route {
    return []*interfaces.Route{
        {
            Path:    "/users",
            Handler: userHandler.GetUsers,
            Method:  "GET",
        },
        {
            Path:    "/user/",
            Handler: userHandler.GetUserByID,
            Method:  "GET",
        },
       
    }
}
