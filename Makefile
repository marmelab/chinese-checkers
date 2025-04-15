.PHONY: help build tidy run test format lint clean docker-build docker-exec docker-run docker-lint docker-stop

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

APP_NAME=chinese-checkers
DOCKER_IMAGE=$(APP_NAME)-dev
DOCKER_CONTAINER=$(APP_NAME)-dev-instance

# Internal target for building (only used by make inside docker)
build-internal:
	mkdir -p bin/
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

# Internal target for running (only used by make inside docker)
run-internal: build-internal
	./bin/$(APP_NAME)

build: ## Build the Go binary (inside Docker).
	docker exec $(DOCKER_CONTAINER) make build-internal

tidy: ## Tidy `go.mod` and `go.sum` files (inside Docker).
	docker exec $(DOCKER_CONTAINER) go mod tidy

run:  ## Build and run the application (inside Docker).
	docker exec -it $(DOCKER_CONTAINER) make run-internal

test: ## Run Go tests (inside Docker).
	docker exec $(DOCKER_CONTAINER) go test -v ./...

format: ## Format Go code using `go fmt` (inside Docker).
	docker exec $(DOCKER_CONTAINER) go fmt ./...

lint: ## Run `staticcheck` linter (inside Docker).
	docker exec $(DOCKER_CONTAINER) staticcheck ./...

clean: ## Remove the built binary (inside Docker).
	docker exec $(DOCKER_CONTAINER) rm -f bin/$(APP_NAME)

# -- Docker management --

docker-build:  ## Build the Docker development image.
	docker build -t $(DOCKER_IMAGE) .

docker-run: ## Start the development Docker container in the background.
	docker run -d --rm \
		--name $(DOCKER_CONTAINER) \
		-v .:/app \
		-w /app \
		$(DOCKER_IMAGE)

docker-stop: ## Stop the running development Docker container.
	docker stop $(DOCKER_CONTAINER)

docker-exec: ## Run a command inside the docker container - Example: make docker-exec CMD="ls -l"
	@docker exec $(DOCKER_CONTAINER) $(CMD)


