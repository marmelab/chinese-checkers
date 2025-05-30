# Executables (local)
DOCKER_COMP = docker compose

# Paths
WEB_DIR = cd web

# Docker containers
PHP_CONT = $(DOCKER_COMP) exec php
PHP_RUN = $(DOCKER_COMP) run --rm php
CHINESE_CHECKERS_RUN = $(DOCKER_COMP) -f compose.cli.yaml run --rm chinese-checkers

# Docker compose files.
DOCKER_COMPOSE_WEB_MAIN=compose.yaml
DOCKER_COMPOSE_WEB_PROD=compose.prod.yaml
DOCKER_COMPOSE_CLI_MAIN=compose.cli.yaml

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
.PHONY        : help install build-cli build-api start-cli deps lint vet check clean test-engine up prepare-web-app start-web-app start-web-app-dev down logs sh bash test install-e2e test-e2e build-mobile-app start-mobile-app-dev build-admin-panel start-admin-panel-dev composer vendor composer-install composer-install-dev sf cc generate-jwt-keys

## —— Chinese Checkers ♟️ ——————————————————————————————————————————————————————
help: ## Outputs this help screen.
	@grep -E '(^[a-zA-Z0-9\./_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

install: ## Builds the Docker images for cli and web apps.
	@$(DOCKER_COMP) build --pull

## —— CLI app ⌨️ ———————————————————————————————————————————————————————————————

build-cli: deps ## Build the Go binary (inside Docker).
	@$(CHINESE_CHECKERS_RUN) go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

build-api: deps ## Build the API binary (inside Docker).
	@$(CHINESE_CHECKERS_RUN) go build -o bin/game-api ./api/$(APP_NAME)

start-cli: ## Run the CLI application.
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

test-engine: ## Test the game engine.
	@$(CHINESE_CHECKERS_RUN) go test -v ./...

## —— Web app 🌐 ———————————————————————————————————————————————————————————————

up: ## Start web app in detached mode.
	@$(DOCKER_COMP) up --detach

up-production: ## Start web app in detached mode for production.
	@$(DOCKER_COMP) -f $(DOCKER_COMPOSE_WEB_MAIN) -f $(DOCKER_COMPOSE_WEB_PROD) up --detach

prepare-web-app: ## Prepare web application files.
	@mkdir -p web/config/jwt
	@touch web/config/jwt/public.jwk

start-web-app: install build-api prepare-web-app composer-install up-production generate-jwt-keys cc ## Build and start the web application for production.
start-web-app-dev: install build-api prepare-web-app composer-install-dev up generate-jwt-keys ## Build and start the web application in dev mode.

down: ## Stop web app.
	@$(DOCKER_COMP) down --remove-orphans

logs: ## Show live logs.
	@$(DOCKER_COMP) logs --tail=0 --follow

sh: ## Connect to the FrankenPHP container.
	@$(PHP_CONT) sh

bash: ## Connect to the FrankenPHP container via bash so up and down arrows go to previous commands.
	@$(PHP_CONT) bash

test: test-engine ## Run tests with phpunit, pass the parameter "c=" to add options to phpunit, example: make test c="--group e2e --stop-on-failure".
	@$(eval c ?=)
	@$(DOCKER_COMP) run --rm -e APP_ENV=test php bin/phpunit $(c)

install-e2e: ## Install dependencies for end to end tests with playwright.
	@(cd web && yarn install && yarn playwright install --with-deps)

test-e2e: ## Run end to end tests with playwright.
	@(cd web && yarn playwright test)


## —— Mobile app 📱 ————————————————————————————————————————————————————————————

build-mobile-app: ## Build the mobile app for production use with web app.
	@(cd web && yarn install && yarn build)

start-mobile-app-dev: ## Run the mobile app in dev mode.
	@(cd web && yarn install && yarn dev)

## —— Admin panel 👑 ———————————————————————————————————————————————————————————

build-admin-panel: ## Build the admin panel for production.
	@(cd web && yarn install)
	@(cd admin && yarn install && yarn build)

start-admin-panel-dev: ## Run the admin panel in dev mode.
	@(cd web && yarn install)
	@(cd admin && yarn install && yarn dev)

## —— Composer 🧙 ——————————————————————————————————————————————————————————————
composer: ## Run composer, pass the parameter "c=" to run a given command, example: make composer c='req symfony/orm-pack'.
	@$(eval c ?=)
	@$(COMPOSER) $(c)

vendor: ## Install vendors according to the current composer.lock file.
vendor: c=install --prefer-dist --no-dev --no-progress --no-scripts --no-interaction
vendor: composer

composer-install: ## Install web app dependencies according to the current composer.lock file.
	$(COMPOSER_INSTALL) --no-dev

composer-install-dev: ## Install dependencies for testing and developing the web app.
	$(COMPOSER_INSTALL)

## —— Symfony 🎵 ———————————————————————————————————————————————————————————————
sf: ## List all Symfony commands or pass the parameter "c=" to run a given command, example: make sf c=about.
	@$(eval c ?=)
	@$(SYMFONY) $(c)

cc: c=c:c ## Clear the cache.
cc: sf

generate-jwt-keys: ## Generate keypair for JWT.
	@$(SYMFONY) lexik:jwt:generate-keypair --skip-if-exists
	@$(PHP_CONT) sh -c "/app/jose.phar key:load:key /app/config/jwt/public.pem > /app/config/jwt/public.jwk"
