package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
)

// Game definition structure.
type GameDefinition struct {
	BoardSize           int8
	PlayerPawnsNumber   int8
	GreenTargetAreaMask [][]Cell
	RedTargetAreaMask   [][]Cell
}

// Allowed game definitions.
var gameDefinitions = [...]GameDefinition{
	{
		BoardSize:         5,
		PlayerPawnsNumber: 6,

		GreenTargetAreaMask: [][]Cell{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1},
			{0, 0, 0, 1, 1},
			{0, 0, 1, 1, 1},
		},
		RedTargetAreaMask: [][]Cell{
			{1, 1, 1, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
	},
	{
		BoardSize:         7,
		PlayerPawnsNumber: 10,

		GreenTargetAreaMask: [][]Cell{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0, 1, 1},
			{0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 1, 1, 1, 1},
		},
		RedTargetAreaMask: [][]Cell{
			{1, 1, 1, 1, 0, 0, 0},
			{1, 1, 1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		},
	},
}

// The main board state.
type BoardState struct {
	Board          [][]Cell         `json:"board"`
	CurrentPlayer  Player           `json:"currentPlayer"`
	LastMove       []CellIdentifier `json:"lastMove,omitempty"`
	stateFile      *string
	gameDefinition *GameDefinition
}

// The default board for a 5x5 game.
var DefaultBoard5 = BoardState{
	Board: [][]Cell{
		{1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 0, 0, 0, 2},
		{0, 0, 0, 2, 2},
		{0, 0, 2, 2, 2},
	},
	CurrentPlayer:  Green,
	stateFile:      nil,
	gameDefinition: &gameDefinitions[0],
}

// The default board for a 7x7 game.
var DefaultBoard7 = BoardState{
	Board: [][]Cell{
		{1, 1, 1, 1, 0, 0, 0},
		{1, 1, 1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 2},
		{0, 0, 0, 0, 0, 2, 2},
		{0, 0, 0, 0, 2, 2, 2},
		{0, 0, 0, 2, 2, 2, 2},
	},
	CurrentPlayer:  Green,
	stateFile:      nil,
	gameDefinition: &gameDefinitions[1],
}

// Initialize a default board state for a 5x5 game.
func NewDefaultBoard5() *BoardState {
	board := DefaultBoard5.Clone()
	// Chose a random player to start.
	board.CurrentPlayer = RandomPlayer()
	return board
}

// Initialize a default board state for a 7x7 game.
func NewDefaultBoard7() *BoardState {
	board := DefaultBoard7.Clone()
	// Chose a random player to start.
	board.CurrentPlayer = RandomPlayer()
	return board
}

// Initialize a board from a state file.
func NewBoardFromStateFile(filePath string) (*BoardState, error) {
	// Fully read the provided file.
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var board BoardState
	err = json.Unmarshal(fileData, &board)

	if err != nil {
		return nil, err
	}

	// Try to detect a corresponding game definition.
	for _, gameDefinition := range gameDefinitions {
		// Check if there is the right count of rows in the board.
		if len(board.Board) == int(gameDefinition.BoardSize) {
			// The game definition board size matches the current board size.
			boardGameDefinition := gameDefinition
			board.gameDefinition = &boardGameDefinition
			break
		}
	}

	// Check board validity.
	err = board.CheckBoardValidity()

	if err != nil {
		return nil, err
	}

	// Store the used file path in the board state.
	board.stateFile = &filePath

	return &board, nil
}

// CheckBoardValidity that the board is valid.
// Automatically called after loading a board from a state file.
func (board *BoardState) CheckBoardValidity() error {
	// Could not detect a game definition, the loaded game state is invalid.
	if board.gameDefinition == nil {
		return errors.New("invalid game state, please provide a valid game state")
	}

	// Check that there are the right count of rows in the board.
	if len(board.Board) != int(board.gameDefinition.BoardSize) {
		return errors.New("invalid game state, please provide a valid game state")
	}

	if board.CurrentPlayer < Green || board.CurrentPlayer > Red {
		// Invalid player ID, return an error.
		return fmt.Errorf("%d is not a valid player ID", board.CurrentPlayer)
	}

	// Count of pawns for each player (index 0 = player 1, index 1 = player 2)
	playerPawnsCounts := []int{0, 0}
	// Check that every row has the right count of columns.
	for _, row := range board.Board {
		if len(row) != len(board.Board) {
			return errors.New("invalid game state, please provide a valid game state")
		}
		// Count the pawns of each player in the current row.
		for _, cell := range row {
			if cell > EmptyCell {
				// There is a player on the current cell.
				if cell > RedCell {
					// Invalid player ID, return an error.
					return fmt.Errorf("%d is not a valid player ID", cell)
				}

				// Increment player pawns count.
				playerPawnsCounts[cell-1] += 1
			}
		}
	}

	// Check that there are enough pawns for a player in the board.
	if playerPawnsCounts[0] != int(board.gameDefinition.PlayerPawnsNumber) || playerPawnsCounts[1] != int(board.gameDefinition.PlayerPawnsNumber) {
		return errors.New("invalid game state, please provide a valid game state")
	}

	// No error.
	return nil
}

// Clone the board state.
func (board *BoardState) Clone() *BoardState {
	// Initialize a new board.
	clonedBoard := &BoardState{
		Board:          copyCellsArray(board.Board),
		CurrentPlayer:  board.CurrentPlayer,
		stateFile:      board.stateFile,
		gameDefinition: new(GameDefinition), // Initialize a game definition pointer.
	}

	// Clone the game definition pointer.
	*clonedBoard.gameDefinition = *board.gameDefinition
	clonedBoard.gameDefinition.GreenTargetAreaMask = copyCellsArray(board.gameDefinition.GreenTargetAreaMask)
	clonedBoard.gameDefinition.RedTargetAreaMask = copyCellsArray(board.gameDefinition.RedTargetAreaMask)

	return clonedBoard
}

// Save the board state in memory.
func (board *BoardState) SaveState(filePath string) error {
	// Convert the board to JSON.
	boardJson, err := json.Marshal(board)
	if err != nil {
		return err
	}
	// Write the new state file.
	return os.WriteFile(filePath, boardJson, 0644)
}

// Check that the provided move is legal.
// A move is legal:
// - when the pawn moves to an adjacent cell (simple move)
// - when the pawn jumps over another one (jump move)
// The returned boolean indicates if the provided move is a simple move.
func (board *BoardState) CheckMoveLegality(from CellIdentifier, to CellIdentifier) (bool, error) {
	// Compute the column diff of the move.
	columnDiff := math.Abs(float64(from.Column - to.Column))
	// Compute the row diff of the move.
	rowDiff := math.Abs(float64(from.Row - to.Row))

	// Check that the target cell is free.
	if board.Board[to.Row][to.Column] != EmptyCell {
		return false, fmt.Errorf("there is already a pawn on %s", to.String())
	}

	if columnDiff+rowDiff == 1 {
		// Only 1 difference, the move is legal.
		return true, nil
	}

	// Check jumps over a pawn.

	if columnDiff >= 2 && rowDiff == 0 {
		// Moving straightly in a row, check that there is ONE pawn in the middle.
		if columnDiff > 2 {
			return false, errors.New("a pawn cannot jump over more than one pawn")
		}
		// Check that there is a pawn in the middle of the jump.
		middleColumn := from.Column + ((to.Column - from.Column) / 2)
		if board.Board[from.Row][middleColumn] != EmptyCell {
			return false, nil
		}
	}
	if rowDiff >= 2 && columnDiff == 0 {
		// Moving straightly in a column, check that there is ONE pawn in the middle.
		if rowDiff > 2 {
			return false, errors.New("a pawn cannot jump over more than one pawn")
		}
		// Check that there is a pawn in the middle of the jump.
		middleRow := from.Row + ((to.Row - from.Row) / 2)
		if board.Board[middleRow][from.Column] != EmptyCell {
			return false, nil
		}
	}

	// The move is illegal (more than 1 difference, or no difference).

	if rowDiff == 1 && columnDiff == 1 {
		// Detected a diagonal move, return a specific error.
		return false, errors.New("a pawn cannot move in diagonal")
	}

	return false, fmt.Errorf("'%s' cannot be reached from '%s'", to.String(), from.String())
}

// Check legality of all successive moves.
// Simple move = move to an adjacent cell.
func (board *BoardState) CheckMovesLegality(moveList []CellIdentifier, disallowSimpleMoves bool) error {
	// Check move list size.
	if len(moveList) < 2 {
		return errors.New("you must provide at least two cells in a move")
	}

	// Check the first move of the list.
	isSimpleMove, err := board.CheckMoveLegality(moveList[0], moveList[1])
	if err != nil {
		return err
	}

	// If simple moves are disallowed, check that the current move does not target an adjacent cell.
	if disallowSimpleMoves && isSimpleMove {
		return errors.New("cannot move to an adjacent cell after a jump")
	}

	if len(moveList) == 2 {
		// Only 2 positions in the list = only one move, and it's legal.
		return nil
	} else {
		// More than 2 positions in the list, check other moves recursively.

		// If the current move is a simple move, no more moves are allowed, the move list should have ended.
		if isSimpleMove {
			return errors.New("cannot continue moving after moving to an adjacent cell")
		}

		// The current move is a legal jump, check the other moves with simple moves disallowed.
		return board.CheckMovesLegality(moveList[1:], true)
	}
}

// Move a pawn of the board.
func (board *BoardState) MovePawn(serializedMoveList string) error {
	// If there is a winner, moving a pawn is disallowed.
	winner := board.GetWinner()
	if winner != None {
		return fmt.Errorf("cannot move a pawn: %s has won", winner.Name())
	}

	// Parse the move list.
	moveList, err := board.ParseMoveList(serializedMoveList)
	if err != nil {
		return err
	}

	// Ensure that there is a pawn at start position.
	startPawn := board.Board[moveList[0].Row][moveList[0].Column]
	if startPawn == EmptyCell {
		return fmt.Errorf("there is no pawn on %s", moveList[0].String())
	}
	// Ensure that the current player can move this pawn.
	if startPawn != Cell(board.CurrentPlayer) {
		return fmt.Errorf("you cannot move a %s pawn", Player(startPawn).Color())
	}

	// Check all successive moves legality, allowing only one simple move.
	if err = board.CheckMovesLegality(moveList, false); err != nil {
		return err
	}

	// Move the start pawn to the end position.
	board.Board[moveList[len(moveList)-1].Row][moveList[len(moveList)-1].Column] = startPawn
	// Remove the start pawn from its previous position.
	board.Board[moveList[0].Row][moveList[0].Column] = 0

	// Store the current move list as the last move.
	board.LastMove = moveList

	// After moving a pawn, switch player turn.
	if board.CurrentPlayer == Green {
		board.CurrentPlayer = Red
	} else {
		board.CurrentPlayer = Green
	}

	return nil
}

// Move a pawn of the board and save the new board state to the stored state file.
func (board *BoardState) MovePawnAndSave(serializedMoveList string) error {
	// Try to move a pawn using the provided move list.
	if err := board.MovePawn(serializedMoveList); err != nil {
		return err
	}
	if board.stateFile != nil {
		// There is a state file, save the new board state to it.
		if err := board.SaveState(*board.stateFile); err != nil {
			return err
		}
	}
	return nil
}

// Count pawns of each player that are in the player target area.
func (board BoardState) CountPawnsInTargetAreas() (greenPawns int8, redPawns int8) {
	// Evaluate all cells of the board to determine if there is a pawn in the target area.
	for rowIndex, row := range board.Board {
		for columnIndex, cell := range row {
			// Initialize a cell position.
			cellPos := CellIdentifier{int8(rowIndex), int8(columnIndex)}

			// Increment the pawns counter of the player if it is in the target area mask.
			if cell == GreenCell && cellPos.InMask(board.gameDefinition.GreenTargetAreaMask) {
				greenPawns++
			} else if cell == RedCell && cellPos.InMask(board.gameDefinition.RedTargetAreaMask) {
				redPawns++
			}
		}
	}
	return greenPawns, redPawns
}

// Get the winner.
// Return None if there is no winner.
func (board BoardState) GetWinner() Player {
	// Get green and red pawns in the target area of each player.
	greenPawns, redPawns := board.CountPawnsInTargetAreas()

	// Check if the green player has won.
	if greenPawns == int8(board.gameDefinition.PlayerPawnsNumber) {
		return Green
	}

	// Check if the red player has won.
	if redPawns == int8(board.gameDefinition.PlayerPawnsNumber) {
		return Red
	}

	return None
}

// Find out who is the previous player.
func (board BoardState) GetPreviousPlayer() Player {
	previousPlayer := Green
	if board.CurrentPlayer == Green {
		previousPlayer = Red
	}
	return previousPlayer
}
