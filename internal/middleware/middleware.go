package middleware

import (
	"fmt"
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
			return
		}

		// Split the header value to get the token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.JSONResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
			return
		}

		accessToken := parts[1]

		user_email, err := helper.VerifyToken(accessToken)

		if err != nil {
			panic(err)
		}

		fmt.Println("The message is", user_email)

		c.Set("user_email", user_email)

		c.Next()
	}
}
