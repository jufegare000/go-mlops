package db

import (
	"database/sql"
	mysqlconf "msyql"
)

func ConnectToDB() *sql.DB {
	return mysqlconf.InitMySQLDB()
}
