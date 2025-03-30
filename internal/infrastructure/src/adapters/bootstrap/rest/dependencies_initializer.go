package rest

import "ports/main/db"

func InitializeAllServices() {
	dbConnection := db.ConnectToDB()
	InitializeQuotesService(dbConnection)
}
