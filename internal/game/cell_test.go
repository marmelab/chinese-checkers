package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCellIdentifierParser(t *testing.T) {
	output, err := ParseCellIdentifier("a3")
	assert.Nil(t, err)
	assert.Equal(t, &CellIdentifier{0, 2}, output, "should be the same identifier")

	output, err = ParseCellIdentifier(" H9")
	assert.Nil(t, err)
	assert.Equal(t, &CellIdentifier{7, 8}, output, "should be the same identifier")
}

func TestInvalidCellIdentifierParsing(t *testing.T) {
	_, err := ParseCellIdentifier("a")
	assert.Equal(t, err.Error(), "invalid cell format", "should be an invalid format error")
}

func TestMoveListParser(t *testing.T) {
	output, err := ParseMoveList("a1,b2,c3")
	assert.Nil(t, err)
	assert.Equal(t, []CellIdentifier{{0, 0}, {1, 1}, {2, 2}}, output)

	output, err = ParseMoveList("b4, b3, b2, a2")
	assert.Nil(t, err)
	assert.Equal(t, []CellIdentifier{{1, 3}, {1, 2}, {1, 1}, {0, 1}}, output)
}

func TestMoveListSizeError(t *testing.T) {
	_, err := ParseMoveList("a1")
	assert.Equal(t, "you must provide at least two cells in a move", err.Error(), "should be a too small move list error")
}

func TestMoveListFormatError(t *testing.T) {
	_, err := ParseMoveList("a1,b,c3")
	assert.Equal(t, "invalid cell format", err.Error(), "should be an invalid format error")
}
