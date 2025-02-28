package data

import (
	"time"
)

type Movie struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int       `json:"year,omitempty"`
	Runtime   int       `json:"runtime,omitempty,string"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int       `json:"version"`
}
