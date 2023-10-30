package main

import (
	"crudwithmysqlgorm/connections"
	"crudwithmysqlgorm/models"
	"crudwithmysqlgorm/utilities"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	utilities.CreateExample(db)
	utilities.QueryExample(db)

}

func init() {
	// Initiliase environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error occurred while initialising env variables %v", err)
	}

	// Migrate DB Models
	db = connections.DBConnector()

	// db.AutoMigrate(&models.Person{})
	if err := db.AutoMigrate(&models.Person{}, &models.Engineer{}); err != nil {
		log.Fatalf("Error while migrating models %v", err)
	}
}
