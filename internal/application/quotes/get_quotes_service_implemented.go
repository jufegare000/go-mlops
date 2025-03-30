package quotes

func NewGetQuotesService(dao GetQuotesDao) GetQuotesService {
	return &getQuotesServiceImpl{dao: dao}
}

type getQuotesServiceImpl struct {
	dao GetQuotesDao
}

func (g getQuotesServiceImpl) GetAllQuotes() ([]QuoteDTO, error) {
	quotes := make([]QuoteDTO, 0)
	quotes = append(quotes, QuoteDTO{1, "Mock", "Blabla"})
	return quotes, nil
}

func (g getQuotesServiceImpl) GetQuotesByAuthor(author string) ([]QuoteDTO, error) {
	quotes := make([]QuoteDTO, 0)
	quotes = append(quotes, QuoteDTO{1, "Mock", "Blabla"})
	return quotes, nil
}
