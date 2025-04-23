package game

import (
	"errors"
	"strings"
)

type Cell int8

type CellIdentifier [2]int8

// Parse a cell identifier from the serialized cell identifier string.
func ParseCellIdentifier(serializedCellIdentifier string) (*CellIdentifier, error) {
	if len(serializedCellIdentifier) == 2 {
		// Ensure the string is lowercased.
		serializedCellIdentifier = strings.ToLower(serializedCellIdentifier)

		// Get the shift from 'a' character.
		row := serializedCellIdentifier[0] - 'a'
		// Get the shift from '1' character.
		column := serializedCellIdentifier[1] - '1'

		// Return the built cell identifier.
		return &CellIdentifier{int8(row), int8(column)}, nil
	} else {
		return nil, errors.New("invalid cell format")
	}
}
