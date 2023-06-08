BUILDHASH = $(shell git rev-parse --verify HEAD | cut -c 1-7)
VERSION = 1.0.0

docker-up:
	@docker-compose build
	@docker-compose up -d

docker-down:
	@docker-compose down