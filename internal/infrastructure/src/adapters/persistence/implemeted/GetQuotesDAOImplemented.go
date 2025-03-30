package implemeted

import (
	"database/sql"
	"github.com/jufegare000/go-mlops/internal/infrastructure/src/adapters/persistence/entities/quotes"
)

type GetQuotesDAOImplemented struct {
	db *sql.DB
}

func NewGetQuotesDAOImplemented(db *sql.DB) *GetQuotesDAOImplemented {
	return &GetQuotesDAOImplemented{db: db}
}
func (dao *GetQuotesDAOImplemented) GetAllQuotes() ([]quotes.QuoteEntity, error) {
	rows, err := dao.db.Query("SELECT id, literary_work_id, description FROM quotes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []quotes.QuoteEntity

	for rows.Next() {
		var q quotes.QuoteEntity
		err := rows.Scan(&q.ID, &q.WorkId, &q.Description)
		if err != nil {
			return nil, err
		}
		result = append(result, q)
	}

	return result, nil
}
