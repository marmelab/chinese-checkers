package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const ongoingGameStateTestPath = "../../tests/states/ongoing-game.json"
const invalidGameStateTestPath = "../../tests/states/invalid-board.json"
const invalidPawnsCountTestPath = "../../tests/states/invalid-pawns-count.json"

func TestLoadBoard(t *testing.T) {
	expected := &BoardState{
		Board: [][]Cell{
			{0, 1, 1, 0, 0},
			{0, 1, 0, 0, 0},
			{1, 1, 0, 2, 0},
			{1, 0, 2, 2, 2},
			{0, 0, 0, 2, 2},
		},
		CurrentPlayer: 2,
	}
	board, err := NewBoardFromStateFile(ongoingGameStateTestPath)

	assert.Nil(t, err)
	assert.Equal(t, expected, board, "should be an ongoing game board")
}

func TestLoadInvalidBoard(t *testing.T) {
	board, err := NewBoardFromStateFile(invalidGameStateTestPath)

	assert.Nil(t, board)
	assert.Equal(t, "invalid game state, please provide a valid game state", err.Error(), "should return an invalid game state error")
}

func TestLoadBoardWithInvalidPawnsCount(t *testing.T) {
	board, err := NewBoardFromStateFile(invalidPawnsCountTestPath)

	assert.Nil(t, board)
	assert.Equal(t, "invalid game state, please provide a valid game state", err.Error(), "should return an invalid game state error")
}

func TestNewDefaultBoard(t *testing.T) {
	expected := &DefaultBoard
	board := NewDefaultBoard()
	assert.Equal(t, expected, board, "should be the default board")
}

func TestInitBoardWithFilePath(t *testing.T) {
	expected := &BoardState{
		Board: [][]Cell{
			{0, 1, 1, 0, 0},
			{0, 1, 0, 0, 0},
			{1, 1, 0, 2, 0},
			{1, 0, 2, 2, 2},
			{0, 0, 0, 2, 2},
		},
		CurrentPlayer: 2,
	}
	board, err := NewBoardFromStateFile(ongoingGameStateTestPath)

	assert.Nil(t, err)
	assert.Equal(t, expected, board, "should be an ongoing game board")
}

func TestBoardCloning(t *testing.T) {
	// Test to clone the default dashboard.
	clonedBoard := DefaultBoard.Clone()
	assert.Equal(t, &DefaultBoard, clonedBoard, "should be the same board")
	assert.NotSame(t, &DefaultBoard, clonedBoard, "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard.Board, &clonedBoard.Board, "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard.Board[0], &clonedBoard.Board[0], "shouldn't be the same pointer")

	// Test to clone a loaded dashboard.
	board, err := NewBoardFromStateFile(ongoingGameStateTestPath)
	assert.Nil(t, err)
	clonedBoard = board.Clone()
	assert.Equal(t, board, clonedBoard, "should be the same board")
	assert.NotSame(t, &DefaultBoard, clonedBoard, "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard.Board, &clonedBoard.Board, "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard.Board[0], &clonedBoard.Board[0], "shouldn't be the same pointer")
}
