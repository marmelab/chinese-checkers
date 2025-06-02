package game

import "math"

type BestMoveContext struct {
	player    Player
	move      []CellIdentifier
	bestMove  *[]CellIdentifier
	bestScore *int
}

func (game *BoardState) tryMove(context BestMoveContext, from CellIdentifier, to CellIdentifier) {
	// If the provided cell is already part of the move, do not explore it.
	for _, cell := range context.move {
		if cell == to {
			return
		}
	}

	// Check cell existence in the board.
	if to.Row < 0 || to.Row >= int8(len(game.Board)) {
		return
	}
	if to.Column < 0 || to.Column >= int8(len(game.Board[to.Row])) {
		return
	}

	move := []CellIdentifier{from, to}

	if isSimpleMove, err := game.CheckMoveLegality(from, to); err == nil {

		virtualGame := game.Clone()
		virtualGame.CurrentPlayer = context.player
		if err := virtualGame.movePawn(move); err != nil {
			return
		}

		// Full move from the starting position.
		fullMove := appendToMove(context.move, to)

		// Get the distance score of the current player after the move: the lower, the better.
		distance := virtualGame.EvaluateDistance()
		distanceScore := distance.Green
		if context.player == Red {
			distanceScore = distance.Red
		}

		if distanceScore < *context.bestScore {
			*context.bestMove = fullMove
			*context.bestScore = distanceScore
		}

		if !isSimpleMove {
			// If we just did a jump move, we can continue to jump.
			virtualGame.FindBestMoveFromCell(BestMoveContext{
				player:    context.player,
				move:      fullMove,
				bestMove:  context.bestMove,
				bestScore: context.bestScore,
			}, false)
		}
	}
}

func (game *BoardState) FindBestMoveFromCell(context BestMoveContext, canDoSimpleMoves bool) {
	from := context.move[len(context.move)-1]

	// Check simple paths.
	if canDoSimpleMoves {
		game.tryMove(context, from, CellIdentifier{from.Row, from.Column - 1})
		game.tryMove(context, from, CellIdentifier{from.Row, from.Column + 1})
		game.tryMove(context, from, CellIdentifier{from.Row - 1, from.Column})
		game.tryMove(context, from, CellIdentifier{from.Row + 1, from.Column})
	}

	// Check jump paths.
	game.tryMove(context, from, CellIdentifier{from.Row, from.Column - 2})
	game.tryMove(context, from, CellIdentifier{from.Row, from.Column + 2})
	game.tryMove(context, from, CellIdentifier{from.Row - 2, from.Column})
	game.tryMove(context, from, CellIdentifier{from.Row + 2, from.Column})
}

// FindBestMove of the provided player.
func (game *BoardState) FindBestMove(player Player) (bestMove []CellIdentifier) {
	bestScore := math.MaxInt

	for rowIndex, row := range game.Board {
		for columnIndex, cell := range row {
			if cell == Cell(player) {
				from := CellIdentifier{int8(rowIndex), int8(columnIndex)}
				game.FindBestMoveFromCell(BestMoveContext{
					player:    player,
					move:      []CellIdentifier{from},
					bestMove:  &bestMove,
					bestScore: &bestScore,
				}, true)
			}
		}
	}

	return
}

// A tree of valid move paths.
type PathsTree struct {
	Cell  CellIdentifier   `json:"cell"`
	Move  []CellIdentifier `json:"move"`
	Paths []PathsTree      `json:"paths"`
}

// Append the cell to a copied move list.
func appendToMove(move []CellIdentifier, cell CellIdentifier) (newMove []CellIdentifier) {
	newMove = make([]CellIdentifier, len(move)+1)
	copy(newMove, move)
	newMove[len(move)] = cell
	return
}

// Append the move candidate to the paths tree if the move is valid.
func (game BoardState) tryToAppendMoveCandidate(paths *PathsTree, to CellIdentifier) {
	// If the provided cell is already part of the move, do not explore it.
	for _, cell := range paths.Move {
		if cell == to {
			return
		}
	}

	// Check cell existence in the board.
	if to.Row < 0 || to.Row >= int8(len(game.Board)) {
		return
	}
	if to.Column < 0 || to.Column >= int8(len(game.Board[to.Row])) {
		return
	}

	if isSimpleMove, err := game.CheckMoveLegality(paths.Cell, to); err == nil {
		if isSimpleMove {
			// In the case of a simple move, no more paths are allowed.
			paths.Paths = append(paths.Paths, PathsTree{
				Cell:  to,
				Move:  appendToMove(paths.Move, to),
				Paths: []PathsTree{},
			})
		} else {
			virtualGame := game.Clone()
			if err := virtualGame.movePawn([]CellIdentifier{paths.Cell, to}); err != nil {
				return
			}
			// Keep the same current player to continue moving.
			virtualGame.CurrentPlayer = game.CurrentPlayer

			// In the case of a jump move, simple paths are no more allowed.
			paths.Paths = append(paths.Paths, virtualGame.FindValidMovesFrom(to, paths.Move, false))
		}
	}
}

// Find all valid move paths from a provided cell.
func (game BoardState) FindValidMovesFrom(cell CellIdentifier, previousCells []CellIdentifier, canDoSimpleMoves bool) (paths PathsTree) {
	paths.Cell = cell
	paths.Move = appendToMove(previousCells, cell)
	paths.Paths = []PathsTree{}

	// Try to add simple paths.
	if canDoSimpleMoves {
		game.tryToAppendMoveCandidate(&paths, CellIdentifier{paths.Cell.Row, paths.Cell.Column + 1})
		game.tryToAppendMoveCandidate(&paths, CellIdentifier{paths.Cell.Row, paths.Cell.Column - 1})
		game.tryToAppendMoveCandidate(&paths, CellIdentifier{paths.Cell.Row + 1, paths.Cell.Column})
		game.tryToAppendMoveCandidate(&paths, CellIdentifier{paths.Cell.Row - 1, paths.Cell.Column})
	}

	// Try to add jump paths.
	game.tryToAppendMoveCandidate(&paths, CellIdentifier{paths.Cell.Row, paths.Cell.Column + 2})
	game.tryToAppendMoveCandidate(&paths, CellIdentifier{paths.Cell.Row, paths.Cell.Column - 2})
	game.tryToAppendMoveCandidate(&paths, CellIdentifier{paths.Cell.Row + 2, paths.Cell.Column})
	game.tryToAppendMoveCandidate(&paths, CellIdentifier{paths.Cell.Row - 2, paths.Cell.Column})

	return
}
