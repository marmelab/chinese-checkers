package game

import (
	"fmt"
	"io"
)

// Print the game board to the console.
func (board *BoardState) Print(writer io.Writer) {
	fmt.Fprintln(writer, "     1    2    3    4    5  ")
	fmt.Fprintln(writer, " . +----+----+----+----+----+")
	for rowIndex, row := range board.Board {
		fmt.Fprintf(writer, " %c ", rune(int('a')+rowIndex))
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
		fmt.Fprintln(writer, " . +----+----+----+----+----+")
	}
}
