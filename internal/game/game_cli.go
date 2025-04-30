package game

import (
	"fmt"
	"io"

	"github.com/go-color-term/go-color-term/coloring"
)

// Print the last move.
func (board *BoardState) PrintLastMove(writer io.Writer) {
	// Only print the last move if there actually is a last move.
	if len(board.LastMove) > 0 {
		// Find out who is the previous player.
		previousPlayer := board.GetPreviousPlayer()

		// Print the previous player move.
		fmt.Fprintf(writer, "Last move from %s player: ", previousPlayer.ColoredName())

		lastIndex := len(board.LastMove) - 1
		for cellIndex, cell := range board.LastMove {
			// Print each cell string, with a comma to separate move parts.
			fmt.Fprint(writer, cell.String())
			if cellIndex != lastIndex {
				fmt.Fprint(writer, ", ")
			}
		}

		fmt.Fprintln(writer)
	}
}

// Print the game board to the console.
func (board *BoardState) Print(writer io.Writer) {
	// Find out who is the previous player.
	previousPlayer := board.GetPreviousPlayer()

	// Columns indices.
	fmt.Fprint(writer, "    ")
	for columnIndex := range board.Board[0] {
		fmt.Fprintf(writer, " %c   ", rune(int('1')+columnIndex))
	}
	fmt.Fprintln(writer)

	// First board edge.
	fmt.Fprint(writer, " . +")
	for range board.Board[0] {
		fmt.Fprintf(writer, "----+")
	}
	fmt.Fprintln(writer)

	for rowIndex, row := range board.Board {
		fmt.Fprintf(writer, " %c ", rune(int('a')+rowIndex))
		fmt.Fprint(writer, "|")

		for columnIndex, cell := range row {
			// Initialize a cell position.
			cellPos := CellIdentifier{int8(rowIndex), int8(columnIndex)}

			if cell == GreenCell {
				pawnCellStr := " 🟢 "
				if cellPos.InMask(board.gameDefinition.GreenTargetAreaMask) {
					// The cell is in the target area, add a background color.
					pawnCellStr = coloring.For(pawnCellStr).Background().Rgb(0, 70, 0).String()
				}
				fmt.Fprint(writer, pawnCellStr)
			} else if cell == RedCell {
				pawnCellStr := " 🔴 "
				if cellPos.InMask(board.gameDefinition.RedTargetAreaMask) {
					// The cell is in the target area, add a background color.
					pawnCellStr = coloring.For(pawnCellStr).Background().Rgb(84, 0, 0).String()
				}
				fmt.Fprint(writer, pawnCellStr)
			} else {
				emptyCellStr := "    "

				// Show a dimmed cell if it was the previous position of a cell.
				for _, cellPos := range board.LastMove {
					// Try to find a previous cell matching the current one.
					if cellPos.Row == int8(rowIndex) && cellPos.Column == int8(columnIndex) {
						// Build a small circle with a dimmed color of the previous player.
						emptyCellBuilder := coloring.For("  ⬤ ")
						if previousPlayer == Green {
							emptyCellBuilder = emptyCellBuilder.Rgb(0, 70, 0)
						} else {
							emptyCellBuilder = emptyCellBuilder.Rgb(84, 0, 0)
						}
						emptyCellStr = emptyCellBuilder.String()
						break
					}
				}
				fmt.Fprint(writer, emptyCellStr)
			}

			fmt.Fprint(writer, "|")
		}
		fmt.Fprintln(writer)

		// Board edge between rows.
		fmt.Fprint(writer, " . +")
		for range board.Board[0] {
			fmt.Fprintf(writer, "----+")
		}
		fmt.Fprintln(writer)
	}
}

// Print the current game score.
func (board *BoardState) PrintScore(writer io.Writer) {
	greenPawns, redPawns := board.CountPawnsInTargetAreas()
	fmt.Fprintf(writer,
		"     "+Green.ColoredName()+": %d/%d, "+Red.ColoredName()+": %d/%d\n",
		greenPawns, board.gameDefinition.PlayerPawnsNumber,
		redPawns, board.gameDefinition.PlayerPawnsNumber,
	)
}
