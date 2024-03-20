// internal/interfaces/http/user_handler.go
package route

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/asifrahaman13/clean/src/internal/application"
)

type UserHandler struct {
    UserService *application.UserService
}

func NewUserHandler(userService *application.UserService) *UserHandler {
    return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Path[len("/user/"):]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := h.UserService.GetUserByID(id)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "User: %v", user)
}
