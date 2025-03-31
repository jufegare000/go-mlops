package main

import (
	mysqlconf "msyql"
	rest_client_adapter "ports/main"
	"ports/main/config"
)

func main() {
	config.SetUpEnvVariables()
	mysqlconf.InitMySQLDB()
	rest_client_adapter.StartRestServer()
}
