package movie

import (
	"context"
	"fmt"

	modelmovies "github.com/c3llus/monolith-movie-service/model/movies"
)

func (s *MoviesService) GetMoviesByTitle(ctx context.Context, title string) ([]*modelmovies.Movie, error) {

	movies, err := s.OmdbRepo.GetMoviesByTitle(ctx, title)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v", err))
		return nil, err
	}

	return movies, nil
}

func (s *MoviesService) GetMovieByImdbId(ctx context.Context, imdbId string) (*modelmovies.Movie, error) {

	movie, err := s.OmdbRepo.GetMovieByImdbId(ctx, imdbId)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v", err))
		return nil, err
	}

	return movie, nil
}
