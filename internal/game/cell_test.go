package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCellIdentifierParser(t *testing.T) {
	output, err := DefaultBoard5.ParseCellIdentifier("a3")
	assert.Nil(t, err)
	assert.Equal(t, &CellIdentifier{0, 2}, output, "should be the same identifier")

	output, err = DefaultBoard5.ParseCellIdentifier(" E5")
	assert.Nil(t, err)
	assert.Equal(t, &CellIdentifier{4, 4}, output, "should be the same identifier")
}

func TestInvalidCellIdentifierParsing(t *testing.T) {
	_, err := DefaultBoard5.ParseCellIdentifier("a")
	assert.Equal(t, err.Error(), "invalid cell format 'a'", "should be an invalid format error")

	_, err = DefaultBoard5.ParseCellIdentifier("h9")
	assert.Equal(t, err.Error(), "h9 is not a valid cell", "should be an invalid cell error")
}

func TestMoveListParser(t *testing.T) {
	output, err := DefaultBoard5.ParseMoveList("a1,b2,c3")
	assert.Nil(t, err)
	assert.Equal(t, []CellIdentifier{{0, 0}, {1, 1}, {2, 2}}, output)

	output, err = DefaultBoard5.ParseMoveList("b4, b3, b2, a2")
	assert.Nil(t, err)
	assert.Equal(t, []CellIdentifier{{1, 3}, {1, 2}, {1, 1}, {0, 1}}, output)
}

func TestMoveListSizeError(t *testing.T) {
	_, err := DefaultBoard5.ParseMoveList("a1")
	assert.Equal(t, "you must provide at least two cells in a move", err.Error(), "should be a too small move list error")
}

func TestMoveListFormatError(t *testing.T) {
	_, err := DefaultBoard5.ParseMoveList("a1,b,c3")
	assert.Equal(t, "invalid cell format 'b'", err.Error(), "should be an invalid format error")
}

func TestCellInMask(t *testing.T) {
	cell, err := DefaultBoard5.ParseCellIdentifier("a1")
	assert.Nil(t, err)
	assert.True(t, cell.InMask([][]Cell{
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
	assert.False(t, cell.InMask([][]Cell{
		{0, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}))
	assert.False(t, cell.InMask([][]Cell{}))
}

func TestCellInTargetAreas(t *testing.T) {
	{
		cell, err := DefaultBoard5.ParseCellIdentifier("c1")
		assert.Nil(t, err)
		assert.False(t, cell.InMask(DefaultBoard5.gameDefinition.GreenTargetAreaMask))
		assert.True(t, cell.InMask(DefaultBoard5.gameDefinition.RedTargetAreaMask))
	}

	{
		cell, err := DefaultBoard5.ParseCellIdentifier("d4")
		assert.Nil(t, err)
		assert.True(t, cell.InMask(DefaultBoard5.gameDefinition.GreenTargetAreaMask))
		assert.False(t, cell.InMask(DefaultBoard5.gameDefinition.RedTargetAreaMask))
	}

	{
		cell, err := DefaultBoard5.ParseCellIdentifier("e2")
		assert.Nil(t, err)
		assert.False(t, cell.InMask(DefaultBoard5.gameDefinition.GreenTargetAreaMask))
		assert.False(t, cell.InMask(DefaultBoard5.gameDefinition.RedTargetAreaMask))
	}
}
