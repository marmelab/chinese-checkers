package game

import (
	"errors"
	"fmt"
	"strings"
)

type Cell int8

type CellIdentifier [2]int8

// Parse a cell identifier from the serialized cell identifier string.
func ParseCellIdentifier(serializedCellIdentifier string) (*CellIdentifier, error) {
	// Ensure the string is lowercased and trim whitespaces.
	serializedCellIdentifier = strings.ToLower(strings.Trim(serializedCellIdentifier, " \t\n"))

	if len(serializedCellIdentifier) == 2 {
		// Get the shift from 'a' character.
		row := serializedCellIdentifier[0] - 'a'
		// Get the shift from '1' character.
		column := serializedCellIdentifier[1] - '1'

		// Return the built cell identifier.
		return &CellIdentifier{int8(row), int8(column)}, nil
	} else {
		return nil, fmt.Errorf("invalid cell format '%s'", serializedCellIdentifier)
	}
}

// Parse a move list (a list of cell identifiers) from the serialized move list string.
func ParseMoveList(serializedMoveList string) ([]CellIdentifier, error) {
	// Extract all the parts of the move list.
	serializedCells := strings.Split(serializedMoveList, ",")
	if len(serializedCells) < 2 {
		return nil, errors.New("you must provide at least two cells in a move")
	}

	// Parse all identifiers of the list.
	moveList := make([]CellIdentifier, len(serializedCells))
	for moveIndex, serializedCellIdentifier := range serializedCells {
		cellIdentifier, err := ParseCellIdentifier(serializedCellIdentifier)
		if err != nil {
			return nil, err
		}
		moveList[moveIndex] = *cellIdentifier
	}

	// Return the move list.
	return moveList, nil
}
