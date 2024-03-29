package handlers

import (
	// "fmt"
	"github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/asifrahaman13/hirego/internal/core/ports"
	"github.com/asifrahaman13/hirego/internal/helper"
	// "github.com/asifrahaman13/hirego/internal/helper"
	"github.com/gin-gonic/gin"
)

var UserHandler *userHandler

type userHandler struct {
	// *base
	userService ports.UserService
}

func (h *userHandler) Initialize(userserv ports.UserService) {
	UserHandler = &userHandler{
		userService: userserv,
	}
}

func (h *userHandler) Signup(c *gin.Context) {
	var user domain.User
	c.BindJSON(&user)

	message, err := h.userService.Signup(user)

	if err != nil {
		panic(err)
	}
	helper.JSONResponse(c, 200, message, nil)
	
}

// func (h *userHandler) Signup(c *gin.Context) {

// 	// The BindJSON is used to extract the JSON data from the request body.
// 	var user domain.User
// 	c.BindJSON(&user)

// 	// Call the signup service to signup the user.
// 	message, err := h.userService.Signup(user)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Next call the helper function to send the response.
// 	helper.JSONResponse(c, 200, message, nil)
// }

// func (h *userHandler) Login(c *gin.Context) {
// 	// The BindJSON is used to extract the JSON data from the request body.
// 	var user domain.User
// 	c.BindJSON(&user)

// 	// Call the signup service to signup the user.
// 	message, err := h.userService.Login(user)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Next call the helper function to send the response.
// 	helper.JSONResponse(c, 200, message, nil)
// }

// func (h *userHandler) ProtectedRoute(c *gin.Context) {

// 	message := c.MustGet("user_email").(map[string]interface{})

// 	fmt.Println("The message obtained is", message)

// 	helper.JSONResponse(c, 200, message, nil)
// }

// func (h *userHandler) UserInformation(c *gin.Context) {
// 	// The BindJSON is used to extract the JSON data from the request body.
// 	var user domain.UserInformation
// 	c.BindJSON(&user)

// 	// Update the user's email based on the user_email from the context
// 	userMap := c.MustGet("user_email").(map[string]interface{})

// 	user.Email = userMap["username"].(string)

// 	fmt.Println("The user information is", user)

// 	// Call the signup service to signup the user.
// 	message, err := h.userService.UserInformation(user)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Next call the helper function to send the response.
// 	helper.JSONResponse(c, 200, message, nil)
// }

// func (h *userHandler) GetUserInformation(c *gin.Context) {

// 	// Update the user's email based on the user_email from the context
// 	userMap := c.MustGet("user_email").(map[string]interface{})

// 	// Call the signup service to signup the user.
// 	message, err := h.userService.GetUserInformation(userMap["username"].(string))

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Next call the helper function to send the response.
// 	helper.JSONResponse(c, 200, message, nil)
// }

// func (h *userHandler) SetUserWrorkInformation(c *gin.Context) {

// 	var workinformation domain.WorkInformation
// 	c.BindJSON(&workinformation)

// 	// Update the user's email based on the user_email from the context
// 	userMap := c.MustGet("user_email").(map[string]interface{})

// 	// Call the signup service to signup the user.
// 	message, err := h.userService.SetUserWrorkInformation(userMap["username"].(string), workinformation)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Next call the helper function to send the response.
// 	helper.JSONResponse(c, 200, message, nil)
// }

// func (h *userHandler) GetUserWorkInformation(c *gin.Context) {

// 	var username domain.UserName
// 	c.BindJSON(&username)


// 	// Call the signup service to signup the user.
// 	message, err := h.userService.GetUserWorkInformation(username)

// 	if err != nil {
// 		panic(err)
// 	}

// 	// Next call the helper function to send the response.
// 	helper.JSONResponse(c, 200, message, nil)
// }