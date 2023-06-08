package http

import (
	"context"
	"net/http"
	"time"

	"github.com/c3llus/monolith-movie-service/common/errors"
	modelhttp "github.com/c3llus/monolith-movie-service/model/http"
	modelmovies "github.com/c3llus/monolith-movie-service/model/movies"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetDetails(c *gin.Context) {
	ctx := c.Request.Context()

	// default timeout
	ctx, cancelCtx := context.WithTimeout(ctx, time.Second*10)
	defer cancelCtx()

	// get searched title
	imdbId := c.Param("id")

	// search movies
	movie, err := s.movieService.GetMovieByImdbId(ctx, imdbId)
	if err != nil {
		c.JSON(errors.GetHTTPStatus(err), gin.H{"error": errors.GetErrorMessage(err)})
		return
	}

	c.JSON(http.StatusOK, modelhttp.SearchResponse{
		Movies: []modelhttp.Movie{parseInternalMovieToClientMovie(movie)},
	})
}

func parseInternalMovieToClientMovie(movie *modelmovies.Movie) modelhttp.Movie {
	return modelhttp.Movie{
		Id:         movie.Id,
		Title:      movie.Title,
		Year:       movie.Year,
		ImdbId:     movie.ImdbId,
		PosterLink: movie.PosterLink,
	}
}
