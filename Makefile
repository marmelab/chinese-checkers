# Executables (local)
DOCKER_COMP = docker compose

# Paths
WEB_DIR = cd web

# Docker containers
PHP_CONT = $(DOCKER_COMP) exec php
PHP_RUN = $(DOCKER_COMP) run --rm php
CHINESE_CHECKERS_RUN = $(DOCKER_COMP) run --rm chinese-checkers

# Executables
PHP      = $(PHP_CONT) php
COMPOSER = $(PHP_CONT) composer
SYMFONY  = $(PHP) bin/console

COMPOSER_INSTALL = $(PHP_RUN) composer install --prefer-dist --no-progress --no-scripts --no-interaction

# Application
APP_NAME=chinese-checkers
GO_PACKAGE=github.com/marmelab/chinese-checkers

# Misc
.DEFAULT_GOAL = help
.PHONY        : help install build run deps lint vet check clean up start down logs sh bash test composer vendor composer-install composer-install-test sf cc

## —— Chinese Checkers ♟️ ——————————————————————————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9\./_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

install: ## Builds the Docker images for cli and web apps.
	@$(DOCKER_COMP) build --pull

## —— CLI app ⌨️ ———————————————————————————————————————————————————————————————

build: deps ## Build the Go binary (inside Docker).
	@$(CHINESE_CHECKERS_RUN) go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

run: ## Run the application.
	@$(CHINESE_CHECKERS_RUN) go run $(GO_PACKAGE)/cmd/$(APP_NAME) $(APP_ARGS)

deps: ## Tidy `go.mod` and `go.sum` files (inside Docker).
	@$(CHINESE_CHECKERS_RUN) go mod tidy

lint: ## Run `staticcheck` linter (inside Docker).
	@$(CHINESE_CHECKERS_RUN) staticcheck ./...

vet: ## Run `go vet` (inside Docker).
	@$(CHINESE_CHECKERS_RUN) go vet ./...

check: lint vet ## Run `staticcheck` and `go vet` (inside Docker).

clean: ## Remove the built binary (inside Docker).
	@$(CHINESE_CHECKERS_RUN) rm -f bin/$(APP_NAME)

## —— Web app 🌐 ———————————————————————————————————————————————————————————————

up: ## Start web app in detached mode.
	@$(DOCKER_COMP) up --detach

start: build up ## Build and start the web application

down: ## Stop web app
	@$(DOCKER_COMP) down --remove-orphans

logs: ## Show live logs
	@$(DOCKER_COMP) logs --tail=0 --follow

sh: ## Connect to the FrankenPHP container
	@$(PHP_CONT) sh

bash: ## Connect to the FrankenPHP container via bash so up and down arrows go to previous commands
	@$(PHP_CONT) bash

test: ## Run tests with phpunit, pass the parameter "c=" to add options to phpunit, example: make test c="--group e2e --stop-on-failure"
	@$(CHINESE_CHECKERS_RUN) go test -v ./...
	@$(eval c ?=)
	@$(DOCKER_COMP) run --rm -e APP_ENV=test php bin/phpunit $(c)


## —— Composer 🧙 ——————————————————————————————————————————————————————————————
composer: ## Run composer, pass the parameter "c=" to run a given command, example: make composer c='req symfony/orm-pack'
	@$(eval c ?=)
	@$(COMPOSER) $(c)

vendor: ## Install vendors according to the current composer.lock file
vendor: c=install --prefer-dist --no-dev --no-progress --no-scripts --no-interaction
vendor: composer

composer-install: ## Install web app dependencies according to the current composer.lock file
	$(COMPOSER_INSTALL) --no-dev

composer-install-test: ## Install dependencies for testing the web app
	$(COMPOSER_INSTALL)

## —— Symfony 🎵 ———————————————————————————————————————————————————————————————
sf: ## List all Symfony commands or pass the parameter "c=" to run a given command, example: make sf c=about
	@$(eval c ?=)
	@$(SYMFONY) $(c)

cc: c=c:c ## Clear the cache
cc: sf
