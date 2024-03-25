package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{

	return func(c *gin.Context) {
		fmt.Println(" Updated one. -------------------------------------------------------------------------------------------------------------------Auth middleware called")

		c.Next()
	}
}
