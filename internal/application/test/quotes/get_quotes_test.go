package quotes

import (
	"github.com/jufegare000/go-mlops/internal/application/src/quotes"
	"testing"
)

func TestGetQuotes(t *testing.T) {
	t.Run("get a simple list of quotes", func(T *testing.T) {
		quotesResult, _ := quotes.GetAllQuotes()
		if len(quotesResult) < 1 {
			t.Error("there must be at least one quote")
		}
	})
}
