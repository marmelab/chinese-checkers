# Chinese Checkers CLI

A simple two-player Chinese Checkers game playable in the terminal, written in Go.

## Features

- Standard Go project structure (`cmd/`, `internal/`).
- `Makefile` for automating common tasks (build, test, lint, run, Docker management).
- `Dockerfile` for a consistent Go development environment with necessary tools (including `staticcheck`).
- Linting configured with `staticcheck`.
- Basic `HelloWorld` example function with a corresponding unit test.
- Ready for further development of the Chinese Checkers game logic.

## 🚀 Getting Started

Clone the repository, then:

### Tidy Go Modules:

Ensure dependencies are clean (optional, as Docker build might handle it, but good practice).

```bash
make tidy
```

### Build the Docker Image:

This creates the development environment container image.

```bash
make docker-build
```

### Run the Development Container:

This starts the container in detached mode and mounts your local code into `/app` inside the container.

```bash
make docker-run
```

Your development container is now running in the background.

### Accessing the Container Shell:

If you need direct shell access inside the container:

```bash
docker exec -it chinese-checkers-dev-instance bash
# Now you are inside the container at /app
# You can run make commands directly (run `make help` for more informations)
# Type 'exit' to leave the container shell
```

### Stopping the Container:

When you're done developing for the session:

```bash
make docker-stop
```

## Makefile Targets

The `Makefile` provides several commands for convenience. Run `make help` to see all available targets. Key targets include:

- `help`: Display available Make targets and their descriptions.
- `tidy`: Tidy `go.mod` and `go.sum` files (runs locally).
- `build`: Build the Go binary (runs locally).
- `run`: Build and run the application (runs locally).
- `test`: Run Go tests (runs locally).
- `format`: Format Go code using `go fmt` (runs locally).
- `lint`: Run `staticcheck` linter (requires local installation).
- `clean`: Remove the built binary.
- ***
- `docker-build`: Build the Docker development image.
- `docker-run`: Start the development Docker container in the background.
- `docker-stop`: Stop the running development Docker container.
- `docker-lint`: Run the `staticcheck` linter inside the Docker container.
- `docker-exec TARGET=<target>`: Execute any `make` target (e.g., `test`, `build`) inside the Docker container.

## 🧱 Project Structure

```bash
cmd/chinese-checkers/ # Main entry point
internal/game/        # Core game logic
bin/                  # Binary file
```

## 🛠️ Requirements

Before you begin, ensure you have the following installed:

- **Go:** Version 1.24 or later (check `go.mod` for the exact version).
- **Docker:** To build and run the development container. [Install Docker](https://docs.docker.com/get-docker/)
- **Make:** To use the Makefile automation. (Usually pre-installed on Linux/macOS, may need installation on Windows).
- **(Optional)** `staticcheck`: If you want to run `make lint-local` directly on your host machine. [Install staticcheck](https://staticcheck.io/docs/getting-started/)
