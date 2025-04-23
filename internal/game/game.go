package game

import (
	"encoding/json"
	"os"
)

type Cell int8
type PlayerId int8

// The main board state.
type BoardState struct {
	Board         [][]Cell `json:"board"`
	CurrentPlayer PlayerId `json:"currentPlayer"`
}

// The default board.
var DefaultBoard = BoardState{
	Board: [][]Cell{
		{1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 0, 0, 0, 2},
		{0, 0, 0, 2, 2},
		{0, 0, 2, 2, 2},
	},
	CurrentPlayer: 1,
}

// Load a board from a state file.
func LoadBoard(filePath string) (*BoardState, error) {
	// Fully read the provided file.
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var board BoardState
	err = json.Unmarshal(fileData, &board)

	return &board, err
}

// Initialize a board state from the provided state file.
func InitBoard(filePath *string) (*BoardState, error) {
	if filePath != nil {
		// A file path has been provided, load it.
		return LoadBoard(*filePath)
	} else {
		// No file path, return the default board.
		return &DefaultBoard, nil
	}
}
