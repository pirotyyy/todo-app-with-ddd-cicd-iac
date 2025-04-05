DOCKER_COMPOSE := docker compose

setup:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b ./bin v2.0.2
up:
	$(DOCKER_COMPOSE) up -d
down:
	$(DOCKER_COMPOSE) down
build:
	$(DOCKER_COMPOSE) build --no-cache
lint:
	./bin/golangci-lint run ./...