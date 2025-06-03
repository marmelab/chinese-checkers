package game

import (
	"math"
)

func euclideanDistance(x int, y int) float64 {
	return math.Sqrt(float64(x*x + y*y))
}

type Evaluation = struct {
	Green int `json:"green"`
	Red   int `json:"red"`
}

type DistanceEvaluation = struct {
	Green float64 `json:"green"`
	Red   float64 `json:"red"`
}

// Evaluate the distance score of each player.
// The lower the better: a low score indicates that the player is near their target area.
func (board *BoardState) EvaluateDistance() (result DistanceEvaluation) {
	result.Green = 0
	result.Red = 0

	greenTarget := CellIdentifier{int8(len(board.Board) - 1), int8(len(board.Board) - 1)}
	redTarget := CellIdentifier{0, 0}

	for rowIndex, row := range board.Board {
		for columnIndex, cell := range row {
			cellPos := CellIdentifier{int8(rowIndex), int8(columnIndex)}

			modifier := 1.0

			if cell == GreenCell {
				// Pawns already in the target area are less important.
				if cellPos.InMask(board.gameDefinition.GreenTargetAreaMask) {
					modifier = 0.5
				}

				rowTargetDiff := int(greenTarget.Row) - rowIndex
				colTargetDiff := int(greenTarget.Column) - columnIndex

				result.Green += modifier * euclideanDistance(rowTargetDiff, colTargetDiff)
			} else if cell == RedCell {
				// Pawns already in the target area are less important.
				if cellPos.InMask(board.gameDefinition.RedTargetAreaMask) {
					modifier = 0.5
				}

				rowTargetDiff := int(redTarget.Row) - rowIndex
				colTargetDiff := int(redTarget.Column) - columnIndex

				result.Red += modifier * euclideanDistance(rowTargetDiff, colTargetDiff)
			}
		}
	}

	return
}

// Evaluate the chances of winning of players.
func (board *BoardState) Evaluate() (result Evaluation) {
	distance := board.EvaluateDistance()

	distancesTotal := distance.Green + distance.Red

	result.Green = int(math.Round((distance.Red / distancesTotal) * 100))
	result.Red = int(math.Round((distance.Green / distancesTotal) * 100))

	return
}
