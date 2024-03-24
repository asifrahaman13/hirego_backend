package route

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/asifrahaman13/hirego/src/internal/application"
	"github.com/asifrahaman13/hirego/src/internal/domain"
)

type AuthHandler struct {
	AuthService *application.AuthService
}

func NewAuthHandler(authService *application.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	// signup function

	fmt.Print("signup called")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var user *domain.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	message, err := h.AuthService.Signup(user)

	if err != nil {
		panic(err)
	}

	fmt.Println("This service is also called.")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
