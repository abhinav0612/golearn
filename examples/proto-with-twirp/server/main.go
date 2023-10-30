package main

import (
	"context"
	"fmt"
	"net/http"
	"protowithtwirp/rpc/notes"
	"time"

	"github.com/twitchtv/twirp"
)

type notesService struct {
	Notes     []notes.Note
	CurrentId int32
}

func (s *notesService) CreateNote(ctx context.Context, request *notes.CreateNoteRequest) (*notes.Note, error) {
	fmt.Println("Inside CreateNote")
	if len(request.Text) < 4 {
		return nil, twirp.InvalidArgument.Error("Text length should be greater than 4")
	}
	note := notes.Note{
		Id:        s.CurrentId,
		Text:      request.Text,
		CreatedAt: time.Now().UnixMilli(),
	}

	s.Notes = append(s.Notes, note)
	s.CurrentId++

	return &note, nil
}

func (s *notesService) GetAllNotes(ctx context.Context, request *notes.GetAllNotesRequest) (*notes.GetAllNotesResponse, error) {
	fmt.Println("Inside GetAllNotes")
	allNotes := make([]*notes.Note, 0)
	for _, note := range s.Notes {
		n := note
		allNotes = append(allNotes, &n)
	}
	return &notes.GetAllNotesResponse{
		Notes: allNotes,
	}, nil
}

func main() {

	notesServer := notes.NewNoteServiceServer(&notesService{})

	mux := http.NewServeMux()
	mux.Handle(notesServer.PathPrefix(), notesServer)

	err := http.ListenAndServe(":9000", notesServer)
	if err != nil {
		panic(err)
	}
}
