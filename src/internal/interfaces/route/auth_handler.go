package route

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/asifrahaman13/hirego/src/internal/application"
	"github.com/asifrahaman13/hirego/src/internal/domain"
	"github.com/asifrahaman13/hirego/src/internal/helper"
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

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// login function
	var requestBody map[string]interface{}

	// Decode the incoming request json into golang data type map structure.
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		panic(err)
	}

	// Extract the username from the request body.
	username, ok := requestBody["email"].(string)

	// If the username is not found in the request body, return an error. otherwise create the token.
	access_token, err := helper.CreateToken(username)

	if err != nil {
		panic(err)

	}

	if !ok {
		panic("username not found")
	}
    
	// Return the access token in the response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(access_token)
}



func (h *AuthHandler) VerifyToken(w http.ResponseWriter, r *http.Request) {
	
	// verify token function

	// Extract out the token string from the authorization header. 
	tokenString := r.Header.Get("Authorization")
    
	// Extract the token from the token string without the name Bearer.
	extracted_token := tokenString[7:]
    
	// Print the extracted token.
	fmt.Println("tokenString: ", extracted_token)
    
	// Verify the token and return the response. 
	claims, err := helper.VerifyToken(extracted_token)

	if err != nil {
		panic(err)
	}
    
	// Return the claims in the response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(claims)
}