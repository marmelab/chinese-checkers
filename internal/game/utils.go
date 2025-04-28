package game

// Copy a 2-dimensional array of cells.
func copyCellsArray(cells [][]Cell) [][]Cell {
	// Initialize all rows.
	clonedCells := make([][]Cell, len(cells))

	// Clone all rows.
	for rowIndex, row := range cells {
		// Clone the current row.
		clonedCells[rowIndex] = make([]Cell, len(row))
		copy(clonedCells[rowIndex], row)
	}

	return clonedCells
}
