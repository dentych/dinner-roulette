package middleware

import (
	"fmt"
	"github.com/dentych/dinner-dash/logging"
	"github.com/dentych/dinner-dash/security"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const headerAuthorization = "Authorization"

// None header: eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(headerAuthorization)
		authHeader = strings.TrimSpace(authHeader)

		if !strings.HasPrefix(authHeader, "Bearer ") || len(authHeader) < 1 {
			logging.Error.Println("Error:", fmt.Errorf("auth header missing"))
			abortRequest(c)
			return
		}

		authHeader = strings.Split(authHeader, " ")[1]

		if claims, err := security.ValidateJwt(authHeader); err != nil {
			logging.Error.Printf("Error validating JWT: %v", err)
			abortRequest(c)
			return
		} else {
			c.Set("userid", claims.UserId)
			c.Next()
		}
	}
}

func abortRequest(context *gin.Context) {
	context.JSON(http.StatusUnauthorized, "unauthorized")
	context.Abort()
}
