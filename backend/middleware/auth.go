package middleware

import (
	"fmt"
	"github.com/dentych/dinner-dash/logging"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
)

const headerAuthorization = "Authorization"

// None header: eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(headerAuthorization)
		if len(authHeader) < 1 {
			logging.Error.Println("Error:", fmt.Errorf("auth header missing"))
			abortRequest(c)
			return
		}

		token, err := jwt.ParseWithClaims(authHeader, &jwt.StandardClaims{}, hmacSecret())
		if err != nil {
			logging.Error.Println("Error:", err)
			abortRequest(c)
			return
		}

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			c.Set("user", claims.Subject)
			if claims.ExpiresAt == 0 {
				logging.Error.Println("Error:", fmt.Errorf("ExpiresAt is 0, which usually means it hasn't been set in the JWT"))
				abortRequest(c)
				return
			}
			c.Next()
		} else {
			logging.Error.Println("Error:", fmt.Errorf("invalid claims or token. Claims OK: %v. Token valid: %v", ok, token.Valid))
			abortRequest(c)
		}
	}
}

func hmacSecret() jwt.Keyfunc {
	signingSecret := os.Getenv("JWT_SECRET")
	if len(signingSecret) == 0 {
		signingSecret = "secret"
	}

	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signingSecret), nil
	}
}

func abortRequest(context *gin.Context) {
	context.JSON(401, "unauthorized")
	context.Abort()
}
