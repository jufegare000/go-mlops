package httpadapter

import (
	"github.com/gin-gonic/gin"
)

func RestRouter() {
	router := gin.Default()

	router.GET("/hello", func(ginContext *gin.Context) {
		ginContext.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	_ = router.Run(":8080")
}
