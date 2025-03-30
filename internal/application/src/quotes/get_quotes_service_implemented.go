package quotes

func GetAllQuotes() ([]QuoteDTO, error) {
	quotes := make([]QuoteDTO, 0)
	quotes = append(quotes, QuoteDTO{1, "Mock", "Blabla"})
	return quotes, nil
}

func GetQuotesByAuthor(author string) ([]QuoteDTO, error) {
	quotes := make([]QuoteDTO, 0)
	quotes = append(quotes, QuoteDTO{1, "Mock", "Blabla"})
	return quotes, nil
}
