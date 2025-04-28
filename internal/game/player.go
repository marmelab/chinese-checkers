package game

import (
	"math/rand"

	"github.com/go-color-term/go-color-term/coloring"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Player int8

const (
	Green Player = 1
	Red   Player = 2
)

// Generate a random player.
func RandomPlayer() Player {
	return Player(rand.Intn(2) + 1)
}

// Get color name of the current player ID.
func (player Player) Color() string {
	if player == 1 {
		return "green"
	}
	return "red"
}

// Get the name of the player.
func (player Player) Name() string {
	return cases.Title(language.English).String(player.Color())
}

// Get the colored name of the player.
func (player Player) ColoredName() string {
	if player == 1 {
		return coloring.For(player.Name()).Bold().Green().String()
	}
	return coloring.For(player.Name()).Bold().Red().String()
}

// Green player target area mask.
var GreenTargetAreaMask = [][]Cell{
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 1},
	{0, 0, 0, 1, 1},
	{0, 0, 1, 1, 1},
}

// Green player target area mask.
var RedTargetAreaMask = [][]Cell{
	{1, 1, 1, 0, 0},
	{1, 1, 0, 0, 0},
	{1, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0},
}
