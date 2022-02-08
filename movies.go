package events

import "github.com/gofrs/uuid"

type MovieReleased struct {
	MovieId   uuid.UUID `json:"movieId"`
	Title     string    `json:"title"`
	PosterUrl string    `json:"posterUrl"`
	Genres    []string  `json:"genres"`
	Runtime   int       `json:"runtime"`
	Rating    string    `json:"rating"`
}

type MovieRated struct {
	RatingId uuid.UUID `json:"ratingId"`
	MovieId  string    `json:"movieId"`
	Username string    `json:"username"`
	Rating   int       `json:"rating"`
}
