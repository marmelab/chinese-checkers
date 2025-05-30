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

		// Red distance: sqrt(0*0 + 4*4) = 4
		// Green distance: 0

		chances := board.Evaluate()
		assert.Equal(t, 100, chances.Green)
		assert.Equal(t, 0, chances.Red)
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

		// Red distance: sqrt(0*0 + 4*4) + sqrt(0*0 + 5*5) = 9
		// Green distance: sqrt(1*1 + 3*3) ~= 3

		chances := board.Evaluate()
		assert.Equal(t, 75, chances.Green)
		assert.Equal(t, 25, chances.Red)
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

		// Red distance: sqrt(0*0 + 4*4) + sqrt(0*0 + 5*5) = 9
		// Green distance: sqrt(1*1 + 3*3) + sqrt(2*2 + 3*3) ~= 7

		chances := board.Evaluate()
		assert.Equal(t, 56, chances.Green)
		assert.Equal(t, 44, chances.Red)
	}

	{
		board := NewDefaultBoard7()

		// Starting position: all players have the same chance to win.

		chances := board.Evaluate()
		assert.Equal(t, 50, chances.Green)
		assert.Equal(t, 50, chances.Red)
	}
}
