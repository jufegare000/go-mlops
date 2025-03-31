package quotes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"quotes"
)

func QuotesControllerImplemented(router *gin.Engine) {
	router.GET("/quotes", getQuotesList)
}

func getQuotes() ([]quotes.QuoteDTO, error) {
	return quotes.GetAllQuotes()
}

func getQuotesList(ginContext *gin.Context) {
	list, _ := getQuotes()
	jsonData := parseQuotes(list)
	ginContext.JSON(200, jsonData)
}

func parseQuotes(list []quotes.QuoteDTO) string {
	jsonData, _ := json.Marshal(list)
	jsonString := string(jsonData)
	return jsonString
}
