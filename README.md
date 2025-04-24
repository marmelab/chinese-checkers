# Chinese Checkers CLI

A simple two-player Chinese Checkers game playable in the terminal, written in Go.

## Features ✨

- Standard Go project structure (`cmd/`, `internal/`).
- `Makefile` for automating common development tasks (build, test, lint, run).
- `Dockerfile` for a consistent Go development environment with necessary tools (including `staticcheck`).
- Linting configured with `staticcheck`.
- Basic `HelloWorld` example function with a corresponding unit test.
- Development workflow entirely containerized via Docker.
- Ready for further development of the Chinese Checkers game logic!

## 🛠️ Requirements

Before you begin, ensure you have the following installed:

- **Docker:** To build and run the development container. [Install Docker](https://docs.docker.com/get-docker/)
- **Make:** To use the Makefile automation. (Usually pre-installed on Linux/macOS, may need installation on Windows).

_(Note: Go and staticcheck do NOT need to be installed locally.)_

## 🚀 Getting Started

Clone the repository, then:

1.  **Install the development Docker Image:**
    This creates the container image with the Go environment and tools. Run this once initially, or when the `Dockerfile` changes.

    ```bash
    make install
    ```

2.  **Run the CLI:**
    This runs the project inside the container.

    ```bash
		# Show help.
		APP_ARGS="--help" make run
		# Run with arguments.
		APP_ARGS="--state-file tests/states/ongoing-game.json --move c1,d1" make run
    ```

3.  **Test:**
    Runs all Go tests inside the container.

    ```bash
    make test
    ```

4.  **Using Make Commands:**
    With the container running, use standard `make` commands directly from your terminal. They automatically execute _inside_ the Docker container:

    ```bash
    make build                          Build the Go binary (inside Docker).
    make check                          Run `staticcheck` and `go vet` (inside Docker).
    make clean                          Remove the built binary (inside Docker).
    make deps                           Tidy `go.mod` and `go.sum` files (inside Docker).
    make docker-exec                    Run a command inside the docker container - Example: make docker-exec CMD="ls -l"
    make install                        Build the Docker development image.
    make lint                           Run `staticcheck` linter (inside Docker).
    make run                            run the application.
    make test                           Run Go tests (inside Docker).
    make vet                            Run `go vet` (inside Docker).
    ```

    _(See Makefile Targets below or run `make help` for all commands)_

5.  **Accessing the Container Shell (Optional):**
    If you need direct shell access inside the container:

    ```bash
    docker exec -it chinese-checkers-dev-instance bash
    # Now you are inside the container at /app
    # You can run Go commands (go test...) or make commands directly here.
    # Type 'exit' to leave the container shell
    ```

## 🧱 Project Structure (main folders)

```bash
cmd/chinese-checkers/ # Main entry point
internal/game/        # Core game logic
bin/                  # Binary file
```
