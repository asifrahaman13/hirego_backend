package route

import "github.com/asifrahaman13/hirego/src/internal/interfaces"


func AuthRoute(authHandler *AuthHandler) []*interfaces.Route {
	// sample route file

	return []*interfaces.Route{
		{
			Path: "/auth/signup",
			Handler: authHandler.Signup,
			Method: "POST",
		},
		{
			Path: "/auth/login",
			Handler: authHandler.Login,
			Method: "POST",
		},{
			Path: "/auth/username",
			Handler: authHandler.VerifyToken,
			Method: "GET",
		},
	}
}