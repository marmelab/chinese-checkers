package game

import (
	"encoding/json"
	"errors"
	"fmt"
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
	}

	// Clone all rows of the board.
	for rowIndex, row := range board.Board {
		// Clone the current row.
		clonedBoard.Board[rowIndex] = make([]Cell, len(row))
		copy(clonedBoard.Board[rowIndex], row)
	}

	return clonedBoard
}

// Check that the board position is valid.
func (board *BoardState) CheckBoardPositionValidity(position CellIdentifier) error {
	if
	// Check that the row index is valid.
	(position[0] >= int8(len(board.Board)) || position[0] < 0) ||
		// Check that the column index is valid.
		(position[1] >= int8(len(board.Board[position[0]])) || position[1] < 0) {
		return fmt.Errorf("%s is not a valid cell", position.String())
	}
	return nil
}

// Move a pawn of the board.
func (board *BoardState) MovePawn(serializedMoveList string) error {
	// Parse the move list.
	moveList, err := ParseMoveList(serializedMoveList)
	if err != nil {
		return err
	}

	// Ensure that the position is valid for the current board.
	if err = board.CheckBoardPositionValidity(moveList[0]); err != nil {
		return err
	}
	// Ensure that there is a pawn at start position.
	startPawn := board.Board[moveList[0][0]][moveList[0][1]]
	if startPawn == 0 {
		return fmt.Errorf("there is no pawn on %s", moveList[0].String())
	}

	// Ensure that the position is valid for the current board.
	if err = board.CheckBoardPositionValidity(moveList[len(moveList)-1]); err != nil {
		return err
	}
	// Ensure that there is no pawn at the end position.
	endPawn := board.Board[moveList[len(moveList)-1][0]][moveList[len(moveList)-1][1]]
	if endPawn != 0 {
		return fmt.Errorf("there already is a pawn on %s", moveList[len(moveList)-1].String())
	}

	// Move the start pawn to the end position.
	board.Board[moveList[len(moveList)-1][0]][moveList[len(moveList)-1][1]] = startPawn
	// Remove the start pawn from its previous position.
	board.Board[moveList[0][0]][moveList[0][1]] = 0

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
