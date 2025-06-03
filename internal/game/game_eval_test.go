package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameEvaluation(t *testing.T) {
	{
		board := NewDefaultBoard7()
		board.Board = [][]Cell{
			{2, 2, 2, 2, 2, 0, 0},
			{2, 2, 2, 0, 0, 0, 0},
			{2, 2, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0, 1, 1},
			{0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 1, 1, 1, 1},
		}

		chances := board.Evaluate()
		assert.Equal(t, 65, chances.Green)
		assert.Equal(t, 35, chances.Red)
	}

	{
		board := NewDefaultBoard7()
		board.Board = [][]Cell{
			{2, 2, 2, 2, 2, 2, 0},
			{2, 2, 2, 0, 0, 0, 0},
			{2, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0, 1, 1},
			{0, 0, 0, 1, 0, 1, 1},
			{0, 0, 0, 1, 1, 1, 1},
		}

		chances := board.Evaluate()
		assert.Equal(t, 71, chances.Green)
		assert.Equal(t, 29, chances.Red)
	}

	{
		board := NewDefaultBoard7()
		board.Board = [][]Cell{
			{2, 2, 2, 2, 2, 2, 0},
			{2, 2, 2, 0, 0, 0, 0},
			{2, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 1, 0, 1, 1},
			{0, 0, 0, 1, 0, 1, 1},
			{0, 0, 0, 1, 0, 1, 1},
		}

		chances := board.Evaluate()
		assert.Equal(t, 65, chances.Green)
		assert.Equal(t, 35, chances.Red)
	}

	{
		board := NewDefaultBoard7()

		// Starting position: all players have the same chance to win.

		chances := board.Evaluate()
		assert.Equal(t, 50, chances.Green)
		assert.Equal(t, 50, chances.Red)
	}
}
