package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/c3llus/monolith-movie-service/common/errors"

	modelhttp "github.com/c3llus/monolith-movie-service/model/http"
	modelmovies "github.com/c3llus/monolith-movie-service/model/movies"
)

func (s *Server) Search(c *gin.Context) {

	ctx := c.Request.Context()

	// default timeout
	ctx, cancelCtx := context.WithTimeout(ctx, time.Second*10)
	defer cancelCtx()

	// get searched title
	title := c.Param("title")

	// search movies
	movies, err := s.movieService.GetMoviesByTitle(ctx, title)
	if err != nil {
		c.JSON(errors.GetHTTPStatus(err), gin.H{"error": errors.GetErrorMessage(err)})
		return
	}

	c.JSON(http.StatusOK, modelhttp.SearchResponse{
		Movies: parseInternalMoviesToClientMovies(movies),
	})
}

// TODO: move somewhere else
func parseInternalMoviesToClientMovies(movie []*modelmovies.Movie) []modelhttp.Movie {
	var movies []modelhttp.Movie
	for i := range movie {
		movies = append(movies, parseInternalMovieToClientMovie(movie[i]))
	}
	return movies
}
