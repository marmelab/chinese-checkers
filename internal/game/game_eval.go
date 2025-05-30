package game

import (
	"math"
)

func euclideanDistance(x int, y int) int {
	return int(math.Round(math.Sqrt(float64(x*x + y*y))))
}

type Evaluation = struct {
	Green int `json:"green"`
	Red   int `json:"red"`
}

// Evaluate the chances of winning of players.
func (board *BoardState) Evaluate() (result Evaluation) {
	greenDistance := 0
	redDistance := 0

	for rowIndex, row := range board.Board {
		for columnIndex, cell := range row {
			cellPos := CellIdentifier{int8(rowIndex), int8(columnIndex)}

			if cell == GreenCell {
				if cellPos.InMask(board.gameDefinition.GreenTargetAreaMask) {
					continue
				}

				rowTargetDiff := len(board.Board) - 1 - rowIndex
				colTargetDiff := len(board.Board) - 1 - columnIndex

				greenDistance += euclideanDistance(rowTargetDiff, colTargetDiff)
			} else if cell == RedCell {
				if cellPos.InMask(board.gameDefinition.RedTargetAreaMask) {
					continue
				}

				redDistance += euclideanDistance(rowIndex, columnIndex)
			}
		}
	}

	distancesTotal := float64(greenDistance + redDistance)

	result.Green = int(math.Round((float64(redDistance) / distancesTotal) * 100))
	result.Red = int(math.Round((float64(greenDistance) / distancesTotal) * 100))

	return
}
