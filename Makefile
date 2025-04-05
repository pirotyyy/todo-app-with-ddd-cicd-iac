DOCKER_COMPOSE := docker compose

up:
	$(DOCKER_COMPOSE) up -d
down:
	$(DOCKER_COMPOSE) down
build:
	$(DOCKER_COMPOSE) build --no-cache
lint:
	go tool golangci-lint run ./...