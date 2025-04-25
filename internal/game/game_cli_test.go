package game

import (
	"bytes"
	"testing"

	"github.com/go-color-term/go-color-term/coloring"
	"github.com/stretchr/testify/assert"
)

func TestBoardPrinting(t *testing.T) {
	expected := `     1    2    3    4    5  
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
	expected := `     1    2    3    4    5  
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

	board, err := NewBoardFromStateFile(ongoingGameStateTestPath)

	var output bytes.Buffer
	board.Print(&output)

	assert.Nil(t, err)
	assert.Equal(t, expected, output.String(), "should have printed an ongoing game board")
}

func TestBoardPrintingWithPawnsInTargetArea(t *testing.T) {
	expected := `     1    2    3    4    5  
 . +----+----+----+----+----+
 a | 🟢 | 🟢 |` + coloring.For(" 🔴 ").Background().Rgb(84, 0, 0).String() + `|    |    |
 . +----+----+----+----+----+
 b | 🟢 |` + coloring.For(" 🔴 ").Background().Rgb(84, 0, 0).String() + `|    |    |    |
 . +----+----+----+----+----+
 c |` + coloring.For(" 🔴 ").Background().Rgb(84, 0, 0).String() + `|    |    |    |` + coloring.For(" 🟢 ").Background().Rgb(0, 70, 0).String() + `|
 . +----+----+----+----+----+
 d |    |    |    |` + coloring.For(" 🟢 ").Background().Rgb(0, 70, 0).String() + `| 🔴 |
 . +----+----+----+----+----+
 e |    |    |` + coloring.For(" 🟢 ").Background().Rgb(0, 70, 0).String() + `| 🔴 | 🔴 |
 . +----+----+----+----+----+
`

	board := NewDefaultBoard()
	board.Board = [][]Cell{
		{1, 1, 2, 0, 0},
		{1, 2, 0, 0, 0},
		{2, 0, 0, 0, 1},
		{0, 0, 0, 1, 2},
		{0, 0, 1, 2, 2},
	}

	var output bytes.Buffer
	board.Print(&output)

	assert.Equal(t, expected, output.String(), "should have printed a default board")
}

func TestBoardScorePrinting(t *testing.T) {
	{
		expected := coloring.For("Green").Bold().Green().String() + ": 0/6, " + coloring.For("Red").Bold().Red().String() + ": 0/6\n"

		board := NewDefaultBoard()

		var output bytes.Buffer
		board.PrintScore(&output)

		assert.Equal(t, expected, output.String(), "should have printed an accurate scoreboard")
	}

	{
		expected := coloring.For("Green").Bold().Green().String() + ": 3/6, " + coloring.For("Red").Bold().Red().String() + ": 3/6\n"

		board := NewDefaultBoard()
		board.Board = [][]Cell{
			{1, 1, 2, 0, 0},
			{1, 2, 0, 0, 0},
			{2, 0, 0, 0, 1},
			{0, 0, 0, 1, 2},
			{0, 0, 1, 2, 2},
		}

		var output bytes.Buffer
		board.PrintScore(&output)

		assert.Equal(t, expected, output.String(), "should have printed an accurate scoreboard")
	}
}
