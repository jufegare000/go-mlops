package quotes

import (
	"github.com/jufegare000/go-mlops/internal/infrastructure/tests/ports/config"
	mysqlconf "msyql"
	"testing"
)

func TestGetQuotesDAO_FindAll(t *testing.T) {
	db := config.SetupInMemoryDB(t)
	dao := mysqlconf.NewGetQuotesDAOImplemented(db)

	quotes, err := dao.GetAllQuotes()
	if err != nil {
		t.Fatalf("DAO error: %v", err)
	}

	if len(quotes) != 2 {
		t.Errorf("Expected 2 quotes, got %d", len(quotes))
	}
}
