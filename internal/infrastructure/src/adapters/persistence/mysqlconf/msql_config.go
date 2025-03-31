package mysqlconf

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMySQLDB() *sql.DB {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	name := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Error connecting to DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("❌ DB unreachable: %v", err)
	}

	fmt.Println("✅ Connected to MySQL")
	return DB
}
