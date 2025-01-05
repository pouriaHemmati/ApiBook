package data

import "time"

type Book struct {
	ID        int64     `json:"id"`
	CreateAt  time.Time `json:"-"`
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"`
	Page      int       `json:"page,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	Version   int32     `json:"version"`
}
