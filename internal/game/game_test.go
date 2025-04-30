package game

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ongoingGameStateTestPath = "../../tests/states/ongoing-game.json"
const invalidGameStateTestPath = "../../tests/states/invalid-board.json"
const invalidPawnsCountTestPath = "../../tests/states/invalid-pawns-count.json"
const invalidCurrentPlayerTestPath = "../../tests/states/invalid-current-player.json"
const invalidPlayerInBoardTestPath = "../../tests/states/invalid-player-in-board.json"
const ongoing7x7GameStateTestPath = "../../tests/states/7x7-ongoing-game.json"
const invalidPawnsCount7x7TestPath = "../../tests/states/7x7-invalid-pawns-count.json"

func TestLoadBoard(t *testing.T) {
	stateFilePath := ongoingGameStateTestPath
	expected := &BoardState{
		Board: [][]Cell{
			{0, 1, 1, 0, 0},
			{0, 1, 0, 0, 0},
			{1, 1, 0, 2, 0},
			{1, 0, 2, 2, 2},
			{0, 0, 0, 2, 2},
		},
		CurrentPlayer:  Red,
		stateFile:      &stateFilePath,
		gameDefinition: &gameDefinitions[0],
	}
	board, err := NewBoardFromStateFile(ongoingGameStateTestPath)

	assert.Nil(t, err)
	assert.Equal(t, expected, board, "should be an ongoing game board")
}

func TestLoadInvalidBoard(t *testing.T) {
	board, err := NewBoardFromStateFile(invalidGameStateTestPath)

	assert.Nil(t, board)
	assert.Equal(t, "invalid game state, please provide a valid game state", err.Error(), "should return an invalid game state error")
}

func TestLoadBoardWithInvalidPawnsCount(t *testing.T) {
	board, err := NewBoardFromStateFile(invalidPawnsCountTestPath)

	assert.Nil(t, board)
	assert.Equal(t, "invalid game state, please provide a valid game state", err.Error(), "should return an invalid game state error")
}

func TestLoadBoardWithInvalidCurrentPlayer(t *testing.T) {
	board, err := NewBoardFromStateFile(invalidCurrentPlayerTestPath)

	assert.Nil(t, board)
	assert.Equal(t, "4 is not a valid player ID", err.Error(), "should return an invalid player error")
}

func TestLoadBoardWithInvalidPlayerInBoard(t *testing.T) {
	board, err := NewBoardFromStateFile(invalidPlayerInBoardTestPath)

	assert.Nil(t, board)
	assert.Equal(t, "3 is not a valid player ID", err.Error(), "should return an invalid player error")
}

func TestLoadBoardFromString(t *testing.T) {
	expected := &BoardState{
		Board: [][]Cell{
			{0, 1, 1, 0, 0},
			{0, 1, 0, 0, 0},
			{1, 1, 0, 2, 0},
			{1, 0, 2, 2, 2},
			{0, 0, 0, 2, 2},
		},
		CurrentPlayer:  Red,
		gameDefinition: &gameDefinitions[0],
	}

	board, err := NewBoardFromState([]byte(`
{
  "board": [
    [0, 1, 1, 0, 0],
    [0, 1, 0, 0, 0],
    [1, 1, 0, 2, 0],
    [1, 0, 2, 2, 2],
    [0, 0, 0, 2, 2]
  ],
  "currentPlayer": 2
}
`))

	assert.Nil(t, err)
	assert.Equal(t, expected, board, "should be an ongoing game board")
}

func TestNewDefaultBoard(t *testing.T) {
	expected := &DefaultBoard5
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure equality.
	assert.Equal(t, expected, board, "should be the default board")
}

