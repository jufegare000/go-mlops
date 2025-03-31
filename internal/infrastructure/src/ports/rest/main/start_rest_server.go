package rest_client_adapter

import (
	"fmt"
	"ports/main/rest_client_adapter"
)

func StartRestServer() {
	fmt.Println("Welcome to this application on")
	rest_client_adapter.RestRouter()
}
