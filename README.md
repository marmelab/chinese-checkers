# Chinese Checkers CLI

A simple two-player Chinese Checkers game playable in the terminal, written in Go.

## Features ✨

- Standard Go project structure (`cmd/`, `internal/`).
- Web application (`web/`).
- `Makefile` for automating common development tasks (build, test, lint, run).
- `Dockerfile` for a consistent Go development environment with necessary tools (including `staticcheck`).
- `web/Dockerfile` for easy deployment and local tests.
- Linting configured with `staticcheck`.

## 🛠️ Requirements

Before you begin, ensure you have the following installed:

- **Docker:** To build and run the development container. [Install Docker](https://docs.docker.com/get-docker/)
- **Make:** To use the Makefile automation. (Usually pre-installed on Linux/macOS, may need installation on Windows).
- **Yarn:** To install and run e2e tests.
- **Corepack:** To install latest stable yarn. Run `sudo corepack enable` before any `yarn install`.

_(Note: PHP, Go and staticcheck do NOT need to be installed locally.)_

## 🚀 Getting Started

Clone the repository, then:

1.  **Install the development Docker Image:**
    This creates the container images with the Go environment and tools, and the PHP environment. Run this once initially, or when the `Dockerfile` or `web/Dockerfile` changes.

    ```bash
    make install
    ```

2.  **Run the CLI:**
    This runs the CLI game inside the container.

    ```bash
		# Show help.
		APP_ARGS="--help" make run
		# Run with arguments.
		APP_ARGS="--state-file tests/states/ongoing-game.json --move c1,d1" make run
    ```

3.  **Test:**
    Runs all PHP and Go tests inside containers.

    ```bash
    make test
    ```

4.  **Using Make Commands:**
    With the container running, use standard `make` commands directly from your terminal. They automatically execute _inside_ the Docker container:

    ```bash
     —— Chinese Checkers ♟️ ——————————————————————————————————————————————————————
    help                           Outputs this help screen.
    install                        Builds the Docker images for cli and web apps.
     —— CLI app ⌨️ ———————————————————————————————————————————————————————————————
    build-cli                      Build the Go binary (inside Docker).
    build-api                      Build the API binary (inside Docker).
    start-cli                      Run the CLI application.
    deps                           Tidy `go.mod` and `go.sum` files (inside Docker).
    lint                           Run `staticcheck` linter (inside Docker).
    vet                            Run `go vet` (inside Docker).
    check                          Run `staticcheck` and `go vet` (inside Docker).
    clean                          Remove the built binary (inside Docker).
    test-engine                    Test the game engine.
     —— Web app 🌐 ———————————————————————————————————————————————————————————————
    up                             Start web app in detached mode.
    up-production                  Start web app in detached mode for production.
    prepare-web-app                Prepare web application files.
    start-web-app                  Build and start the web application for production.
    start-web-app-dev              Build and start the web application in dev mode.
    down                           Stop web app.
    logs                           Show live logs.
    sh                             Connect to the FrankenPHP container.
    bash                           Connect to the FrankenPHP container via bash so up and down arrows go to previous commands.
    test                           Run tests with phpunit, pass the parameter "c=" to add options to phpunit, example: make test c="--group e2e --stop-on-failure".
    install-e2e                    Install dependencies for end to end tests with playwright.
    test-e2e                       Run end to end tests with playwright.
     —— Mobile app 📱 ————————————————————————————————————————————————————————————
    build-mobile-app               Build the mobile app for production use with web app.
    start-mobile-app-dev           Run the mobile app in dev mode.
     —— Admin panel 👑 ———————————————————————————————————————————————————————————
    build-admin-panel              Build the admin panel for production.
    start-admin-panel-dev          Run the admin panel in dev mode.
     —— Composer 🧙 ——————————————————————————————————————————————————————————————
    composer                       Run composer, pass the parameter "c=" to run a given command, example: make composer c='req symfony/orm-pack'.
    vendor                         Install vendors according to the current composer.lock file.
    composer-install               Install web app dependencies according to the current composer.lock file.
    composer-install-dev           Install dependencies for testing and developing the web app.
     —— Symfony 🎵 ———————————————————————————————————————————————————————————————
    sf                             List all Symfony commands or pass the parameter "c=" to run a given command, example: make sf c=about.
    cc                             Clear the cache.
    generate-jwt-keys              Generate keypair for JWT.
    ```

    _(See Makefile Targets below or run `make help` for all commands)_

## 🧱 Project Structure (main folders)

```bash
cmd/chinese-checkers/ # Main entry point
internal/game/        # Core game logic
bin/                  # Binary file
web/                  # Web application (Symfony)
web/app/              # Mobile application (React)
admin/                # Admin panel (React Admin)
```

## 👷 Testing

### Web application

You should not need to configure anything to start the web application in development mode:

```shell
make start-web-app-dev
```

You can then access the web application on [`https://localhost`](https://localhost).

### Mobile app

```shell
# Start the web application (backend).
make start-web-app-dev
# Start the mobile application.
make start-mobile-app-dev
```

You can then access the web application on [`http://localhost:5173/app`](http://localhost:5173/app).

### Admin panel

```shell
# Start the web application (backend).
make start-web-app-dev
# Start the admin panel.
make start-admin-panel-dev
```

## ⚙️ Game data generation

You can generate fake game data to test various game situations, statistics and so on.

### Generate accounts

```shell
# Generate 20 fake accounts.
make sf c="app:generate-fake-accounts"
# Generate 50 fake accounts.
make sf c="app:generate-fake-accounts --count 50"
```

### Generate games

```shell
# Generate 20 ongoing games, 5 won games and 5 pending games (awaiting an opponent).
make sf c="app:generate-fake-games"
# Generate 50 ongoing games, 10 won games and 3 pending games (awaiting an opponent).
make sf c="app:generate-fake-games --ongoing-games 50 --won-games 10 --pending-games 3"
```

## 🧪 Running tests

### Unit tests

`make install` is optional is you already did it before.

```shell
make install
make start-web-app-dev
make test
```

### End to end tests

`make install` is optional is you already did it before.

```shell
make install
make install-e2e # Install playwright and playwright browsers locally.
make start-web-app # Start the web application to run tests on it.
make test-e2e # Run end to end tests.
```

## 🌐 Deploy web application

To deploy the web application in production, copy `.env.sample` to `.env`
and set the application secret using the provided command.

```shell
openssl rand -hex 64
```

Build the mobile app with `make build-mobile-app`.
You'll need to run it again if anything has changed in the mobile app.

You can then start the application with `make start-web-app`.
