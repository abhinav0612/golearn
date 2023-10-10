package main

import (
	"fmt"
	"httpserver/handlers"
	"net/http"
)

func main() {
	fmt.Println("Starting Simple HTTP Server...")

	http.HandleFunc("/", handlers.WelcomeHandler)

	http.ListenAndServe(":8080", nil)
}
