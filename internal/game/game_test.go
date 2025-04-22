package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	expected := "Hello, Chinese Checkers!"
	got := Hello()
	assert.Equal(t, expected, got, "they should be equal")
}
