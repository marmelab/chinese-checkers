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

		// Red distance score: sqrt(0*0 + 4*4) = 4
		// Green distance score: 0

		scores := board.Evaluate()
		assert.Equal(t, 100, scores.GreenScore)
		assert.Equal(t, 0, scores.RedScore)
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

		// Red distance score: sqrt(0*0 + 4*4) + sqrt(0*0 + 5*5) = 9
		// Green distance score: sqrt(1*1 + 3*3) ~= 3

		scores := board.Evaluate()
		assert.Equal(t, 75, scores.GreenScore)
		assert.Equal(t, 25, scores.RedScore)
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

		// Red distance score: sqrt(0*0 + 4*4) + sqrt(0*0 + 5*5) = 9
		// Green distance score: sqrt(1*1 + 3*3) + sqrt(2*2 + 3*3) ~= 7

		scores := board.Evaluate()
		assert.Equal(t, 56, scores.GreenScore)
		assert.Equal(t, 44, scores.RedScore)
	}

	{
		board := NewDefaultBoard7()

		// Starting position: all players have the same chance to win.

		scores := board.Evaluate()
		assert.Equal(t, 50, scores.GreenScore)
		assert.Equal(t, 50, scores.RedScore)
	}
}
