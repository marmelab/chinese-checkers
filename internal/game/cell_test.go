package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCellIdentifierParser(t *testing.T) {
	output, err := ParseCellIdentifier("a3")
	assert.Nil(t, err)
	assert.Equal(t, &CellIdentifier{0, 2}, output, "should be the same identifier")

	output, err = ParseCellIdentifier("h9")
	assert.Nil(t, err)
	assert.Equal(t, &CellIdentifier{7, 8}, output, "should be the same identifier")
}

func TestInvalidCellIdentifierParsing(t *testing.T) {
	input := "a"

	_, err := ParseCellIdentifier(input)

	assert.Equal(t, err.Error(), "invalid cell format", "should be an invalid format error")
}
