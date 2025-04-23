package game

import (
	"os"
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

func TestMovePawnInDefaultBoard(t *testing.T) {
	expected := &BoardState{
		Board: [][]Cell{
			{1, 1, 1, 0, 0},
			{0, 1, 0, 0, 0},
			{1, 0, 1, 0, 2},
			{0, 0, 0, 2, 2},
			{0, 0, 2, 2, 2},
		},
		CurrentPlayer: 1,
	}

	board := NewDefaultBoard()
	err := board.MovePawn("b1,c1,c2,c3")
	assert.Nil(t, err)
	assert.Equal(t, expected, board)
}

func TestMovePawnInOngoingGameBoard(t *testing.T) {
	expected := &BoardState{
		Board: [][]Cell{
			{0, 1, 1, 0, 0},
			{0, 1, 0, 0, 0},
			{1, 0, 0, 2, 0},
			{1, 1, 2, 2, 2},
			{0, 0, 0, 2, 2},
		},
		CurrentPlayer: 2,
	}

	board, err := NewBoardFromStateFile(ongoingGameStateTestPath)
	assert.Nil(t, err)
	err = board.MovePawn("c2,d2")
	assert.Nil(t, err)
	assert.Equal(t, expected, board)
}

func TestMovePawnWithNoPawnOnStartPosition(t *testing.T) {
	board := NewDefaultBoard()
	err := board.MovePawn("e1,c1,c2,c3")
	assert.Equal(t, "there is no pawn on e1", err.Error(), "should return an error with no pawn on start position")
	assert.Equal(t, &DefaultBoard, board, "should be unchanged")
}

func TestMovePawnWithAPawnOnEndPosition(t *testing.T) {
	board := NewDefaultBoard()
	err := board.MovePawn("a1,b1")
	assert.Equal(t, "there already is a pawn on b1", err.Error(), "should return an error with a pawn on end position")
	assert.Equal(t, &DefaultBoard, board, "should be unchanged")
}

func TestMovePawnErrorWithOnlyOneCell(t *testing.T) {
	board := NewDefaultBoard()
	err := board.MovePawn("a1")
	assert.Equal(t, "you must provide at least two cells in a move", err.Error(), "should be a too small move list error")
	assert.Equal(t, &DefaultBoard, board, "should be unchanged")
}

func TestMovePawnErrorWithInvalidFormat(t *testing.T) {
	board := NewDefaultBoard()
	err := board.MovePawn("a,b2")
	assert.Equal(t, "invalid cell format 'a'", err.Error(), "should be an invalid format error")
	assert.Equal(t, &DefaultBoard, board, "should be unchanged")
}

func TestMovePawnWithInvalidPosition(t *testing.T) {
	board := NewDefaultBoard()
	err := board.MovePawn("c1,d1,e1,f1")
	assert.Equal(t, "f1 is not a valid cell", err.Error(), "should be an invalid cell error")
	assert.Equal(t, &DefaultBoard, board, "should be unchanged")

	err = board.MovePawn("91,d1,e1")
	assert.Equal(t, "91 is not a valid cell", err.Error(), "should be an invalid cell error")
	assert.Equal(t, &DefaultBoard, board, "should be unchanged")
}

func TestSaveBoardState(t *testing.T) {
	testFilePath := "testFile.json"

	board := NewDefaultBoard()
	assert.Nil(t, board.SaveState(testFilePath))
	loadedBoard, err := NewBoardFromStateFile(testFilePath)
	assert.Nil(t, err)
	assert.Equal(t, &DefaultBoard, loadedBoard, "saved board must be the default board")

	os.Remove(testFilePath)
}
