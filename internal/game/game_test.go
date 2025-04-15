package game

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Chinese Checkers!"
	got := Hello()

	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}
