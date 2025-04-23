package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

// Size of the board (in rows and columns).
const BoardSize = 5

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

	if err != nil {
		return nil, err
	}

	// Check board validity.
	err = CheckBoard(&board)

	if err != nil {
		return nil, err
	}

	return &board, nil
}

// Check that the board is valid.
// Automatically called after loading a board from a state file.
func CheckBoard(board *BoardState) error {
	// Check that there are the right count of rows in the board.
	if len(board.Board) != BoardSize {
		return errors.New("invalid game state, please provide a valid game state")
	}

	// Check that every row has the right count of columns.
	for _, row := range board.Board {
		if len(row) != len(board.Board) {
			return errors.New("invalid game state, please provide a valid game state")
		}
	}

	// No error.
	return nil
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

// Print the game board to the console.
func PrintBoard(writer io.Writer, board *BoardState) {
	fmt.Fprintln(writer, "    1    2    3    4    5  ")
	fmt.Fprintln(writer, ". +----+----+----+----+----+")
	for rowIndex, row := range board.Board {
		fmt.Fprintf(writer, "%c ", rune(int('a')+rowIndex))
		fmt.Fprint(writer, "|")
		for _, cell := range row {
			if cell == 1 {
				fmt.Fprint(writer, " 🟢 ")
			} else if cell == 2 {
				fmt.Fprint(writer, " 🔴 ")
			} else {
				fmt.Fprint(writer, "    ")
			}
			fmt.Fprint(writer, "|")
		}
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, ". +----+----+----+----+----+")
	}
}
