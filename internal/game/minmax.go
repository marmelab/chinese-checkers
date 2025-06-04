package game

import (
	"math"
	"time"
)

type MinMaxOptions = struct {
	maxDepth int
	maxTime  int64
}

func (game *BoardState) minMaxEvaluateMove(depth int, minimizing bool, player Player, options MinMaxOptions, move []CellIdentifier) (float64, error) {
	virtualGame := game.Clone()
	if err := virtualGame.movePawn(move); err != nil {
		return 0, err
	}

	if virtualGame.GetWinner() != None {
		if virtualGame.GetWinner() == player {
			// Try to win as soon as possible.
			return float64(math.MinInt + depth*10_000), nil
		} else {
			// Try to lose as late as possible.
			return float64(math.MaxInt - depth*10_000), nil
		}
	}

	var moveScore float64
	if depth >= options.maxDepth || time.Now().Unix() > options.maxTime {
		// Reached max depth or max time, perform a simple evaluation.
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

func (game *BoardState) DefaultMinMaxBestMove(maxTime int64) (bestMove []CellIdentifier, bestScore float64) {
	return game.MinMaxBestMove(0, true, game.CurrentPlayer, MinMaxOptions{maxDepth: 3, maxTime: maxTime})
}
