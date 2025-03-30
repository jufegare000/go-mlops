package rest_client_adapter

import (
	"fmt"
	"ports/main/config"
	"ports/main/db"
	"ports/main/rest_client_adapter"
)

func StartRestServer() {
	fmt.Println("Welcome to this application on")
	config.SetUpEnvVariables()
	db.ConnectToDB()
	rest_client_adapter.RestRouter()
}
