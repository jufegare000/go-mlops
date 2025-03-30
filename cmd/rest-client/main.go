package main

import (
	"github.com/jufegare000/go-mlops/internal/infrastructure/src/adapters/bootstrap/rest"
	restclientadapter "ports/main"
	"ports/main/config"
)

func main() {
	config.SetUpEnvVariables()
	rest.InitializeAllServices()
	restclientadapter.StartRestServer()
}
