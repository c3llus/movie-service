package configs

type Config struct {
	OmdbClient OmdbClient `yaml:"omdb_client"`
}

type OmdbClient struct {
	Host string `yaml:"host"`
	Key  string `yaml:"key"`
}
