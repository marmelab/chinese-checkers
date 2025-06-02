package game

import "math"

func (game *BoardState) tryMove(bestScore *int, bestMove *[]CellIdentifier, player Player, from CellIdentifier, to CellIdentifier) {
	// Check cell existence in the board.
	if to.Row < 0 || to.Row >= int8(len(game.Board)) {
		return
	}
	if to.Column < 0 || to.Column >= int8(len(game.Board[to.Row])) {
		return
	}

	move := []CellIdentifier{from, to}

	virtualGame := game.Clone()
	virtualGame.CurrentPlayer = player
	if err := virtualGame.movePawn(move); err == nil {
		// Get the distance score of the current player after the move: the lower, the better.
		distance := virtualGame.EvaluateDistance()
		distanceScore := distance.Green
		if player == Red {
			distanceScore = distance.Red
		}

		if distanceScore < *bestScore {
			*bestMove = move
			*bestScore = distanceScore
		}
	}
}

// FindBestMove of the provided player.
func (game *BoardState) FindBestMove(player Player) (bestMove []CellIdentifier) {
	bestScore := math.MaxInt

	for rowIndex, row := range game.Board {
		for columnIndex, cell := range row {
			if cell == Cell(player) {
				from := CellIdentifier{int8(rowIndex), int8(columnIndex)}

				game.tryMove(&bestScore, &bestMove, player, from, CellIdentifier{from.Row, from.Column - 1})
				game.tryMove(&bestScore, &bestMove, player, from, CellIdentifier{from.Row, from.Column + 1})
				game.tryMove(&bestScore, &bestMove, player, from, CellIdentifier{from.Row - 1, from.Column})
				game.tryMove(&bestScore, &bestMove, player, from, CellIdentifier{from.Row + 1, from.Column})
			}
		}
	}

	return
}
