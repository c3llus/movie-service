package movie

import "github.com/segmentio/ksuid"

type Movie struct {
	Id         ksuid.KSUID `json:"id"`
	Title      string      `json:"title"`
	Year       string      `json:"year"`
	ImdbId     string      `json:"imdb_id"`
	PosterLink string      `json:"poster_link"`
}
