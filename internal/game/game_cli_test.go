package game

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	board, err := LoadBoard(ongoingGameStateTestPath)

	var output bytes.Buffer
	board.Print(&output)

	assert.Nil(t, err)
	assert.Equal(t, expected, output.String(), "should have printed an ongoing game board")
}
