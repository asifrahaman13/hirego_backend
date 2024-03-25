package handlers

import (
	"fmt"

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

	helper.Response(c, 200, "hey there done", nil)

}
