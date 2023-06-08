package movie

import (
	"context"

	// repositories
	omdbrepo "github.com/c3llus/monolith-movie-service/repository/omdb"

	// models
	modelmovies "github.com/c3llus/monolith-movie-service/model/movies"
)

// inversion
type MoviesServiceProvider interface {
	GetMovieByImdbId(ctx context.Context, id string) (*modelmovies.Movie, error)
	GetMoviesByTitle(ctx context.Context, title string) ([]*modelmovies.Movie, error)
}

type MoviesService struct {
	OmdbRepo omdbrepo.OmdbRepositoryProvider
	// movieRepo
}

func NewMovieService(
	omdbRepo omdbrepo.OmdbRepositoryProvider,
) *MoviesService {
	return &MoviesService{
		OmdbRepo: omdbRepo,
	}
}
