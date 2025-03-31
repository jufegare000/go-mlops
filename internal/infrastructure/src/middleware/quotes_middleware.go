package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jufegare000/go-mlops/internal/application/quotes"
)

const quotesServiceKey = "quotesService"

func InjectQuotesService(service quotes.GetQuotesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(quotesServiceKey, service)
		c.Next()
	}
}

func GetQuotesServiceFromContext(c *gin.Context) quotes.GetQuotesService {
	if value, exists := c.Get(quotesServiceKey); exists {
		if service, ok := value.(quotes.GetQuotesService); ok {
			return service
		}
	}
	return nil
}
