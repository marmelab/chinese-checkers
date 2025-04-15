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

1.  **Build the Docker Image:**
    This creates the container image with the Go environment and tools. Run this once initially, or when the `Dockerfile` changes.

    ```bash
    make docker-build
    ```

2.  **Run the Development Container:**
    This starts the container in the background and mounts your local code into `/app` inside it. **This container must be running** for most `make` commands below to work.

    ```bash
    make docker-run
    ```

3.  **Tidy Go Modules (via Docker):**
    Ensures Go module files (`go.mod`, `go.sum`) are consistent.

    ```bash
    make tidy
    ```

4.  **Using Make Commands (via Docker):**
    With the container running, use standard `make` commands directly from your terminal. They automatically execute _inside_ the Docker container:

    ```bash
    make test      # Runs 'go test' inside the container
    make lint      # Runs 'staticcheck' inside the container
    make build     # Runs 'go build' inside the container
    make format    # Runs 'go fmt' inside the container
    make run       # Builds and runs the app inside the container
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

6.  **Stopping the Container:**
    When you're done developing for the session:
    ```bash
    make docker-stop
    ```

## 🎯 Makefile Targets

The `Makefile` provides several commands for convenience, executed via Docker. Run `make help` to see all available targets. Key targets include:

- `help`: Display this help screen.
- `build`: Build the Go binary (inside Docker).
- `run`: Build and run the application (inside Docker).
- `test`: Run Go tests (inside Docker).
- `format`: Format Go code using `go fmt` (inside Docker).
- `lint`: Run `staticcheck` linter (inside Docker).
- `tidy`: Tidy `go.mod` and `go.sum` files (inside Docker).
- `clean`: Remove the built binary (via Docker).

---

- `docker-build`: Build the Docker development image.
- `docker-run`: Start the development Docker container (Required!).
- `docker-stop`: Stop the running development Docker container.
- `docker-exec CMD="..."`: Execute an arbitrary command inside the container.

## 🧱 Project Structure (main folders)

```bash
cmd/chinese-checkers/ # Main entry point
internal/game/        # Core game logic
bin/                  # Binary file
```
