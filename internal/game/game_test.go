package game

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ongoingGameStateTestPath = "../../tests/states/ongoing-game.json"
const invalidGameStateTestPath = "../../tests/states/invalid-board.json"

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
	board, err := LoadBoard(ongoingGameStateTestPath)

	assert.Nil(t, err)
	assert.Equal(t, expected, board, "should be an ongoing game board")
}

func TestLoadInvalidBoard(t *testing.T) {
	board, err := LoadBoard(invalidGameStateTestPath)

	assert.Nil(t, board)
	assert.Equal(t, "invalid game state, please provide a valid game state", err.Error(), "should return an invalid game state error")
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
	path := ongoingGameStateTestPath
	board, err := InitBoard(&path)

	assert.Nil(t, err)
	assert.Equal(t, expected, board, "should be an ongoing game board")
}

func TestBoardPrinting(t *testing.T) {
	expected := `    1    2    3    4    5  
. +----+----+----+----+----+
a | 🟢 | 🟢 | 🟢 |    |    |
. +----+----+----+----+----+
b | 🟢 | 🟢 |    |    |    |
. +----+----+----+----+----+
c | 🟢 |    |    |    | 🔴 |
. +----+----+----+----+----+
d |    |    |    | 🔴 | 🔴 |
. +----+----+----+----+----+
e |    |    | 🔴 | 🔴 | 🔴 |
. +----+----+----+----+----+
`

	var output bytes.Buffer
	DefaultBoard.Print(&output)

	assert.Equal(t, expected, output.String(), "should have printed a default board")
}

func TestOngoingGameBoardPrinting(t *testing.T) {
	expected := `    1    2    3    4    5  
. +----+----+----+----+----+
a |    | 🟢 | 🟢 |    |    |
. +----+----+----+----+----+
b |    | 🟢 |    |    |    |
. +----+----+----+----+----+
c | 🟢 | 🟢 |    | 🔴 |    |
. +----+----+----+----+----+
d | 🟢 |    | 🔴 | 🔴 | 🔴 |
. +----+----+----+----+----+
e |    |    |    | 🔴 | 🔴 |
. +----+----+----+----+----+
`

	path := ongoingGameStateTestPath
	board, err := InitBoard(&path)

	var output bytes.Buffer
	board.Print(&output)

	assert.Nil(t, err)
	assert.Equal(t, expected, output.String(), "should have printed an ongoing game board")
}
