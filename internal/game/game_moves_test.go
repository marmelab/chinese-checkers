package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindBestMove(t *testing.T) {
	{
		board := NewDefaultBoard7()

		bestMove := board.FindBestMove(Green)
		assert.Equal(t, []CellIdentifier{{0, 3}, {0, 4}}, bestMove)
	}

	{
		board := NewDefaultBoard7()
		board.Board = [][]Cell{
			{1, 1, 1, 0, 2, 0, 0},
			{1, 1, 1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 2, 2},
			{0, 0, 0, 0, 2, 2, 2},
			{0, 0, 0, 2, 2, 2, 2},
		}

		{
			bestMove := board.FindBestMove(Green)
			assert.Equal(t, []CellIdentifier{{2, 6}, {3, 6}}, bestMove)
		}
		{
			bestMove := board.FindBestMove(Red)
			assert.Equal(t, []CellIdentifier{{0, 4}, {0, 3}}, bestMove)
		}
	}
}
