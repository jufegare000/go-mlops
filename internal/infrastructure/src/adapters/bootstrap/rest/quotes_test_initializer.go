package rest

import (
	"database/sql"
	quotes2 "github.com/jufegare000/go-mlops/internal/application/quotes"
	"github.com/jufegare000/go-mlops/internal/infrastructure/src/adapters/persistence/implemeted"
)

func InitializeQuotesService(db *sql.DB) quotes2.GetQuotesService {
	dao := implemeted.NewGetQuotesDAOImplemented(db)
	service := quotes2.NewGetQuotesService(dao)
	return service
}
