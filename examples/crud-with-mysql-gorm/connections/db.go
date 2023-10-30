package connections

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnector() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dbString := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dbString)
	db, err := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error occurred while connecting to database %v", err)
	}
	return db
}
