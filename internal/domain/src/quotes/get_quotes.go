package quotes

type GetQuotes interface {
	Execute() ([]Quote, error)
}
