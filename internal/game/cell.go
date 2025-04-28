package game

import (
	"errors"
	"fmt"
	"strings"
)

type Cell int8

const (
	EmptyCell Cell = 0
	GreenCell Cell = Cell(Green)
	RedCell   Cell = Cell(Red)
)

type CellIdentifier struct {
	Row    int8
	Column int8
}

// Convert a cell identifier to its string format.
func (cellIdentifier CellIdentifier) String() string {
	return fmt.Sprintf("%c%c", cellIdentifier.Row+'a', cellIdentifier.Column+'1')
}

// Parse a cell identifier from the serialized cell identifier string.
func (board BoardState) ParseCellIdentifier(serializedCellIdentifier string) (*CellIdentifier, error) {
	// Ensure the string is lowercased and trim whitespaces.
	serializedCellIdentifier = strings.ToLower(strings.Trim(serializedCellIdentifier, " \t\n"))

	if len(serializedCellIdentifier) == 2 {
		// Get the shift from 'a' character.
		row := int8(serializedCellIdentifier[0] - 'a')
		// Get the shift from '1' character.
		column := int8(serializedCellIdentifier[1] - '1')

		// Check cell validity in the board.
		if row < 0 || row >= int8(len(board.Board)) {
			return nil, fmt.Errorf("%s is not a valid cell", serializedCellIdentifier)
		}
		if column < 0 || column >= int8(len(board.Board[row])) {
			return nil, fmt.Errorf("%s is not a valid cell", serializedCellIdentifier)
		}

		// Return the built cell identifier.
		return &CellIdentifier{row, column}, nil
	} else {
		return nil, fmt.Errorf("invalid cell format '%s'", serializedCellIdentifier)
	}
}

// Parse a move list (a list of cell identifiers) from the serialized move list string.
func (board BoardState) ParseMoveList(serializedMoveList string) ([]CellIdentifier, error) {
	// Extract all the parts of the move list.
	serializedCells := strings.Split(serializedMoveList, ",")
	if len(serializedCells) < 2 {
		return nil, errors.New("you must provide at least two cells in a move")
	}

	// Parse all identifiers of the list.
	moveList := make([]CellIdentifier, len(serializedCells))
	for moveIndex, serializedCellIdentifier := range serializedCells {
		cellIdentifier, err := board.ParseCellIdentifier(serializedCellIdentifier)
		if err != nil {
			return nil, err
		}
		moveList[moveIndex] = *cellIdentifier
	}

	// Return the move list.
	return moveList, nil
}

// Determine if the cell is in the provided mask.
func (cell CellIdentifier) InMask(cellsMask [][]Cell) bool {
	// Detect overflows in the provided mask.
	if cell.Row >= int8(len(cellsMask)) || cell.Column >= int8(len(cellsMask[cell.Row])) {
		return false
	}

	return cellsMask[cell.Row][cell.Column] > 0
}
