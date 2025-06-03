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

// FindTargetCell for the provided player, used in distance evaluation.
func (board *BoardState) FindTargetCell(player Player) CellIdentifier {
	// Take the other corner of the board when for green target.
	cornerOrigin := 0
	cornerFactor := 1
	if player == Green {
		cornerOrigin = len(board.Board) - 1
		cornerFactor = -1
	}

	// First check the edge of the target area.
	for i := range 4 {
		if board.Board[cornerOrigin+0][cornerOrigin+cornerFactor*i] != Cell(player) {
			return CellIdentifier{Row: int8(cornerOrigin + 0), Column: int8(cornerOrigin + cornerFactor*i)}
		}
		if board.Board[cornerOrigin+cornerFactor*i][cornerOrigin+0] != Cell(player) {
			return CellIdentifier{Row: int8(cornerOrigin + cornerFactor*i), Column: int8(cornerOrigin + 0)}
		}
	}
	// Then check the 3 other cells of the target area.
	for i := range 2 {
		if board.Board[cornerOrigin+cornerFactor*1][cornerOrigin+cornerFactor*(i+1)] != Cell(player) {
			return CellIdentifier{Row: int8(cornerOrigin + cornerFactor*1), Column: int8(cornerOrigin + cornerFactor*(i+1))}
		}
		if board.Board[cornerOrigin+cornerFactor*(i+1)][cornerOrigin+cornerFactor*1] != Cell(player) {
			return CellIdentifier{Row: int8(cornerOrigin + cornerFactor*(i+1)), Column: int8(cornerOrigin + cornerFactor*1)}
		}
	}

	return CellIdentifier{int8(cornerOrigin), int8(cornerOrigin)}
}

// Evaluate the distance score of each player.
// The lower the better: a low score indicates that the player is near their target area.
func (board *BoardState) EvaluateDistance() (result DistanceEvaluation) {
	result.Green = 0
	result.Red = 0

	greenTarget := board.FindTargetCell(Green)
	redTarget := board.FindTargetCell(Red)

	for rowIndex, row := range board.Board {
		for columnIndex, cell := range row {
			cellPos := CellIdentifier{int8(rowIndex), int8(columnIndex)}

			modifier := 1.0

			if cell == GreenCell {
				// Pawns already in the target area are less important.
				if cellPos.InMask(board.gameDefinition.GreenTargetAreaMask) {
					modifier = 0.25
				}

				rowTargetDiff := int(greenTarget.Row) - rowIndex
				colTargetDiff := int(greenTarget.Column) - columnIndex

				result.Green += modifier * euclideanDistance(rowTargetDiff, colTargetDiff)
			} else if cell == RedCell {
				// Pawns already in the target area are less important.
				if cellPos.InMask(board.gameDefinition.RedTargetAreaMask) {
					modifier = 0.25
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
