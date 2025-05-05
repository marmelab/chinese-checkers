package main

import (
	"github.com/labstack/echo/v4"
	"github.com/marmelab/chinese-checkers/internal/game"
	"io"
	"net/http"
)

// Error response structure.
type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleMove(c echo.Context) error {
	// Read the full body.
	body, err := io.ReadAll(c.Request().Body)

	// Cannot read the body, return an internal error.
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "internal server error",
		})
	}

	// Get the provided move path of the pawn.
	path := c.QueryParam("path")
	if len(path) == 0 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "missing move path",
		})
	}

	// Initialize a board state from the request body.
	board, err := game.NewBoardFromState(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
	}

	// Try to move the pawn as requested.
	err = board.MovePawn(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
	}

	// Return the updated board.
	return c.JSON(http.StatusOK, board)
}

func main() {
	// Initialize the API.
	e := echo.New()

	e.POST("/move", HandleMove)

	// Start the API server.
	e.Logger.Fatal(e.Start(":3003"))
}
