package game

import (
	"math"
	"time"
)

type MinMaxOptions = struct {
	maxDepth int
	maxTime  int64
}

func (game *BoardState) minMaxEvaluateMove(depth int, maximizing bool, player Player, alpha float64, beta float64, options MinMaxOptions, move []CellIdentifier) (float64, float64, error) {
	virtualGame := game.Clone()
	if err := virtualGame.movePawn(move); err != nil {
		return 0, 0, err
	}

	// Compute the direct score after the move to prioritize moves with the best direct score when minmax move score is the same
	// (we want to improve our score as soon as possible).
	distances := virtualGame.EvaluateDistance()
	directScore := -distances.Green + distances.Red
	if player == Red {
		directScore = -distances.Red + distances.Green
	}

	if virtualGame.GetWinner() != None {
		if virtualGame.GetWinner() == player {
			// Try to win as soon as possible.
			return float64(math.MaxInt - depth*10_000), directScore, nil
		} else {
			// Try to lose as late as possible.
			return float64(math.MinInt + depth*10_000), directScore, nil
		}
	}

	var moveScore float64
	if depth >= options.maxDepth || time.Now().Unix() > options.maxTime {
		// Reached max depth or max time, use a simple evaluation.
		moveScore = directScore
	} else {
		// Deep exploration of possible moves.
		_, moveScore = virtualGame.MinMaxBestMove(depth+1, !maximizing, player, alpha, beta, options)
	}

	return moveScore, directScore, nil
}

func (game *BoardState) MinMaxBestMove(depth int, maximizing bool, player Player, alpha float64, beta float64, options MinMaxOptions) (bestMove []CellIdentifier, bestScore float64) {
	var bestDirectScore float64
	if maximizing {
		bestScore = math.MinInt
		bestDirectScore = math.MinInt
	} else {
		bestScore = math.MaxInt
		bestDirectScore = math.MaxInt
	}

	for rowIndex, row := range game.Board {
		for columnIndex, cell := range row {
			if cell == Cell(game.CurrentPlayer) {
				from := CellIdentifier{int8(rowIndex), int8(columnIndex)}
				for _, move := range game.FindAllPossibleMoves(from) {
					eval, directScore, err := game.minMaxEvaluateMove(depth, maximizing, player, alpha, beta, options, move)
					if err == nil {
						if maximizing {
							if eval > bestScore || (eval == bestScore && directScore > bestDirectScore) {
								bestMove = move
								bestScore = eval
								bestDirectScore = directScore
							}

							if bestScore >= beta {
								return
							}
							alpha = math.Max(alpha, bestScore)
						} else {
							if eval < bestScore || (eval == bestScore && directScore < bestDirectScore) {
								bestMove = move
								bestScore = eval
								bestDirectScore = directScore
							}

							if bestScore <= alpha {
								return
							}
							beta = math.Min(beta, bestScore)
						}
					}
				}
			}
		}
	}

	return
}

func (game *BoardState) DefaultMinMaxBestMove(maxTime int64) (bestMove []CellIdentifier, bestScore float64) {
	return game.MinMaxBestMove(0, true, game.CurrentPlayer, math.MinInt, math.MaxInt, MinMaxOptions{maxDepth: 4, maxTime: maxTime})
}
