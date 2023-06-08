package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/c3llus/monolith-movie-service/common/configs"

	// services
	movieservice "github.com/c3llus/monolith-movie-service/service/movie"

	// clients
	client "github.com/c3llus/monolith-movie-service/app/resource/client"

	// repositories
	omdbrepo "github.com/c3llus/monolith-movie-service/repository/omdb"
)

type Server struct {
	movieService movieservice.MoviesServiceProvider
}

func InitServer(
	config *configs.Config,
) *Server {

	// init client(s)
	omdbClient := client.NewOmdbClient(&config.OmdbClient)

	// init repositor(ies)
	omdbRepository := omdbrepo.NewOmdbRepository(omdbClient)

	// init service(s)
	movieService := movieservice.NewMovieService(
		omdbRepository,
	)

	return &Server{
		movieService: movieService,
	}
}

func (s *Server) ServeHTTP() {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: s.RegisterHandler(),
	}

	// serve the server
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	s.gracefulShutdown(httpServer)

}

// gracefulShutdown implements graceful shutdown
func (s *Server) gracefulShutdown(httpServer *http.Server) {

	ctx, stopCtx := context.WithCancel(context.Background())

	go func() {

		signals := make(chan os.Signal, 1)

		// wait for the sigterm
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// We received an os signal, shut down.
		if err := httpServer.Shutdown(ctx); err != nil {
			fmt.Println(fmt.Errorf("graceful shutdown failed: %w", err))
		} else {
			fmt.Println("graceful shutdown succeed")
		}

		stopCtx()

	}()

	<-ctx.Done()
}

func (s *Server) RegisterHandler() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/ping", s.ping)

	r.GET("/search/:title", s.Search)
	r.GET("/details/:id", s.GetDetails)

	return r
}

func (s *Server) ping(c *gin.Context) {

	ctx := c.Request.Context()

	// add tracing&observability
	ctx, cancelCtx := context.WithTimeout(ctx, time.Second*10)
	defer cancelCtx()

	c.JSON(http.StatusOK, "pong")
}
