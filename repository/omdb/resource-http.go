package omdb

import (
	"compress/gzip"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/c3llus/monolith-movie-service/app/resource/client"
	"github.com/c3llus/monolith-movie-service/common/errors"
	modelomdbclient "github.com/c3llus/monolith-movie-service/model/client/omdb"
	modelmovies "github.com/c3llus/monolith-movie-service/model/movies"
	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/ksuid"
)

type omdbRepoClientProvider interface {
	getMoviesByTitle(ctx context.Context, title string) ([]*modelmovies.Movie, error)
	getMovieByImdbId(ctx context.Context, imdbId string) (*modelmovies.Movie, error)
}

type omdbRepoClient struct {
	client client.HTTPClientProvider
}

func (r *omdbRepoClient) getMoviesByTitle(ctx context.Context, title string) ([]*modelmovies.Movie, error) {

	var (
		omdbResponse modelomdbclient.OmdbGetByTitle
		movies       []*modelmovies.Movie
	)

	// do http call
	resp, err := r.client.GetWithContext(ctx, client.Parameter{
		Path: "",
		QueryParams: map[string]string{
			"apiKey": r.client.GetKey(),
			"s":      title,
			"page":   "1", // FOR THE SAKE OF MVP, WE ONLY USE PAGE 1
		},
	})
	if err != nil {
		return nil, errors.AddTrace(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("OMDB API returned non 200 status code")
	}

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	// read to bytes data
	respBody, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, errors.AddTrace(err)
	}

	// unmarshall
	err = jsoniter.Unmarshal(respBody, &omdbResponse)
	if err != nil {
		return nil, errors.AddTrace(err)
	}

	// parse to our internal struct of Movie
	// TODO: make async parsing
	for _, omdbMovie := range omdbResponse.Search {

		year, err := strconv.Atoi(omdbMovie.Year)
		if err != nil {
			return nil, errors.AddTrace(err)
		}

		movies = append(movies, &modelmovies.Movie{
			Id:         ksuid.New(),
			Title:      omdbMovie.Title,
			Year:       year,
			ImdbId:     omdbMovie.ImdbId,
			PosterLink: omdbMovie.Poster,
		})
	}

	return movies, nil
}

func (r *omdbRepoClient) getMovieByImdbId(ctx context.Context, imdbId string) (*modelmovies.Movie, error) {

	var (
		omdbResponse modelomdbclient.OmdbGetByImdbIdResponse
		movie        *modelmovies.Movie
	)

	// do http call
	resp, err := r.client.GetWithContext(ctx, client.Parameter{
		Path: "/",
		QueryParams: map[string]string{
			"apiKey": r.client.GetKey(),
			"i":      imdbId,
		},
	})
	if err != nil {
		return nil, errors.AddTrace(err)
	}

	// read to bytes data
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.AddTrace(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("OMDB API returned non 200 status code")
	}

	// unmarshall
	err = jsoniter.Unmarshal(respBody, &omdbResponse)
	if err != nil {
		return nil, errors.AddTrace(err)
	}

	// parse to our internal struct of Movie
	// TODO: Make more modular
	year, err := strconv.Atoi(omdbResponse.Year)
	if err != nil {
		return nil, errors.AddTrace(err)
	}

	movie = &modelmovies.Movie{
		Id:         ksuid.New(),
		Title:      omdbResponse.Title,
		Year:       year,
		ImdbId:     omdbResponse.ImdbId,
		PosterLink: omdbResponse.Poster,
		// etc....
		// takut ga keburu mas wkwk
	}

	return movie, nil

}
