package handlers

import (
	// "fmt"
	// "github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/asifrahaman13/hirego/internal/core/ports"
	// "github.com/asifrahaman13/hirego/internal/helper"
	// "github.com/gin-gonic/gin"
)

var HRHandler *userHandler

type hrHandler struct {
	// *base
	userService ports.HRService
}

func (h *hrHandler) Initialize(hrserv ports.HRService) {
	HRHandler = &userHandler{
		userService: hrserv,
	}
}