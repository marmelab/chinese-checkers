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

2.  **Start the Development Container:**
    This starts the container in the background and run the application. **This container must be running** for most `make` commands below to work.

    ```bash
    make docker-run
    ```

3.  **Test:**
    Runs all Go tests inside the container.

    ```bash
    make test
    ```

4.  **Stop the Development Container:**
    When you're done developing for the session:

    ```bash
    make stop
    ```

5.  **Using Make Commands:**
    With the container running, use standard `make` commands directly from your terminal. They automatically execute _inside_ the Docker container:

    ```bash
    make help        # Show help message with available make targets
    make deps        # Tidy `go.mod` and `go.sum` files inside the container
    make lint        # Runs 'staticcheck' inside the container
    make build       # Runs 'go build' inside the container
    make run         # Builds and runs the app inside the container
    make clean       # Remove the built binary (inside Docker).
    make docker-exec # Run a command inside the docker container - Example: make docker-exec CMD="ls -l"
    ```

    _(See Makefile Targets below or run `make help` for all commands)_

6.  **Accessing the Container Shell (Optional):**
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
