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

func TestWinnerApiNoWinner(t *testing.T) {
	// Build the request.
	request := httptest.NewRequest(http.MethodPost, "/winner", strings.NewReader(`
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
	err := HandleWinner(context)
	assert.Nil(t, err)

	// Parse the response.
	responseBody, err := io.ReadAll(response.Body)
	assert.Nil(t, err, "should read the response body without error")
	var winner int8
	err = json.Unmarshal(responseBody, &winner)
	assert.Nil(t, err, "should parse the winner without error")

	assert.Equal(t, int8(0), winner, "should have no winner")
}

func TestWinnerApiGreenWinner(t *testing.T) {
	// Build the request.
	request := httptest.NewRequest(http.MethodPost, "/winner", strings.NewReader(`
{
  "board": [
    [0, 2, 2, 2, 0, 0, 0],
    [2, 2, 2, 2, 0, 0, 0],
    [2, 2, 0, 0, 0, 0, 0],
    [2, 0, 0, 0, 0, 0, 1],
    [0, 0, 0, 0, 0, 1, 1],
    [0, 0, 0, 0, 1, 1, 1],
    [0, 0, 0, 1, 1, 1, 1]
  ],
  "currentPlayer": 2
}
`))
	response := httptest.NewRecorder()

	// Run the request.
	e := echo.New()
	context := e.NewContext(request, response)
	err := HandleWinner(context)
	assert.Nil(t, err)

	// Parse the response.
	responseBody, err := io.ReadAll(response.Body)
	assert.Nil(t, err, "should read the response body without error")
	var winner int8
	err = json.Unmarshal(responseBody, &winner)
	assert.Nil(t, err, "should parse the winner without error")

	assert.Equal(t, int8(1), winner, "should have green as a winner")
}

func TestEvaluationApi(t *testing.T) {
	// Build the request.
	request := httptest.NewRequest(http.MethodPost, "/evaluate", strings.NewReader(`
{
  "board": [
    [0, 1, 1, 0, 0, 0, 0],
    [1, 1, 1, 2, 0, 0, 0],
    [1, 1, 0, 0, 0, 0, 0],
    [1, 0, 2, 0, 0, 0, 2],
    [0, 0, 0, 0, 0, 2, 2],
    [0, 0, 1, 0, 2, 2, 0],
    [0, 0, 0, 2, 1, 2, 2]
  ],
  "currentPlayer": 2
}
`))
	response := httptest.NewRecorder()

	// Run the request.
	e := echo.New()
	context := e.NewContext(request, response)
	err := HandleEvaluate(context)
	assert.Nil(t, err)

	// Parse the response.
	responseBody, err := io.ReadAll(response.Body)
	assert.Nil(t, err, "should read the response body without error")

	evaluation := struct {
		Evaluation game.Evaluation
	}{Evaluation: game.Evaluation{}}
	err = json.Unmarshal(responseBody, &evaluation)
	assert.Nil(t, err, "should parse the evaluation without error")

	assert.Equal(t, struct {
		Evaluation game.Evaluation
	}{Evaluation: game.Evaluation{
		GreenScore: 51,
		RedScore:   49,
	}}, evaluation, "should have correct evaluation scores")
}

func TestHintApi(t *testing.T) {
	// Build the request.
	request := httptest.NewRequest(http.MethodPost, "/hint", strings.NewReader(`
{
  "board": [
    [0, 1, 1, 0, 0, 0, 0],
    [1, 1, 1, 2, 0, 0, 0],
    [1, 1, 0, 0, 0, 0, 0],
    [1, 0, 2, 0, 0, 0, 2],
    [0, 0, 0, 0, 0, 2, 2],
    [0, 0, 1, 0, 2, 2, 0],
    [0, 0, 0, 2, 1, 2, 2]
  ],
  "currentPlayer": 2
}
`))
	response := httptest.NewRecorder()

	// Run the request.
	e := echo.New()
	context := e.NewContext(request, response)
	err := HandleHint(context)
	assert.Nil(t, err)

	// Parse the response.
	responseBody, err := io.ReadAll(response.Body)
	assert.Nil(t, err, "should read the response body without error")

	hint := struct {
		Move []game.CellIdentifier
	}{}
	err = json.Unmarshal(responseBody, &hint)
	assert.Nil(t, err, "should parse the hint without error")

	assert.Equal(t, struct {
		Move []game.CellIdentifier
	}{Move: []game.CellIdentifier{{1, 3}, {0, 3}}}, hint, "should have correct hint")
}
