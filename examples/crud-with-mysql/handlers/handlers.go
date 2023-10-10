package handlers

import (
	"crudwithmysql/connectors"
	"crudwithmysql/models"
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func Add(author models.Author) int64 {
	query := `INSERT INTO authors (name, email) VALUES (?, ?)`
	res, err := db.Exec(query, author.Name, author.Email)
	if err != nil {
		fmt.Println(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}
	return id
}

func Delete(id int64) {
	query := `DELETE FROM authors where author_id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Get(name string) models.Author {
	var author models.Author
	query := "SELECT * FROM authors where name = ?"
	err := db.QueryRow(query, name).Scan(&author.AuthorId, &author.Name, &author.Email)
	if err != nil {
		fmt.Println(err.Error())
	}
	return author
}

func List() []models.Author {
	var authors []models.Author
	query := "SELECT * FROM authors"
	rows, _ := db.Query(query)
	defer rows.Close()
	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.AuthorId, &author.Name, &author.Email)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			authors = append(authors, author)
		}
	}
	return authors
}

func Initialize() {
	db = connectors.DBConnector()
	if err := db.Ping(); err != nil {
		log.Fatalf("Error occured while seting up database connection %v", err)
		return
	}
	fmt.Println("Database connected...")

}
