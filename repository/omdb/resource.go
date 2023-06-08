package omdb

import (
	"context"

	"github.com/c3llus/monolith-movie-service/app/resource/client"

	modelmovies "github.com/c3llus/monolith-movie-service/model/movies"
)

type OmdbRepositoryProvider interface {
	GetMoviesByTitle(ctx context.Context, title string) ([]*modelmovies.Movie, error)
	GetMovieByImdbId(ctx context.Context, imdbId string) (*modelmovies.Movie, error)
}

type OmdbRepository struct {
	client omdbRepoClient
}

func NewOmdbRepository(
	client client.HTTPClientProvider,
) *OmdbRepository {
	return &OmdbRepository{
		client: omdbRepoClient{
			client: client,
		},
	}
}

func (o *OmdbRepository) GetMoviesByTitle(ctx context.Context, title string) ([]*modelmovies.Movie, error) {

	movies, err := o.client.getMoviesByTitle(ctx, title)
	if err != nil {
		return nil, err
	}

	return movies, nil

}

func (o *OmdbRepository) GetMovieByImdbId(ctx context.Context, imdbId string) (*modelmovies.Movie, error) {

	movie, err := o.client.getMovieByImdbId(ctx, imdbId)
	if err != nil {
		return nil, err
	}

	return movie, nil

}
