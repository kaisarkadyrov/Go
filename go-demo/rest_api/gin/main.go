package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct {
	Message string `json:"message"`
}

var test = Test{
	Message: "Hello World",
}

func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, test)
}

func posting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"method": "POST"})
}

func putting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"method": "PUT"})
}

func deleting(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func main() {
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{ // instead of 200 we can write http.StatusOK, H stands for HashMap
			"message": "OK",
		})
	})

	server.GET("/hello", helloWorld)       // endpoint
	server.POST("/somepost", posting)      // endpoint
	server.PUT("/someput", putting)        // endpoint
	server.DELETE("/somedelete", deleting) // endpoint

	server.Run(":8080")
}
