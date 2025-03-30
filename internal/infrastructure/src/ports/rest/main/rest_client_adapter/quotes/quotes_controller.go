package quotes

import (
	"github.com/gin-gonic/gin"
)

func QuotesControllerImplemented(router *gin.Engine) {
	router.GET("/quotes", getQuotesList)
}

func getQuotesList(ginContext *gin.Context) {
	ginContext.JSON(200, gin.H{
		"message": "Hello from quotes list!",
	})
}
