package main

import (
	"fmt"
	"ports/main/config"
	"ports/main/db"
	"ports/main/rest_client_adapter"
)

func main() {
	config.SetUpEnvVariables()
	db.ConnectToDB()
	rest_client_adapter.RestRouter()
	fmt.Println("Welcome to this application on")
}
