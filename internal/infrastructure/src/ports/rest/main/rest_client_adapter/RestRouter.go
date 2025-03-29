package rest_client_adapter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"ports/main/rest_client_adapter/quotes"
)

func RestRouter() {
	router := gin.Default()
	port := os.Getenv("APPLICATION_PORT")
	HelloWorldRestController(router)
	quotes.QuotesController(router)
	_ = router.Run(fmt.Sprintf(":%s", port))
}
