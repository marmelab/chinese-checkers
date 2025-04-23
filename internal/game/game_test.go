package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const ongoingGameTestStatePath = "../../tests/states/ongoing-game.json"

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
	board, err := LoadBoard(ongoingGameTestStatePath)

	assert.Nil(t, err)
	assert.Equal(t, expected, board, "should be an ongoing game board")
}

func TestInitBoardWithoutFilePath(t *testing.T) {
	expected := &DefaultBoard
	board, err := InitBoard(nil)
	assert.Nil(t, err)
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
	path := ongoingGameTestStatePath
	board, err := InitBoard(&path)

	assert.Nil(t, err)
	assert.Equal(t, expected, board, "should be an ongoing game board")
}
