package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/marmelab/chinese-checkers/internal/game"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMoveApi(t *testing.T) {
	// Build the request.
	request := httptest.NewRequest(http.MethodPost, "/move?path=d7,c7", strings.NewReader(`
{
  "board": [
    [0, 1, 1, 0, 0, 0, 0],
    [1, 1, 1, 1, 0, 0, 0],
    [1, 1, 0, 0, 0, 0, 0],
    [1, 0, 1, 0, 0, 0, 2],
    [0, 0, 0, 0, 0, 2, 2],
    [0, 0, 2, 2, 2, 2, 0],
    [0, 0, 0, 2, 0, 2, 2]
  ],
  "currentPlayer": 2
}
`))
	response := httptest.NewRecorder()

	// Run the request.
	e := echo.New()
	context := e.NewContext(request, response)
	err := HandleMove(context)
	assert.Nil(t, err)

	// Parse the response.
	responseBody, err := io.ReadAll(response.Body)
	assert.Nil(t, err, "should read the response body without error")
	board, err := game.NewBoardFromState(responseBody)
	assert.Nil(t, err, "should parse the new board state without error")

	// Check the board.
	assert.Equal(t, [][]game.Cell{
		{0, 1, 1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 2},
		{1, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 2, 2},
		{0, 0, 2, 2, 2, 2, 0},
		{0, 0, 0, 2, 0, 2, 2},
	}, board.Board, "a red pawn should have moved from d7 to c7")
	assert.Equal(t, game.Green, board.CurrentPlayer, "the new current player should be green")
}

func TestMoveApiRequestError(t *testing.T) {
	// Build the request.
	request := httptest.NewRequest(http.MethodPost, "/move", strings.NewReader(`
{
  "board": [
    [0, 1, 1, 0, 0, 0, 0],
    [1, 1, 1, 1, 0, 0, 0],
    [1, 1, 0, 0, 0, 0, 0],
    [1, 0, 1, 0, 0, 0, 2],
    [0, 0, 0, 0, 0, 2, 2],
    [0, 0, 2, 2, 2, 2, 0],
    [0, 0, 0, 2, 0, 2, 2]
  ],
  "currentPlayer": 2
}
`))
	response := httptest.NewRecorder()

	// Run the request.
	e := echo.New()
	context := e.NewContext(request, response)
	err := HandleMove(context)
	assert.Nil(t, err)

	// Parse the response.
	responseBody, err := io.ReadAll(response.Body)
	assert.Nil(t, err, "should read the response body without error")
	// Parse the error.
	var errorResponse ErrorResponse
	err = json.Unmarshal(responseBody, &errorResponse)
	assert.Nil(t, err, "should parse the error response body without error")

	// Check the error.
	assert.Equal(t, "missing move path", errorResponse.Error, "should indicate the path is missing")
}

func TestMoveApiGameError(t *testing.T) {
	// Build the request.
	request := httptest.NewRequest(http.MethodPost, "/move?path=c7,b7", strings.NewReader(`
{
  "board": [
    [0, 1, 1, 0, 0, 0, 0],
    [1, 1, 1, 1, 0, 0, 0],
    [1, 1, 0, 0, 0, 0, 0],
    [1, 0, 1, 0, 0, 0, 2],
    [0, 0, 0, 0, 0, 2, 2],
    [0, 0, 2, 2, 2, 2, 0],
    [0, 0, 0, 2, 0, 2, 2]
  ],
  "currentPlayer": 2
}
`))
	response := httptest.NewRecorder()

	// Run the request.
	e := echo.New()
	context := e.NewContext(request, response)
	err := HandleMove(context)
	assert.Nil(t, err)

	// Parse the response.
	responseBody, err := io.ReadAll(response.Body)
	assert.Nil(t, err, "should read the response body without error")
	// Parse the error.
	var errorResponse ErrorResponse
	err = json.Unmarshal(responseBody, &errorResponse)
	assert.Nil(t, err, "should parse the error response body without error")

	// Check the error.
	assert.Equal(t, "there is no pawn on c7", errorResponse.Error, "should indicate that there is no pawn on c7")
}
