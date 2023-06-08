package main

import (
	"github.com/c3llus/monolith-movie-service/app/http"
	"github.com/c3llus/monolith-movie-service/common/configs"
)

func main() {

	// get configs
	cfg := configs.GetConfig()

	// init server
	httpServer := http.InitServer(&cfg)

	// serve HTTP server
	httpServer.ServeHTTP()
}
