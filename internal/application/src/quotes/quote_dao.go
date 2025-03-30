package quotes

type GetQuotesDao interface {
	GetAllQuotes() ([]QuoteEntity, error)
}
