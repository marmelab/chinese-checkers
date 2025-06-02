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

func HandleWinner(c echo.Context) error {
	board, err := parseGameBoard(c)
	if err != nil {
		return err
	}

	// Return the updated board.
	return c.JSON(http.StatusOK, board.GetWinner())
}

func HandleEvaluate(c echo.Context) error {
	board, err := parseGameBoard(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		Evaluation game.Evaluation `json:"evaluation"`
	}{
		Evaluation: board.Evaluate(),
	})
}

func HandleHint(c echo.Context) error {
	board, err := parseGameBoard(c)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		Move []game.CellIdentifier `json:"move"`
	}{
		Move: board.FindBestMove(board.CurrentPlayer),
	})
}

func HandleValidMoves(c echo.Context) error {
	board, err := parseGameBoard(c)
	if err != nil {
		return err
	}

	// Get the provided starting cell.
	from := c.QueryParam("from")
	if len(from) == 0 {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "missing cell from which to start moving",
		})
	}
	fromCell, err := board.ParseCellIdentifier(from)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, board.FindValidMovesFrom(*fromCell, []game.CellIdentifier{}, true))
}

func parseGameBoard(c echo.Context) (*game.BoardState, error) {
	// Read the full body.
	body, err := io.ReadAll(c.Request().Body)

	// Cannot read the body, return an internal error.
	if err != nil {
		if responseErr := c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "internal server error",
		}); responseErr != nil {
			return nil, responseErr
		}

		return nil, err
	}

	// Initialize a board state from the request body.
	board, err := game.NewBoardFromState(body)
	if err != nil {
		if responseErr := c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: err.Error(),
		}); responseErr != nil {
			return nil, responseErr
		}

		return nil, err
	}

	return board, nil
}

func main() {
	// Initialize the API.
	e := echo.New()

	e.POST("/move", HandleMove)

	e.POST("/winner", HandleWinner)

	e.POST("/evaluate", HandleEvaluate)

	e.POST("/hint", HandleHint)

	e.POST("/valid-moves", HandleValidMoves)

	// Start the API server.
	e.Logger.Fatal(e.Start(":3003"))
}
