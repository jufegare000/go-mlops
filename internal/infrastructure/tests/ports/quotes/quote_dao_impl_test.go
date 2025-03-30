package quotes

import (
	"github.com/jufegare000/go-mlops/internal/infrastructure/src/adapters/persistence/implemeted"
	"github.com/jufegare000/go-mlops/internal/infrastructure/tests/ports/config"
	"testing"
)

func TestGetQuotesDAO_FindAll(t *testing.T) {
	db := config.SetupInMemoryDB(t)
	dao := implemeted.NewGetQuotesDAOImplemented(db)

	quotes, err := dao.GetAllQuotes()
	if err != nil {
		t.Fatalf("DAO error: %v", err)
	}

	if len(quotes) != 2 {
		t.Errorf("Expected 2 quotes, got %d", len(quotes))
	}
}
