package game

import (
	"math"
)

type MinMaxOptions = struct {
	maxDepth int
}

func (game *BoardState) minMaxEvaluateMove(depth int, minimizing bool, player Player, options MinMaxOptions, move []CellIdentifier) (float64, error) {
	virtualGame := game.Clone()
	if err := virtualGame.movePawn(move); err != nil {
		return 0, err
	}

	var moveScore float64
	if depth >= options.maxDepth {
		// Reached the max depth, perform a simple evaluation.
		distances := virtualGame.EvaluateDistance()
		moveScore = distances.Green - distances.Red
		if player == Red {
			moveScore = distances.Red - distances.Green
		}
	} else {
		// Deep exploration of possible moves.
		_, moveScore = virtualGame.MinMaxBestMove(depth+1, !minimizing, player, options)
	}

	return moveScore, nil
}

func (game *BoardState) MinMaxBestMove(depth int, minimizing bool, player Player, options MinMaxOptions) (bestMove []CellIdentifier, bestScore float64) {
	if minimizing {
		bestScore = math.MaxInt
	} else {
		bestScore = math.MinInt
	}

	for rowIndex, row := range game.Board {
		for columnIndex, cell := range row {
			if cell == Cell(game.CurrentPlayer) {
				from := CellIdentifier{int8(rowIndex), int8(columnIndex)}
				for _, move := range game.FindAllPossibleMoves(from) {
					eval, err := game.minMaxEvaluateMove(depth, minimizing, player, options, move)
					if err == nil {
						if minimizing {
							if eval <= bestScore {
								bestMove = move
								bestScore = eval
							}
						} else {
							if eval > bestScore {
								bestMove = move
								bestScore = eval
							}
						}
					}
				}
			}
		}
	}

	return
}

func (game *BoardState) DefaultMinMaxBestMove() (bestMove []CellIdentifier, bestScore float64) {
	return game.MinMaxBestMove(0, true, game.CurrentPlayer, MinMaxOptions{maxDepth: 3})
}
