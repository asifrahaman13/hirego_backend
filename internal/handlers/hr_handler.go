package handlers

import (
	"github.com/asifrahaman13/hirego/internal/core/ports"
)

var HRHandler *hrHandler

type hrHandler struct {
	hrService ports.HRService
}

func (h *hrHandler) Initialize(hrserv ports.HRService) {
	HRHandler = &hrHandler{
		hrService: hrserv,
	}
}
