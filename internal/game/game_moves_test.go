package game

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFindBestMove(t *testing.T) {
	{
		board := NewDefaultBoard7()
		board.CurrentPlayer = Green
		bestMove := board.FindBestMove()
		assert.Equal(t, []CellIdentifier{{2, 0}, {4, 0}}, bestMove)
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
			board.CurrentPlayer = Green
			bestMove := board.FindBestMove()
			assert.Equal(t, []CellIdentifier{{2, 6}, {3, 6}}, bestMove)
		}
		{
			board.CurrentPlayer = Red
			bestMove := board.FindBestMove()
			assert.Equal(t, []CellIdentifier{{0, 4}, {0, 3}}, bestMove)
		}
	}

	{
		board := NewDefaultBoard7()
		board.Board = [][]Cell{
			{1, 1, 1, 1, 0, 0, 0},
			{1, 1, 0, 1, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 2, 0, 0, 2},
			{0, 0, 1, 0, 2, 0, 2},
			{0, 0, 0, 0, 2, 2, 2},
			{0, 0, 0, 0, 2, 2, 2},
		}

		{
			board.CurrentPlayer = Green
			bestMove := board.FindBestMove()
			assert.Equal(t, []CellIdentifier{{0, 3}, {2, 3}, {4, 3}, {4, 5}}, bestMove)
		}
		{
			board.CurrentPlayer = Red
			bestMove := board.FindBestMove()
			assert.Equal(t, []CellIdentifier{{6, 5}, {4, 5}, {4, 3}, {2, 3}}, bestMove)
		}
	}
}

func TestFindPathsSimpleMove(t *testing.T) {
	board := NewDefaultBoard7()
	board.Board = [][]Cell{
		{1, 1, 1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 1, 0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 0, 2, 2},
		{0, 0, 0, 0, 2, 2, 2},
		{0, 0, 0, 2, 2, 2, 2},
	}
	board.CurrentPlayer = Green

	paths := board.FindValidMovesFrom(CellIdentifier{2, 3}, []CellIdentifier{}, true)

	assert.Equal(t, PathsTree{
		Cell: CellIdentifier{2, 3},
		Move: []CellIdentifier{CellIdentifier{2, 3}},
		Paths: []PathsTree{
			{
				Cell:  CellIdentifier{2, 4},
				Move:  []CellIdentifier{CellIdentifier{2, 3}, CellIdentifier{2, 4}},
				Paths: []PathsTree{},
			},
			{
				Cell:  CellIdentifier{2, 2},
				Move:  []CellIdentifier{CellIdentifier{2, 3}, CellIdentifier{2, 2}},
				Paths: []PathsTree{},
			},
			{
				Cell:  CellIdentifier{3, 3},
				Move:  []CellIdentifier{CellIdentifier{2, 3}, CellIdentifier{3, 3}},
				Paths: []PathsTree{},
			},
			{
				Cell:  CellIdentifier{1, 3},
				Move:  []CellIdentifier{CellIdentifier{2, 3}, CellIdentifier{1, 3}},
				Paths: []PathsTree{},
			},
		},
	}, paths, "should have 4 simple moves allowed")
}

func TestFindPathsSimpleJump(t *testing.T) {
	board := NewDefaultBoard7()
	board.Board = [][]Cell{
		{1, 1, 1, 0, 0, 0, 0},
		{1, 1, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 0},
		{1, 0, 0, 2, 0, 0, 2},
		{0, 0, 0, 2, 0, 2, 2},
		{0, 0, 0, 0, 0, 2, 2},
		{0, 0, 0, 2, 0, 2, 2},
	}
	board.CurrentPlayer = Green

	paths := board.FindValidMovesFrom(CellIdentifier{2, 3}, []CellIdentifier{}, true)

	assert.Equal(t, PathsTree{
		Cell: CellIdentifier{2, 3},
		Move: []CellIdentifier{CellIdentifier{2, 3}},
		Paths: []PathsTree{
			{
				Cell:  CellIdentifier{2, 4},
				Move:  []CellIdentifier{CellIdentifier{2, 3}, CellIdentifier{2, 4}},
				Paths: []PathsTree{},
			},
			{
				Cell:  CellIdentifier{0, 3},
				Move:  []CellIdentifier{CellIdentifier{2, 3}, CellIdentifier{0, 3}},
				Paths: []PathsTree{},
			},
		},
	}, paths, "should have one simple move and one jump allowed")
}

func TestFindPathsChainedJump(t *testing.T) {
	board := NewDefaultBoard7()
	board.Board = [][]Cell{
		{1, 1, 0, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 2, 0},
		{0, 1, 1, 1, 0, 0, 0},
		{1, 0, 0, 2, 0, 0, 0},
		{0, 0, 0, 2, 0, 2, 2},
		{0, 0, 0, 0, 0, 2, 2},
		{0, 0, 0, 2, 0, 2, 2},
	}
	board.CurrentPlayer = Green

	paths := board.FindValidMovesFrom(CellIdentifier{2, 3}, []CellIdentifier{}, true)

	assert.Equal(t, PathsTree{
		Cell: CellIdentifier{2, 3},
		Move: []CellIdentifier{CellIdentifier{2, 3}},
		Paths: []PathsTree{
			{
				Cell:  CellIdentifier{2, 4},
				Move:  []CellIdentifier{CellIdentifier{2, 3}, CellIdentifier{2, 4}},
				Paths: []PathsTree{},
			},
			{
				Cell: CellIdentifier{0, 3},
				Move: []CellIdentifier{CellIdentifier{2, 3}, CellIdentifier{0, 3}},
				Paths: []PathsTree{
					{
						Cell: CellIdentifier{0, 5},
						Move: []CellIdentifier{CellIdentifier{2, 3}, CellIdentifier{0, 3}, CellIdentifier{0, 5}},
						Paths: []PathsTree{
							{
								Cell:  CellIdentifier{2, 5},
								Move:  []CellIdentifier{CellIdentifier{2, 3}, CellIdentifier{0, 3}, CellIdentifier{0, 5}, CellIdentifier{2, 5}},
								Paths: []PathsTree{},
							},
						},
					},
				},
			},
		},
	}, paths, "should have one simple move and chained jumps to c6")
}

func TestBestMovesWin(t *testing.T) {
	board := NewDefaultBoard7()
	board.Board = [][]Cell{
		{2, 2, 2, 2, 0, 0, 0},
		{2, 2, 2, 0, 0, 1, 0},
		{0, 2, 2, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1, 1},
		{2, 0, 0, 0, 1, 0, 1},
		{0, 0, 0, 1, 1, 1, 1},
	}
	board.CurrentPlayer = Green

	assert.Nil(t, board.movePawn(board.FindBestMove()))
	assert.Nil(t, board.movePawn(board.FindBestMove()))
	assert.Nil(t, board.movePawn(board.FindBestMove()))
	assert.Nil(t, board.movePawn(board.FindBestMove()))
	board.Print(os.Stdout)
	assert.Nil(t, board.movePawn(board.FindBestMove()))
	board.Print(os.Stdout)

	assert.Equal(t, Green, board.GetWinner())
}
