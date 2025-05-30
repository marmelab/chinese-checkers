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

// Evaluate the distance score of each player.
// The lower the better: a low score indicates that the player is near their target area.
func (board *BoardState) EvaluateDistance() (result Evaluation) {
	result.Green = 0
	result.Red = 0

	for rowIndex, row := range board.Board {
		for columnIndex, cell := range row {
			cellPos := CellIdentifier{int8(rowIndex), int8(columnIndex)}

			if cell == GreenCell {
				if cellPos.InMask(board.gameDefinition.GreenTargetAreaMask) {
					continue
				}

				rowTargetDiff := len(board.Board) - 1 - rowIndex
				colTargetDiff := len(board.Board) - 1 - columnIndex

				result.Green += euclideanDistance(rowTargetDiff, colTargetDiff)
			} else if cell == RedCell {
				if cellPos.InMask(board.gameDefinition.RedTargetAreaMask) {
					continue
				}

				result.Red += euclideanDistance(rowIndex, columnIndex)
			}
		}
	}

	return
}

// Evaluate the chances of winning of players.
func (board *BoardState) Evaluate() (result Evaluation) {
	distance := board.EvaluateDistance()

	distancesTotal := float64(distance.Green + distance.Red)

	result.Green = int(math.Round((float64(distance.Red) / distancesTotal) * 100))
	result.Red = int(math.Round((float64(distance.Green) / distancesTotal) * 100))

	return
}
