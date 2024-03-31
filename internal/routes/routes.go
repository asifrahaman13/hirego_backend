package routes

import (
	"github.com/asifrahaman13/hirego/internal/handlers"
	"github.com/asifrahaman13/hirego/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupV1Routes(router *gin.Engine) {

	v1 := router.Group("/auth")
	{
		v1.POST("/signup", handlers.UserHandler.Signup)
		v1.POST("/login", handlers.UserHandler.Login)
		v1.POST("/hrsignup", handlers.HRHandler.Signup)
		v1.POST("/hrlogin", handlers.HRHandler.Login)
	}
}

func SetupV2Routes(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	v2 := router.Group("/user")
	{ // Add all the available middlewares which needs to be attached to the router.
		for _, middleware := range middlewares {
			v2.Use(middleware)
		}
		v2.POST("/userprofileinformation", handlers.UserHandler.SetUserProfileInformation)
		v2.GET("/userprofileinformation", handlers.UserHandler.GetProfileInformation)
		v2.POST("/userworkinformation", handlers.UserHandler.SetUserWrorkInformation)
		v2.POST("/jobposting", handlers.HRHandler.GetJobPosting)
		v2.GET("/alljobposting", handlers.HRHandler.GetAllJobPosting)
		
	}
}

func SetupHrRoutes(router *gin.Engine, middlewares ...gin.HandlerFunc) {
	hr := router.Group("/hr")
	{
		for _, middleware := range middlewares {
			hr.Use(middleware)
		}
		hr.POST("/hrprofileinformation", handlers.HRHandler.SetHrProfileInformation)
		hr.GET("/hrprofileinformation", handlers.HRHandler.GetProfileInformation)
		hr.POST("/jobposting", handlers.HRHandler.JobPosting) // Route for HR to post a job
		hr.GET("/jobposting", handlers.HRHandler.HrSpecificJobPosting) // Get all the job postings by the particular HR
		hr.POST("/userpublicinformation", handlers.UserHandler.GetUserWorkInformation)
	}
}

// func SetupV3Routes(router *gin.Engine) {
// 	v3 := router.Group("/public")
// 	{

// 		//Sample public route
// 	}
// }

func InitializeRoutes(router *gin.Engine) {

	// Add the middleware to the parent route.
	SetupV1Routes(router)
	SetupV2Routes(router, middleware.AuthMiddleware())
	SetupHrRoutes(router, middleware.AuthMiddleware())
	// SetupV3Routes(router)
}
