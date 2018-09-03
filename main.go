package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	router := gin.Default()

	apiRouter := router.Group("/api")
	apiRouter.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "OK")
	})

	router.Run(":8080")
}
