package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // - directive always hide this in the output
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"` // omitempty shows the output if and only if the field is not empty
	Runtime   int32     `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
