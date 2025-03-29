package rest_client_adapter

import "github.com/gin-gonic/gin"

func HelloWorldRestController(router *gin.Engine) {
	router.GET("/hello", getHelloWorld)
}

func getHelloWorld(ginContext *gin.Context) {
	ginContext.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
