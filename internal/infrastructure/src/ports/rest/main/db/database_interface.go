package db

import mysqlconf "msyql"

func ConnectToDB() {
	mysqlconf.InitMySQLDB()
}
