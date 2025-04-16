.PHONY: help install start stop test build run deps lint clean docker-exec

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

APP_NAME=chinese-checkers
DOCKER_SERVICE=chinese-checkers

# -- Main goals --

install:  ## Build the Docker development image.
	docker-compose build

start: ## Start the development Docker container.
	docker-compose up -d --force-recreate 

stop: ## Stop the running development Docker container.
	docker-compose down

test: ## Run Go tests (inside Docker).
	docker-compose exec $(DOCKER_SERVICE) go test -v ./...

# -- Other DX goals 

build: deps ## Build the Go binary (inside Docker).
	docker-compose exec $(DOCKER_SERVICE) go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

run: ## run the application.
	docker-compose exec $(DOCKER_SERVICE) go run cmd/$(APP_NAME)/main.go 

deps: ## Tidy `go.mod` and `go.sum` files (inside Docker).
	docker-compose exec $(DOCKER_SERVICE) go mod tidy

lint: ## Run `staticcheck` linter (inside Docker).
	docker-compose exec $(DOCKER_SERVICE) staticcheck ./...

clean: ## Remove the built binary (inside Docker).
	docker-compose exec $(DOCKER_SERVICE) rm -f bin/$(APP_NAME)

docker-exec: ## Run a command inside the docker container - Example: make docker-exec CMD="ls -l"
	docker-compose exec $(DOCKER_SERVICE) $(CMD)


