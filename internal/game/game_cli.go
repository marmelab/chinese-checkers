package game

import (
	"fmt"
	"io"

	"github.com/go-color-term/go-color-term/coloring"
)

// Print the game board to the console.
func (board *BoardState) Print(writer io.Writer) {
	fmt.Fprintln(writer, "     1    2    3    4    5  ")
	fmt.Fprintln(writer, " . +----+----+----+----+----+")
	for rowIndex, row := range board.Board {
		fmt.Fprintf(writer, " %c ", rune(int('a')+rowIndex))
		fmt.Fprint(writer, "|")
		for columnIndex, cell := range row {
			// Initialize a cell position.
			cellPos := CellIdentifier{int8(rowIndex), int8(columnIndex)}

			if cell == GreenCell {
				pawnCellStr := " 🟢 "
				if cellPos.InMask(GreenTargetAreaMask) {
					// The cell is in the target area, add a background color.
					pawnCellStr = coloring.For(pawnCellStr).Background().Rgb(0, 70, 0).String()
				}
				fmt.Fprint(writer, pawnCellStr)
			} else if cell == RedCell {
				pawnCellStr := " 🔴 "
				if cellPos.InMask(RedTargetAreaMask) {
					// The cell is in the target area, add a background color.
					pawnCellStr = coloring.For(pawnCellStr).Background().Rgb(84, 0, 0).String()
				}
				fmt.Fprint(writer, pawnCellStr)
			} else {
				fmt.Fprint(writer, "    ")
			}
			fmt.Fprint(writer, "|")
		}
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, " . +----+----+----+----+----+")
	}
}

// Print the current game score.
func (board *BoardState) PrintScore(writer io.Writer) {
	greenPawns, redPawns := board.CountPawnsInTargetAreas()
	fmt.Fprintf(writer,
		coloring.For("Green").Bold().Green().String()+": %d/%d, "+coloring.For("Red").Bold().Red().String()+": %d/%d\n",
		greenPawns, PlayerPawnsNumber,
		redPawns, PlayerPawnsNumber,
	)
}
