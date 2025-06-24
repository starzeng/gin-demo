package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	getTest(router)

	//err := router.Run("127.0.0.1:8081")
	err := router.Run(":8081")
	if err != nil {
		return
	}

}

func getTest(router *gin.Engine) gin.IRoutes {
	return router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"hello":   123,
		})
	})
}
