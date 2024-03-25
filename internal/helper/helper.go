// package helper

// import (
// 	"fmt"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// )

// var secretKey = []byte("SECRET_KEY")

// func CreateToken(username string) (interface{}, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
// 		jwt.MapClaims{
// 			"username": username,
// 			"exp":      time.Now().Add(time.Hour * 24).Unix(),
// 		})

// 	tokenString, err := token.SignedString(secretKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	tokenData := map[string]interface{}{
// 		"access_token": tokenString,
// 	}

// 	return tokenData, nil
// }

// func VerifyToken(tokenString string) (map[string]interface{}, error) {

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, fmt.Errorf("invalid token")
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return nil, fmt.Errorf("invalid token claims")
// 	}

// 	return claims, nil
// }

package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code int, message string, body interface{}) {
	if message == "" {
		message = http.StatusText(code)
	}
	c.JSON(
		code,
		gin.H{
			"code":       code,
			"statusText": http.StatusText(code),
			"message":    message,
			"body":       body,
		},
	)
}

func JSONResponse(c *gin.Context, code int, message interface{}, body interface{}) {
	c.JSON(
		code,
		gin.H{
			"code":    code,
			"message": message,
			"body":    body,
		},
	)
}
