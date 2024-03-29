package handlers

import (
	// "fmt"

	"github.com/asifrahaman13/hirego/internal/core/ports"
	// "github.com/asifrahaman13/hirego/internal/helper"
	// "github.com/gin-gonic/gin"
	// "github.com/asifrahaman13/hirego/internal/helper"
	// "github.com/gin-gonic/gin"
)

var HRHandler *hrHandler

type hrHandler struct {
	// *base
	hrService ports.HRService
}

func (h *hrHandler) Initialize(hrserv ports.HRService) {
	HRHandler = &hrHandler{
		hrService: hrserv,
	}
}

// func (h *hrHandler) Sample(c *gin.Context) {
// 	// h.hrService.Sample()
// 	helper.JSONResponse(c, 200, "Success", nil)
// }