func TestBoardCloning(t *testing.T) {
	// Test to clone the default dashboard.
	clonedBoard := DefaultBoard5.Clone()
	assert.Equal(t, &DefaultBoard5, clonedBoard, "should be the same board")
	assert.NotSame(t, &DefaultBoard5, clonedBoard, "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard5.Board, &clonedBoard.Board, "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard5.Board[0], &clonedBoard.Board[0], "shouldn't be the same pointer")
	assert.NotSame(t, DefaultBoard5.gameDefinition, clonedBoard.gameDefinition, "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard5.gameDefinition.GreenTargetAreaMask[0], &clonedBoard.gameDefinition.GreenTargetAreaMask[0], "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard5.gameDefinition.RedTargetAreaMask[0], &clonedBoard.gameDefinition.RedTargetAreaMask[0], "shouldn't be the same pointer")

	// Test to clone a loaded dashboard.
	board, err := NewBoardFromStateFile(ongoingGameStateTestPath)
	assert.Nil(t, err)
	clonedBoard = board.Clone()
	assert.Equal(t, board, clonedBoard, "should be the same board")
	assert.NotSame(t, &DefaultBoard5, clonedBoard, "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard5.Board, &clonedBoard.Board, "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard5.Board[0], &clonedBoard.Board[0], "shouldn't be the same pointer")
	assert.NotSame(t, DefaultBoard5.gameDefinition, clonedBoard.gameDefinition, "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard5.gameDefinition.GreenTargetAreaMask[0], &clonedBoard.gameDefinition.GreenTargetAreaMask[0], "shouldn't be the same pointer")
	assert.NotSame(t, &DefaultBoard5.gameDefinition.RedTargetAreaMask[0], &clonedBoard.gameDefinition.RedTargetAreaMask[0], "shouldn't be the same pointer")
}

func TestPreviousPlayer(t *testing.T) {
	board := NewDefaultBoard7()

	board.CurrentPlayer = Red
	assert.Equal(t, Green, board.GetPreviousPlayer(), "the previous player must be Green when the current player is Red")

	board.CurrentPlayer = Green
	assert.Equal(t, Red, board.GetPreviousPlayer(), "the previous player must be Red when the current player is Green")
}

func TestMovePawnInDefaultBoard(t *testing.T) {
	expected := &BoardState{
		Board: [][]Cell{
			{1, 1, 1, 0, 0},
			{1, 1, 0, 0, 0},
			{0, 1, 0, 0, 2},
			{0, 0, 0, 2, 2},
			{0, 0, 2, 2, 2},
		},
		CurrentPlayer:  Red,
		gameDefinition: &gameDefinitions[0],
		LastMove:       []CellIdentifier{{2, 0}, {2, 1}},
	}

	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure equality.
	err := board.MovePawn("c1,c2")
	assert.Nil(t, err)
	assert.Equal(t, expected, board)
}

func TestMovePawnInOngoingGameBoard(t *testing.T) {
	stateFilePath := ongoingGameStateTestPath
	expected := &BoardState{
		Board: [][]Cell{
			{0, 1, 1, 0, 0},
			{0, 1, 0, 0, 0},
			{1, 1, 0, 2, 0},
			{1, 0, 2, 2, 2},
			{0, 0, 2, 0, 2},
		},
		CurrentPlayer:  Green,
		stateFile:      &stateFilePath,
		gameDefinition: &gameDefinitions[0],
		LastMove:       []CellIdentifier{{4, 3}, {4, 2}},
	}

	board, err := NewBoardFromStateFile(ongoingGameStateTestPath)
	assert.Nil(t, err)
	err = board.MovePawn("e4,e3")
	assert.Nil(t, err)
	assert.Equal(t, expected, board)
}

func TestMovePawnWithNoPawnOnStartPosition(t *testing.T) {
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure equality and validity.
	err := board.MovePawn("e1,c1,c2,c3")
	assert.Equal(t, "there is no pawn on e1", err.Error(), "should return an error with no pawn on start position")
	assert.Equal(t, &DefaultBoard5, board, "should be unchanged")
}

func TestMovePawnWithAPawnOnEndPosition(t *testing.T) {
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure equality and validity.
	err := board.MovePawn("a1,b1")
	assert.Equal(t, "there is already a pawn on b1", err.Error(), "should return an error with a pawn on end position")
	assert.Equal(t, &DefaultBoard5, board, "should be unchanged")
}

func TestMovePawnErrorWithOnlyOneCell(t *testing.T) {
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure equality and validity.
	err := board.MovePawn("a1")
	assert.Equal(t, "you must provide at least two cells in a move", err.Error(), "should be a too small move list error")
	assert.Equal(t, &DefaultBoard5, board, "should be unchanged")
}

func TestMovePawnErrorWithInvalidFormat(t *testing.T) {
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure equality and validity.
	err := board.MovePawn("a,b2")
	assert.Equal(t, "invalid cell format 'a'", err.Error(), "should be an invalid format error")
	assert.Equal(t, &DefaultBoard5, board, "should be unchanged")
}

func TestMovePawnWithInvalidPosition(t *testing.T) {
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure equality and validity.

	err := board.MovePawn("c1,d1,e1,f1")
	assert.Equal(t, "f1 is not a valid cell", err.Error(), "should be an invalid cell error")
	assert.Equal(t, &DefaultBoard5, board, "should be unchanged")

	err = board.MovePawn("c1,d1,e1,f1,f2,e2")
	assert.Equal(t, "f1 is not a valid cell", err.Error(), "should be an invalid cell error")
	assert.Equal(t, &DefaultBoard5, board, "should be unchanged")

	err = board.MovePawn("91,d1,e1")
	assert.Equal(t, "91 is not a valid cell", err.Error(), "should be an invalid cell error")
	assert.Equal(t, &DefaultBoard5, board, "should be unchanged")
}

func TestSaveBoardState(t *testing.T) {
	expectedBoard := NewDefaultBoard5()
	expectedBoard.CurrentPlayer = Green // Set current player to ensure equality and validity.
	testFilePath := "testFile.json"
	expectedBoard.stateFile = &testFilePath

	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure equality and validity.
	assert.Nil(t, board.SaveState(testFilePath))
	loadedBoard, err := NewBoardFromStateFile(testFilePath)
	assert.Nil(t, err)
	assert.Equal(t, expectedBoard, loadedBoard, "saved board must be the default board")

	os.Remove(testFilePath)
}

func TestMovePawnAndSave(t *testing.T) {
	expectedBoard := NewDefaultBoard5()
	expectedBoard.CurrentPlayer = Green // Ensure the player to start is Green.
	expectedBoard.MovePawn("c1,d1")
	testFilePath := "testFile.json"
	expectedBoard.stateFile = &testFilePath

	{ // Prepare the state file.
		board := NewDefaultBoard5()
		board.CurrentPlayer = Green // Ensure the player to start is Green.
		assert.Nil(t, board.SaveState(testFilePath))
	}

	{ // Load a board from a state file, then move a pawn and save.
		board, err := NewBoardFromStateFile(testFilePath)
		assert.Nil(t, err)
		board.MovePawnAndSave("c1,d1")
	}

	{ // The file should now have been updated with the expected board state.
		board, err := NewBoardFromStateFile(testFilePath)
		assert.Nil(t, err)
		assert.Equal(t, expectedBoard, board, "saved board must be the default board with one moved pawn, and a saved last move")
	}

	os.Remove(testFilePath)
}

func TestMoveLegalityCheck(t *testing.T) {
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure equality and validity.

	a3, _ := board.ParseCellIdentifier("a3")
	a4, _ := board.ParseCellIdentifier("a4")
	_, err := board.CheckMoveLegality(*a3, *a4)
	assert.Nil(t, err, "the move is legal")

	a5, _ := board.ParseCellIdentifier("a5")
	_, err = board.CheckMoveLegality(*a3, *a5)
	assert.Equal(t, "'a5' cannot be reached from 'a3'", err.Error(), "should return an illegal move error")

	e1, _ := board.ParseCellIdentifier("e1")
	_, err = board.CheckMoveLegality(*a3, *e1)
	assert.Equal(t, "'e1' cannot be reached from 'a3'", err.Error(), "should return an illegal move error")

	b4, _ := board.ParseCellIdentifier("b4")
	_, err = board.CheckMoveLegality(*a3, *b4)
	assert.Equal(t, "a pawn cannot move in diagonal", err.Error(), "should return an illegal move error")
}

func TestMovesLegalityCheck(t *testing.T) {
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure equality and validity.

	{ // Legal move from a3 to a4.
		moveList, err := board.ParseMoveList("a3,a4")
		assert.Nil(t, err)
		assert.Nil(t, board.CheckMovesLegality(moveList, false))
	}
	{ // Legal moves from a3 to a4, then from a4 to a5, but only one move is allowed.
		moveList, err := board.ParseMoveList("a3,a4,a5")
		assert.Nil(t, err)
		assert.Equal(t, "cannot continue moving after moving to an adjacent cell", board.CheckMovesLegality(moveList, false).Error(), "should disallow multiple simple moves")
	}

	{ // Legal moves from a3 to a4, from a4 to b4, but illegal move from b4 to a5.
		assert.Nil(t, board.MovePawn("a3,a4"))
		board.CurrentPlayer = Green
		assert.Nil(t, board.MovePawn("a4,b4"))
		board.CurrentPlayer = Green
		moveList, err := board.ParseMoveList("b4,a5")
		assert.Nil(t, err)
		assert.Equal(t, "a pawn cannot move in diagonal", board.CheckMovesLegality(moveList, false).Error(), "should return an illegal move error")

		// Move back to a3
		assert.Nil(t, board.MovePawn("b4,a4"))
		board.CurrentPlayer = Green
		assert.Nil(t, board.MovePawn("a4,a3"))
		board.CurrentPlayer = Green
	}
	{ // Legal move from a3 to a4, but illegal move from a4 to c4.
		assert.Nil(t, board.MovePawn("a3,a4"))
		board.CurrentPlayer = Green
		moveList, err := board.ParseMoveList("a4,c4,a4")
		assert.Nil(t, err)
		assert.Equal(t, "'c4' cannot be reached from 'a4'", board.CheckMovesLegality(moveList, false).Error(), "should return an illegal move error")
	}
}

func TestIllegalMoves(t *testing.T) {
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure validity.

	assert.Equal(t, "'a5' cannot be reached from 'a3'", board.MovePawn("a3,a5").Error(), "should return an illegal move error")
	assert.Nil(t, board.MovePawn("a3,a4"), "the move is legal")
	board.CurrentPlayer = Green // Reset current player to ensure validity.
	assert.Equal(t, "a pawn cannot move in diagonal", board.MovePawn("a4,b3").Error(), "should return an illegal move error")
	assert.Equal(t, "'e1' cannot be reached from 'a4'", board.MovePawn("a4,e1").Error(), "should return an illegal move error")
	assert.Equal(t, "'c3' cannot be reached from 'c1'", board.MovePawn("c1,c3").Error(), "should return an illegal move error")
	assert.Equal(t, "a pawn cannot move in diagonal", board.MovePawn("a4,b3").Error(), "should return an illegal move error")
	assert.Equal(t, "'c2' cannot be reached from 'a4'", board.MovePawn("a4,c2").Error(), "should return an illegal move error")
	assert.Equal(t, "'c3' cannot be reached from 'a4'", board.MovePawn("a4,c3").Error(), "should return an illegal move error")
}

func TestPlayersTurns(t *testing.T) {
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure validity.

	assert.Nil(t, board.MovePawn("a3,a4"), "the move should be allowed")
	assert.Equal(t, Red, board.CurrentPlayer, "red player should now be the one to play")
	assert.Equal(t, "you cannot move a green pawn", board.MovePawn("a4,a5").Error(), "red player shouldn't be allowed to move a green pawn")
	assert.Nil(t, board.MovePawn("e3,e2"), "the move should be allowed")
	assert.Equal(t, Green, board.CurrentPlayer, "green player should now be the one to play")
	assert.Equal(t, "you cannot move a red pawn", board.MovePawn("e2,d2").Error(), "green player shouldn't be allowed to move a red pawn")
}

func TestJumpMoves(t *testing.T) {
	board := NewDefaultBoard5()
	board.CurrentPlayer = Green // Set current player to ensure validity.

	// Disallowed jumps of more than one pawn.
	assert.Equal(t, "a pawn cannot jump over more than one pawn", board.MovePawn("a1,d1").Error(), "should return an illegal move error")
	assert.Equal(t, "a pawn cannot jump over more than one pawn", board.MovePawn("a1,e1").Error(), "should return an illegal move error")
	assert.Equal(t, "a pawn cannot jump over more than one pawn", board.MovePawn("a1,a4").Error(), "should return an illegal move error")
	assert.Equal(t, "a pawn cannot jump over more than one pawn", board.MovePawn("a1,a5").Error(), "should return an illegal move error")

	// Jump on another pawn.
	assert.Equal(t, "there is already a pawn on a3", board.MovePawn("a1,a3").Error(), "should return that the cell is not empty")
	assert.Equal(t, "there is already a pawn on a3", board.MovePawn("a1,a3,a5").Error(), "should return that an intermediate cell is not empty")

	// Valid simple jumps.
	assert.Nil(t, board.MovePawn("a2,a4"), "the move should be allowed")
	assert.Nil(t, board.MovePawn("d5,d3"), "the move should be allowed")
	board.CurrentPlayer = Red
	assert.Nil(t, board.MovePawn("e3,c3"), "the move should be allowed")
	board.CurrentPlayer = Red
	assert.Nil(t, board.MovePawn("d3,e3"), "the move should be allowed")
	board.CurrentPlayer = Red
	assert.Equal(t, "'d2' cannot be reached from 'd4'", board.MovePawn("d4,d2").Error(), "should return an illegal move error")
	board.CurrentPlayer = Green

	// Chained jumps.
	assert.Nil(t, board.MovePawn("a4,a2,c2"), "the chained jumps should be allowed")
	board.CurrentPlayer = Green
	assert.Nil(t, board.MovePawn("c2,a2,a4"), "the chained jumps should be allowed")
	board.CurrentPlayer = Green

	// Block a2 position with a1.
	assert.Nil(t, board.MovePawn("a1,a2"), "the move should be allowed")
	board.CurrentPlayer = Green
	// Error in chained jumps, as a2 is not empty.
	assert.Equal(t, "there is already a pawn on a2", board.MovePawn("a4,a2,c2").Error(), "should return that an intermediate cell is not empty")
	// Free a2 position again.
	assert.Nil(t, board.MovePawn("a2,a1"), "the move should be allowed")
	board.CurrentPlayer = Green

	// Simple move.
	assert.Nil(t, board.MovePawn("b2,b3"), "the move should be allowed")
	board.CurrentPlayer = Green
	// Error in chained jumps, as b2 has been moved to b3.
	assert.Equal(t, "'c2' cannot be reached from 'a2'", board.MovePawn("a4,a2,c2").Error(), "should return an illegal move error")

	// Diagonal jump is disallowed.
	assert.Equal(t, "'c2' cannot be reached from 'a4'", board.MovePawn("a4,c2").Error(), "should return an illegal move error")
}

func TestPawnsInTargetArea(t *testing.T) {
	{
		board := NewDefaultBoard5()
		greenPawns, redPawns := board.CountPawnsInTargetAreas()
		assert.Equal(t, int8(0), greenPawns, "should have no green pawn in target area")
		assert.Equal(t, int8(0), redPawns, "should have no red pawn in target area")
	}

	{
		board := NewDefaultBoard5()
		board.Board = [][]Cell{
			{2, 2, 2, 0, 0},
			{2, 2, 0, 0, 0},
			{2, 0, 0, 0, 1},
			{0, 0, 0, 1, 1},
			{0, 0, 1, 1, 1},
		}
		greenPawns, redPawns := board.CountPawnsInTargetAreas()
		assert.Equal(t, int8(6), greenPawns, "should have 6 green pawns in target area")
		assert.Equal(t, int8(6), redPawns, "should have 6 red pawns in target area")
	}

	{
		board := NewDefaultBoard5()
		board.Board = [][]Cell{
			{2, 0, 0, 0, 2},
			{0, 0, 0, 2, 0},
			{0, 0, 2, 1, 1},
			{0, 2, 1, 0, 1},
			{2, 0, 1, 0, 1},
		}
		greenPawns, redPawns := board.CountPawnsInTargetAreas()
		assert.Equal(t, int8(4), greenPawns, "should have 4 green pawns in target area")
		assert.Equal(t, int8(1), redPawns, "should have 1 red pawn in target area")
	}
}

func TestWinner(t *testing.T) {
	{
		board := NewDefaultBoard5()
		board.Board = [][]Cell{
			{2, 0, 0, 0, 2},
			{0, 0, 0, 2, 0},
			{0, 0, 2, 1, 1},
			{0, 2, 1, 0, 1},
			{2, 0, 1, 0, 1},
		}
		assert.Equal(t, None, board.GetWinner(), "should have no winner")

		// Can move a pawn.
		board.CurrentPlayer = Green
		assert.Nil(t, board.MovePawn("e3,e4"), "should be able to move a pawn")
	}

	{
		board := NewDefaultBoard5()
		board.Board = [][]Cell{
			{2, 2, 2, 0, 2},
			{2, 2, 0, 1, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 1},
			{0, 0, 1, 1, 1},
		}
		assert.Equal(t, None, board.GetWinner(), "should have no winner")

		// Can move a pawn.
		board.CurrentPlayer = Green
		assert.Nil(t, board.MovePawn("b4,b5"), "should be able to move a pawn")
	}

	{
		board := NewDefaultBoard5()
		board.Board = [][]Cell{
			{2, 0, 2, 0, 2},
			{0, 2, 0, 0, 0},
			{0, 0, 0, 0, 1},
			{0, 2, 0, 1, 1},
			{2, 0, 1, 1, 1},
		}
		assert.Equal(t, Green, board.GetWinner(), "should have green as a winner")

		// Cannot move a pawn.
		board.CurrentPlayer = Red
		assert.Equal(t, "cannot move a pawn: Green has won", board.MovePawn("a5,a4").Error(), "shouldn't be able to move a pawn")
	}

	{
		board := NewDefaultBoard5()
		board.Board = [][]Cell{
			{2, 2, 2, 0, 0},
			{2, 2, 0, 1, 0},
			{2, 0, 0, 0, 0},
			{0, 0, 0, 1, 1},
			{0, 0, 1, 1, 1},
		}
		assert.Equal(t, Red, board.GetWinner(), "should have red as a winner")

		// Cannot move a pawn.
		board.CurrentPlayer = Green
		assert.Equal(t, "cannot move a pawn: Red has won", board.MovePawn("b4,b5").Error(), "shouldn't be able to move a pawn")
	}
}

func Test7x7(t *testing.T) {
	{ // Test new default 7x7 board.
		board := NewDefaultBoard7()

		assert.Equal(t, [][]Cell{
			{1, 1, 1, 1, 0, 0, 0},
			{1, 1, 1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 2},
			{0, 0, 0, 0, 0, 2, 2},
			{0, 0, 0, 0, 2, 2, 2},
			{0, 0, 0, 2, 2, 2, 2},
		}, board.Board, "should be the default board")
		assert.Equal(t, gameDefinitions[1], *board.gameDefinition, "should be a 7x7 board")
	}

	{ // Test loading a 7x7 board from a file.
		board, err := NewBoardFromStateFile(ongoing7x7GameStateTestPath)
		assert.Nil(t, err, "should have loaded the state file without error")

		assert.Equal(t, [][]Cell{
			{0, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0},
			{1, 0, 1, 0, 0, 0, 2},
			{0, 0, 0, 0, 0, 2, 2},
			{0, 0, 2, 2, 2, 2, 0},
			{0, 0, 0, 2, 0, 2, 2},
		}, board.Board, "should be the default board")
		assert.Equal(t, gameDefinitions[1], *board.gameDefinition, "should be a 7x7 board")
		assert.Equal(t, Red, board.CurrentPlayer, "red should be the next to play")
	}

	{ // Test loading a 7x7 board with an invalid pawns count.
		_, err := NewBoardFromStateFile(invalidPawnsCount7x7TestPath)
		assert.Equal(t, "invalid game state, please provide a valid game state", err.Error(), "shouldn't load the board successfully")
	}

	{ // Test moving in a 7x7 board.
		board := NewDefaultBoard7()
		board.Board = [][]Cell{
			{1, 1, 1, 0, 0, 0, 1},
			{1, 1, 0, 0, 0, 1, 2},
			{1, 1, 2, 0, 0, 0, 0},
			{0, 0, 0, 2, 0, 1, 0},
			{0, 0, 0, 0, 2, 0, 2},
			{0, 0, 0, 0, 0, 2, 2},
			{0, 0, 0, 0, 2, 2, 2},
		}

		// Try a lot of jumps.
		board.CurrentPlayer = Green
		assert.Nil(t, board.MovePawn("c2,c4,e4,e6,c6,a6"), "should be able to move a pawn")

		assert.Equal(t, [][]Cell{
			{1, 1, 1, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 1, 2},
			{1, 0, 2, 0, 0, 0, 0},
			{0, 0, 0, 2, 0, 1, 0},
			{0, 0, 0, 0, 2, 0, 2},
			{0, 0, 0, 0, 0, 2, 2},
			{0, 0, 0, 0, 2, 2, 2},
		}, board.Board, "the pawn has moved from c2 to c6")
		assert.Equal(t, Red, board.CurrentPlayer, "the new current player is red")

		assert.Nil(t, board.MovePawn("g5,g4"), "should be able to move a pawn")
		assert.Equal(t, [][]Cell{
			{1, 1, 1, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 1, 2},
			{1, 0, 2, 0, 0, 0, 0},
			{0, 0, 0, 2, 0, 1, 0},
			{0, 0, 0, 0, 2, 0, 2},
			{0, 0, 0, 0, 0, 2, 2},
			{0, 0, 0, 2, 0, 2, 2},
		}, board.Board, "the pawn has moved from g5 to g4")
		assert.Equal(t, Green, board.CurrentPlayer, "the new current player is green")
	}

	{ // Test no winner in 7x7 board.
		board := NewDefaultBoard7()
		board.Board = [][]Cell{
			{2, 2, 2, 2, 2, 0, 0},
			{2, 2, 2, 0, 0, 0, 0},
			{2, 2, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 1, 0, 1, 1},
			{0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 1, 1, 1},
		}
		assert.Equal(t, None, board.GetWinner(), "should have no winner")

		// Can move a pawn.
		board.CurrentPlayer = Green
		assert.Nil(t, board.MovePawn("e4,f4"), "should be able to move a pawn")
	}

	{ // Test winner in 7x7 board.
		board := NewDefaultBoard7()
		board.Board = [][]Cell{
			{2, 2, 2, 2, 2, 0, 0},
			{2, 2, 2, 0, 0, 0, 0},
			{2, 2, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0, 1, 1},
			{0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 1, 1, 1, 1},
		}
		assert.Equal(t, Green, board.GetWinner(), "should have green as a winner")

		// Cannot move a pawn.
		board.CurrentPlayer = Red
		assert.Equal(t, "cannot move a pawn: Green has won", board.MovePawn("a5,b5").Error(), "shouldn't be able to move a pawn")
	}
}
