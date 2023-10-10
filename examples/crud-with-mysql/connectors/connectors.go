package connectors

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnector() *sql.DB {
	driver := "mysql"
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbString := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open(driver, dbString)
	if err != nil {
		panic(err.Error())
	}
	return db
}
