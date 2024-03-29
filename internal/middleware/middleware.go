package middleware

import (

	"github.com/asifrahaman13/hirego/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helper.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort() // Abort the request
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort() // Abort the request
			return
		}

		accessToken := parts[1]
		userEmail, err := helper.VerifyToken(accessToken)
		if err != nil {
			helper.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
			c.Abort() // Abort the request
			return
		}

		// Set the user email in the context for later use
		c.Set("user_email", userEmail)

		c.Next()
	}
}
