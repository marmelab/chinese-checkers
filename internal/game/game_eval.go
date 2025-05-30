package game

import (
	"math"
)

func euclideanDistance(x int, y int) int {
	return int(math.Round(math.Sqrt(float64(x*x + y*y))))
}

type Evaluation = struct {
	GreenScore int `json:"green"`
	RedScore   int `json:"red"`
}

// Evaluate the chances of winning of players.
func (board *BoardState) Evaluate() (result Evaluation) {
	greenDistanceScore := 0
	redDistanceScore := 0

	for rowIndex, row := range board.Board {
		for columnIndex, cell := range row {
			cellPos := CellIdentifier{int8(rowIndex), int8(columnIndex)}

			if cell == GreenCell {
				if cellPos.InMask(board.gameDefinition.GreenTargetAreaMask) {
					continue
				}

				rowTargetDiff := len(board.Board) - 1 - rowIndex
				colTargetDiff := len(board.Board) - 1 - columnIndex

				greenDistanceScore += euclideanDistance(rowTargetDiff, colTargetDiff)
			} else if cell == RedCell {
				if cellPos.InMask(board.gameDefinition.RedTargetAreaMask) {
					continue
				}

				redDistanceScore += euclideanDistance(rowIndex, columnIndex)
			}
		}
	}

	totalScores := float64(greenDistanceScore + redDistanceScore)

	result.GreenScore = int(math.Round((float64(redDistanceScore) / totalScores) * 100))
	result.RedScore = int(math.Round((float64(greenDistanceScore) / totalScores) * 100))

	return
}
