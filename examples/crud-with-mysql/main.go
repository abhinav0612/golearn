package main

import (
	"crudwithmysql/handlers"
	"crudwithmysql/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error occured while loading env variables: %v", err)
	}

	handlers.Initialize()

	// Add
	author := models.Author{
		AuthorId: 4,
		Name:     "Jane Doe",
		Email:    "jane@gmail.com",
	}

	id := handlers.Add(author)

	// Get
	fmt.Println("Jane details")
	jane := handlers.Get("Jane Doe")
	fmt.Println(jane)

	// List
	fmt.Println("All authors details")
	authors := handlers.List()
	for _, author := range authors {
		fmt.Println(author)
	}

	// Delete
	handlers.Delete(id)

	// List
	fmt.Println("All authors details")
	authors = handlers.List()
	for _, author := range authors {
		fmt.Println(author)
	}
}
