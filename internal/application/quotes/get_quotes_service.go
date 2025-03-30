package quotes

type GetQuotesService interface {
	GetAllQuotes() ([]QuoteDTO, error)
	GetQuotesByAuthor(author string) ([]QuoteDTO, error)
}
