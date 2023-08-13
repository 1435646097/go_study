package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

	ginServer.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "hello world",
		})
	})
	ginServer.Run(":8080")
}
