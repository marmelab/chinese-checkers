package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
)

// Size of the board (in rows and columns).
const BoardSize = 5

// Number of pawns of a player.
const PlayerPawnsNumber = 6

type PlayerId int8

// The main board state.
type BoardState struct {
	Board         [][]Cell `json:"board"`
	CurrentPlayer PlayerId `json:"currentPlayer"`
	stateFile     *string
}

// The default board.
var DefaultBoard = BoardState{
	Board: [][]Cell{
		{1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 0, 0, 0, 2},
		{0, 0, 0, 2, 2},
		{0, 0, 2, 2, 2},
	},
	CurrentPlayer: 1,
	stateFile:     nil,
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
	// Check that there are the right count of rows in the board.
	if len(board.Board) != BoardSize {
		return errors.New("invalid game state, please provide a valid game state")
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
			if cell > 0 {
				// There is a player on the current cell, increment its count.
				playerPawnsCounts[cell-1] += 1
			}
		}
	}

	// Check that there are enough pawns for a player in the board.
	if playerPawnsCounts[0] != PlayerPawnsNumber || playerPawnsCounts[1] != PlayerPawnsNumber {
		return errors.New("invalid game state, please provide a valid game state")
	}

	// No error.
	return nil
}

// Clone the board state.
func (board *BoardState) Clone() *BoardState {
	// Initialize a new board.
	clonedBoard := &BoardState{
		Board:         make([][]Cell, len(board.Board)),
		CurrentPlayer: board.CurrentPlayer,
		stateFile:     board.stateFile,
	}

	// Clone all rows of the board.
	for rowIndex, row := range board.Board {
		// Clone the current row.
		clonedBoard.Board[rowIndex] = make([]Cell, len(row))
		copy(clonedBoard.Board[rowIndex], row)
	}

	return clonedBoard
}

// Check that the provided move is legal.
// A move is illegal when the pawn only moves to an adjacent cell and not further.
func (board *BoardState) CheckMoveLegality(from CellIdentifier, to CellIdentifier) error {
	// Compute the column diff of the move.
	columnDiff := math.Abs(float64(from.Column - to.Column))
	// Compute the row diff of the move.
	rowDiff := math.Abs(float64(from.Row - to.Row))

	if columnDiff+rowDiff == 1 {
		// Only 1 difference, the move is legal.
		return nil
	}

	// The move is illegal (more than 1 difference, or no difference).

	if rowDiff == 1 && columnDiff == 1 {
		// Detected a diagonal move, return a specific error.
		return errors.New("a pawn cannot move in diagonal")
	}

	return fmt.Errorf("'%s' cannot be reached from '%s'", to.String(), from.String())
}

// Check legality of all successive moves.
func (board *BoardState) CheckMovesLegality(moveList []CellIdentifier) error {
	if len(moveList) == 2 {
		// Only 2 positions in the list = only one move, just check its legality.
		return board.CheckMoveLegality(moveList[0], moveList[1])
	} else {
		// More than 2 positions in the list, check the first move and the other moves recursively.
		if err := board.CheckMoveLegality(moveList[0], moveList[1]); err != nil {
			return err
		}
		// The first move is legal, check the others.
		return board.CheckMovesLegality(moveList[1:])
	}
}

// Move a pawn of the board.
func (board *BoardState) MovePawn(serializedMoveList string) error {
	// Parse the move list.
	moveList, err := board.ParseMoveList(serializedMoveList)
	if err != nil {
		return err
	}

	// Ensure that there is a pawn at start position.
	startPawn := board.Board[moveList[0].Row][moveList[0].Column]
	if startPawn == 0 {
		return fmt.Errorf("there is no pawn on %s", moveList[0].String())
	}

	// Ensure that there is no pawn at the end position.
	endPawn := board.Board[moveList[len(moveList)-1].Row][moveList[len(moveList)-1].Column]
	if endPawn != 0 {
		return fmt.Errorf("there already is a pawn on %s", moveList[len(moveList)-1].String())
	}

	// Check all successive moves legality.
	if err = board.CheckMovesLegality(moveList); err != nil {
		return err
	}
	// Check entire move legality before doing it.
	if err = board.CheckMoveLegality(moveList[0], moveList[len(moveList)-1]); err != nil {
		return err
	}

	// Move the start pawn to the end position.
	board.Board[moveList[len(moveList)-1].Row][moveList[len(moveList)-1].Column] = startPawn
	// Remove the start pawn from its previous position.
	board.Board[moveList[0].Row][moveList[0].Column] = 0

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

// Initialize a default board state.
func NewDefaultBoard() *BoardState {
	return DefaultBoard.Clone()
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
