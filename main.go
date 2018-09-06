package main

import (
	"github.com/dentych/dinner-dash/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	router := gin.Default()

	unprotectedApiRouter := router.Group("/api")
	unprotectedApiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "OK")
	})

	protectedApiRouter := router.Group("/api", middleware.AuthRequired())
	protectedApiRouter.GET("/john", func(c *gin.Context) {
		user, _ := c.Get("User")
		c.JSON(200, "Authenticated as: " + user.(string))
	})

	router.Run(":8080")
}
