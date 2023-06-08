package http

import "github.com/segmentio/ksuid"

// our client-facing response for Movie
type Movie struct {
	Id         ksuid.KSUID `json:"id"`
	Title      string      `json:"title"`
	Year       int         `json:"year"`
	ImdbId     string      `json:"imdb_id"`
	PosterLink string      `json:"poster_link"`
	// Type       string `json:"type"`
}