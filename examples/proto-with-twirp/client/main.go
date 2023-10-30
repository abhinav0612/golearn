package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"protowithtwirp/rpc/notes"
)

func main() {
	client := notes.NewNoteServiceProtobufClient("http://localhost:9000", &http.Client{})

	ctx := context.Background()

	_, err := client.CreateNote(ctx, &notes.CreateNoteRequest{Text: "Sample Notes"})
	if err != nil {
		log.Fatal(err)
	}

	allNotes, err := client.GetAllNotes(ctx, &notes.GetAllNotesRequest{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Notes")
	for _, note := range allNotes.Notes {
		log.Println(note.Id)
	}
}
