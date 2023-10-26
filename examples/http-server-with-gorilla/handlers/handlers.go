package handlers

import (
	"encoding/json"
	"httpserver/data"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type CreateBookRequestBody struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

type UpdateBookReqeustBody struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

func generateNewBookID() int {
	ids := []int{}
	for key := range data.BooksData {
		id, _ := strconv.Atoi(key)
		ids = append(ids, id)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ids)))
	if len(ids) == 0 {
		return 1
	}
	return ids[0] + 1
}

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	book, exists := data.BooksData[id]
	if !exists {
		respondWithJSON(w, http.StatusNotFound, "Book not found")
		return
	}
	respondWithJSON(w, http.StatusOK, book)
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, data.BooksData)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var requestBody CreateBookRequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		log.Fatalf("Error while decoding request body %v", err)
	}
	book := data.Book{
		Name:   requestBody.Name,
		Author: requestBody.Author,
	}
	newId := strconv.Itoa(generateNewBookID())
	data.BooksData[newId] = book
	respondWithJSON(w, http.StatusCreated, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var requestBody UpdateBookReqeustBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		log.Fatalf("Error while decoding request body %v", err)
	}
	id := requestBody.Id
	book, exists := data.BooksData[id]
	if !exists {
		respondWithJSON(w, http.StatusNotFound, "Book not found")
		return
	}
	bookData, ok := book.(data.Book)
	if !ok {
		respondWithJSON(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	bookData.Name = requestBody.Name
	bookData.Author = requestBody.Author
	data.BooksData[id] = bookData
	respondWithJSON(w, http.StatusOK, bookData)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	book, exists := data.BooksData[id]
	if !exists {
		respondWithJSON(w, http.StatusNotFound, "Book not found")
		return
	}
	delete(data.BooksData, id)
	respondWithJSON(w, http.StatusOK, book)
}
