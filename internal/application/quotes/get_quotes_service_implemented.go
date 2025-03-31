package quotes

import (
	mysqlconf "msyql"
)

import "github.com/jinzhu/copier"

func GetAllQuotes() ([]QuoteDTO, error) {
	quotes, _ := getDAO().GetAllQuotes()
	parsedList, _ := ConvertQuotes(quotes)
	return parsedList, nil
}

func getDAO() *mysqlconf.GetQuotesDAOImplemented {
	implementation := mysqlconf.NewGetQuotesDAOImplemented(mysqlconf.DB)
	return implementation
}

func ConvertQuotes(dtos []mysqlconf.QuoteEntity) ([]QuoteDTO, error) {
	var entities []QuoteDTO
	err := copier.Copy(&entities, &dtos)
	return entities, err
}
