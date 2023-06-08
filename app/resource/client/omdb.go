package client

import (
	"github.com/c3llus/monolith-movie-service/common/configs"
)

func NewOmdbClient(
	cfg *configs.OmdbClient,
) HTTPClientProvider {
	return &HTTPClient{
		host: cfg.Host,
		key:  cfg.Key,
	}
}
