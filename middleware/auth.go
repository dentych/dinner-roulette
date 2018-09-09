package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
)

const headerAuthorization = "Authorization"

type MyClaims struct {
	User string `json:"user"`
	jwt.StandardClaims
}

// None header: eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(headerAuthorization)
		if len(authHeader) < 1 {
			abortRequest(c, fmt.Errorf("auth header missing"))
			return
		}

		token, err := jwt.ParseWithClaims(authHeader, &jwt.StandardClaims{}, hmacSecret())
		if err != nil {
			abortRequest(c, err)
			return
		}

		err = token.Claims.Valid()
		if err != nil {
			abortRequest(c, err)
			return
		}

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			c.Set("User", claims.Subject)
			c.Next()
		} else {
			abortRequest(c, fmt.Errorf("invalid claims or token. Claims OK: %v. Token valid: %v", ok, token.Valid))
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

func abortRequest(context *gin.Context, err error) {
	context.JSON(401, "unauthorized")
	fmt.Println("Error while parsing token: ", err)
	context.Abort()
}
