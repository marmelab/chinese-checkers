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
	DefaultBoard5.Print(&output)

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

func TestOngoing7x7GameBoardPrinting(t *testing.T) {
	expected := `     1    2    3    4    5    6    7   
 . +----+----+----+----+----+----+----+
 a |    | 🟢 | 🟢 |    |    |    |    |
 . +----+----+----+----+----+----+----+
 b | 🟢 | 🟢 | 🟢 | 🟢 |    |    |    |
 . +----+----+----+----+----+----+----+
 c | 🟢 | 🟢 |    |    |    |    |    |
 . +----+----+----+----+----+----+----+
 d | 🟢 |    | 🟢 |    |    |    | 🔴 |
 . +----+----+----+----+----+----+----+
 e |    |    |    |    |    | 🔴 | 🔴 |
 . +----+----+----+----+----+----+----+
 f |    |    | 🔴 | 🔴 | 🔴 | 🔴 |    |
 . +----+----+----+----+----+----+----+
 g |    |    |    | 🔴 |    | 🔴 | 🔴 |
 . +----+----+----+----+----+----+----+
`

	board, err := NewBoardFromStateFile(ongoing7x7GameStateTestPath)

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

	board := NewDefaultBoard5()
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
		expected := "     " + coloring.For("Green").Bold().Green().String() + ": 0/6, " + coloring.For("Red").Bold().Red().String() + ": 0/6\n"

		board := NewDefaultBoard5()

		var output bytes.Buffer
		board.PrintScore(&output)

		assert.Equal(t, expected, output.String(), "should have printed an accurate scoreboard")
	}

	{
		expected := "     " + coloring.For("Green").Bold().Green().String() + ": 3/6, " + coloring.For("Red").Bold().Red().String() + ": 3/6\n"

		board := NewDefaultBoard5()
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

	{
		expected := "     " + coloring.For("Green").Bold().Green().String() + ": 3/10, " + coloring.For("Red").Bold().Red().String() + ": 3/10\n"

		board := NewDefaultBoard7()
		board.Board = [][]Cell{
			{1, 1, 2, 0, 0, 0, 1},
			{1, 2, 0, 0, 0, 0, 0},
			{2, 0, 0, 0, 1, 2, 2},
			{0, 0, 0, 1, 2, 0, 0},
			{0, 0, 1, 2, 2, 0, 0},
			{0, 0, 0, 0, 2, 0, 1},
			{0, 0, 0, 0, 2, 1, 1},
		}

		var output bytes.Buffer
		board.PrintScore(&output)

		assert.Equal(t, expected, output.String(), "should have printed an accurate scoreboard")
	}
}

func TestBoardPrintLastMove(t *testing.T) {
	board := NewDefaultBoard7()
	board.CurrentPlayer = Green

	err := board.MovePawn("b2,b4")
	assert.Nil(t, err, "should move the pawn without error")

	var output bytes.Buffer
	board.PrintLastMove(&output)
	assert.Equal(t, "Last move from "+coloring.For("Green").Bold().Green().String()+" player: b2, b4\n", output.String(), "should have an accurate last move")
}

func TestBoardPrintingWithLastMove(t *testing.T) {
	board := NewDefaultBoard7()
	board.CurrentPlayer = Green

	err := board.MovePawn("b2,b4")
	assert.Nil(t, err, "should move the pawn without error")

	expected := `     1    2    3    4    5    6    7   
 . +----+----+----+----+----+----+----+
 a | 🟢 | 🟢 | 🟢 | 🟢 |    |    |    |
 . +----+----+----+----+----+----+----+
 b | 🟢 |` + coloring.For("  ⬤ ").Rgb(0, 70, 0).String() + `| 🟢 | 🟢 |    |    |    |
 . +----+----+----+----+----+----+----+
 c | 🟢 | 🟢 |    |    |    |    |    |
 . +----+----+----+----+----+----+----+
 d | 🟢 |    |    |    |    |    | 🔴 |
 . +----+----+----+----+----+----+----+
 e |    |    |    |    |    | 🔴 | 🔴 |
 . +----+----+----+----+----+----+----+
 f |    |    |    |    | 🔴 | 🔴 | 🔴 |
 . +----+----+----+----+----+----+----+
 g |    |    |    | 🔴 | 🔴 | 🔴 | 🔴 |
 . +----+----+----+----+----+----+----+
`

	var output bytes.Buffer
	board.Print(&output)
	assert.Equal(t, expected, output.String(), "should have printed an default game board with one last move")
}

func TestBoardPrintingWithJumpAsLastMove(t *testing.T) {
	board, err := NewBoardFromStateFile(ongoing7x7GameStateTestPath)
	assert.Nil(t, err, "should load the board without error")
	board.CurrentPlayer = Green

	err = board.MovePawn("b2,d2,d4")
	assert.Nil(t, err, "should move the pawn without error")

	expected := `     1    2    3    4    5    6    7   
 . +----+----+----+----+----+----+----+
 a |    | 🟢 | 🟢 |    |    |    |    |
 . +----+----+----+----+----+----+----+
 b | 🟢 |` + coloring.For("  ⬤ ").Rgb(0, 70, 0).String() + `| 🟢 | 🟢 |    |    |    |
 . +----+----+----+----+----+----+----+
 c | 🟢 | 🟢 |    |    |    |    |    |
 . +----+----+----+----+----+----+----+
 d | 🟢 |` + coloring.For("  ⬤ ").Rgb(0, 70, 0).String() + `| 🟢 | 🟢 |    |    | 🔴 |
 . +----+----+----+----+----+----+----+
 e |    |    |    |    |    | 🔴 | 🔴 |
 . +----+----+----+----+----+----+----+
 f |    |    | 🔴 | 🔴 | 🔴 | 🔴 |    |
 . +----+----+----+----+----+----+----+
 g |    |    |    | 🔴 |    | 🔴 | 🔴 |
 . +----+----+----+----+----+----+----+
`

	var output bytes.Buffer
	board.Print(&output)
	assert.Equal(t, expected, output.String(), "should have printed an default game board with one last move")
}
