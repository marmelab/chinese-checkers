package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/marmelab/chinese-checkers/internal/game"
)

var gameStateFilePath string
var serializedMoveList string

// Initialize a game state by using the provided state file path if there is one.
func InitGameState() (*game.BoardState, error) {
	if len(gameStateFilePath) > 0 {
		return game.NewBoardFromStateFile(gameStateFilePath)
	} else {
		return game.NewDefaultBoard(), nil
	}
}

// Run the chinese checkers command line interface.
func RunCli() error {
	// Declare the main command.
	chineseCheckersCommand := &cobra.Command{
		Use:   "chinese-checkers",
		Short: "Fun chinese checkers game implementation",
		Long:  "A fun chinese checkers game implementation, to play with your sysadmin friends in a stealthy manner.",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Initialize the game state, from the state file if there is one.
			board, err := InitGameState()
			if err != nil {
				return err
			}

			if len(serializedMoveList) > 0 {
				// Move a pawn on the board and save if a state file has been provided.
				if err = board.MovePawnAndSave(serializedMoveList); err != nil {
					return err
				}

				// Print the board with the moved pawn.
				board.Print(os.Stdout)

				return nil
			} else {
				// Run the game loop.
				runGameLoop(board)
				return nil
			}
		},
	}

	// Add game state file flag without shorthand.
	chineseCheckersCommand.PersistentFlags().StringVarP(&gameStateFilePath, "state-file", "", "", "Game state file to read the board from.")
	// Add a required move flag without shorthand.
	chineseCheckersCommand.PersistentFlags().StringVarP(&serializedMoveList, "move", "m", "", "Move a pawn from a start position to an end position.")

	// Execute the command and return the error.
	return chineseCheckersCommand.Execute()
}

// Run the CLI infinite loop to interact with the game.
func runGameLoop(board *game.BoardState) {
	errMsg := ""

	for {
		// Print the current board.
		board.Print(os.Stdout)

		// Print the previous error if there is one.
		if len(errMsg) > 0 {
			println(errMsg)
			errMsg = ""
		}

		// Show the player.
		if board.CurrentPlayer == 1 {
			print("Green")
		} else {
			print("Red")
		}
		print(" to play, ")
		// Prompt the user for a new move list.
		print("move a pawn (e.g. a2,a4): ")
		var input string
		fmt.Scanln(&input)
		println()

		// The provided move is empty, exit the game.
		if len(input) == 0 {
			println("Bye bye!")
			break
		}

		// Try to move the pawn and store the error if there is one.
		if err := board.MovePawnAndSave(input); err != nil {
			errMsg = err.Error()
		}
	}
}
