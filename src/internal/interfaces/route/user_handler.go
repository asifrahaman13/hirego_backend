// internal/interfaces/http/user_handler.go
package route

import (
	"encoding/json"
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

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This api is called.")

	user, err := h.UserService.GetUsers()

	if err != nil {
		panic(err)
	}


    // The http response from mutex is not able to understand the golang struct object.
    // As a result we need to convert it into the json data. 
	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to encode user to JSON", http.StatusInternalServerError)
		return
	}

    // Specify that you want the content to be JSON format. 
	w.Header().Set("Content-Type", "application/json")

    // Return the json data. No need to have the return keyword.
	w.Write(userJSON)
}
