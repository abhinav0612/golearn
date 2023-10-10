package models

import (
	"time"
)

type Posts struct {
	PostId    int
	Title     string
	Content   string
	CreatedAt time.Time
	AuthorId  int
}
