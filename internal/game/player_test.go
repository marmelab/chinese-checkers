package game

import (
	"testing"

	"github.com/go-color-term/go-color-term/coloring"
	"github.com/stretchr/testify/assert"
)

func TestPlayerColor(t *testing.T) {
	assert.Equal(t, "green", Green.Color())
	assert.Equal(t, "red", Red.Color())
}

func TestPlayerName(t *testing.T) {
	assert.Equal(t, "Green", Green.Name())
	assert.Equal(t, "Red", Red.Name())
}

func TestPlayerColoredName(t *testing.T) {
	assert.Equal(t, coloring.For("Green").Bold().Green().String(), Green.ColoredName())
	assert.Equal(t, coloring.For("Red").Bold().Red().String(), Red.ColoredName())
}
