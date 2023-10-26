package main

import (
	"fmt"
	"httpserver/data"
	"httpserver/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Println(data.BooksData)

	r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/books/", handlers.ListBooks).Methods("GET")
	r.HandleFunc("/books/", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	http.ListenAndServe(":9000", r)
}
