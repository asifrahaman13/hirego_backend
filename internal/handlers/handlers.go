package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/asifrahaman13/hirego/internal/core/ports"
	"github.com/asifrahaman13/hirego/internal/helper"
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

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		panic(err)
	}

	fmt.Println("The result is", users)

	// helper.Response(c, 200, "hey there", nil)

	helper.JSONResponse(c, 200, users, nil)

}

func (h *userHandler) Signup(c *gin.Context) {

	// The BindJSON is used to extract the JSON data from the request body.
	var user *domain.User
	c.BindJSON(&user)

	// Call the signup service to signup the user.
	message, err := h.userService.Signup(user)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (h *userHandler) Login(c *gin.Context) {
	// The BindJSON is used to extract the JSON data from the request body.
	var user *domain.User
	c.BindJSON(&user)

	// Call the signup service to signup the user.
	message, err := h.userService.Login(user)

	if err != nil {
		panic(err)
	}

	// Next call the helper function to send the response.
	helper.JSONResponse(c, 200, message, nil)
}

func (h *userHandler) ProtectedRoute(c *gin.Context) {

	// Extract the authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		helper.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	// Split the header value to get the token
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		helper.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	accessToken := parts[1]

	message, err := h.userService.ProtectedRoute(accessToken)

	if err != nil {
		panic(err)
	}

	helper.JSONResponse(c, 200, message, nil)
}
