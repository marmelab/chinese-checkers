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

// Evaluate the distance score of each player.
// The lower the better: a low score indicates that the player is near their target area.
func (board *BoardState) EvaluateDistance() (result Evaluation) {
	result.GreenScore = 0
	result.RedScore = 0

	for rowIndex, row := range board.Board {
		for columnIndex, cell := range row {
			cellPos := CellIdentifier{int8(rowIndex), int8(columnIndex)}

			if cell == GreenCell {
				if cellPos.InMask(board.gameDefinition.GreenTargetAreaMask) {
					continue
				}

				rowTargetDiff := len(board.Board) - 1 - rowIndex
				colTargetDiff := len(board.Board) - 1 - columnIndex

				result.GreenScore += euclideanDistance(rowTargetDiff, colTargetDiff)
			} else if cell == RedCell {
				if cellPos.InMask(board.gameDefinition.RedTargetAreaMask) {
					continue
				}

				result.RedScore += euclideanDistance(rowIndex, columnIndex)
			}
		}
	}

	return
}

// Evaluate the chances of winning of players.
func (board *BoardState) Evaluate() (result Evaluation) {
	distance := board.EvaluateDistance()

	totalScores := float64(distance.GreenScore + distance.RedScore)

	result.GreenScore = int(math.Round((float64(distance.RedScore) / totalScores) * 100))
	result.RedScore = int(math.Round((float64(distance.GreenScore) / totalScores) * 100))

	return
}
