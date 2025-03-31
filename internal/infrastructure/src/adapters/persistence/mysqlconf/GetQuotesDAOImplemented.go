package mysqlconf

import (
	"database/sql"
)

type GetQuotesDAOImplemented struct {
	db *sql.DB
}

func NewGetQuotesDAOImplemented(db *sql.DB) *GetQuotesDAOImplemented {
	return &GetQuotesDAOImplemented{db: db}
}

func (dao *GetQuotesDAOImplemented) GetAllQuotes() ([]QuoteEntity, error) {
	rows, err := dao.db.Query("SELECT id, literary_work_id, description FROM famous_quotes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []QuoteEntity

	for rows.Next() {
		var q QuoteEntity
		err := rows.Scan(&q.ID, &q.WorkId, &q.Description)
		if err != nil {
			return nil, err
		}
		result = append(result, q)
	}

	return result, nil
}
