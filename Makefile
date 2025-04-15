.PHONY: help build tidy run test format lint clean docker-build docker-exec docker-run docker-lint docker-stop

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

APP_NAME=chinese-checkers
DOCKER_IMAGE=$(APP_NAME)-dev
DOCKER_CONTAINER=$(APP_NAME)-dev-instance

# -- Go --

build: ## Build the Go binary (runs locally).
	mkdir -p bin/
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

tidy: ## Tidy `go.mod` and `go.sum` files (runs locally).
	go mod tidy

run: build ## Build and run the application (runs locally).
	./bin/$(APP_NAME)

test: ## Run Go tests (runs locally).
	go test ./...

format: ## Format Go code using `go fmt` (runs locally).
	go fmt ./...

lint: ## Run `staticcheck` linter (requires local installation).
	@command -v staticcheck >/dev/null 2>&1 || { echo "Error: please install Staticcheck."; exit 1; }
	staticcheck ./...

clean: ## Remove the built binary.
	rm -Rf bin

# -- Docker --

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

docker-lint: ## Run the `staticcheck` linter inside the Docker container.
	$(MAKE) docker-exec TARGET=lint

docker-exec:  ## Execute any `make` target (e.g., `test`, `build`) inside the Docker container. - Exemple: make docker-exec TARGET=test
	docker exec $(DOCKER_CONTAINER) make $(TARGET)


