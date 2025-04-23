package main

import (
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

			// Move a pawn on the board.
			if err = board.MovePawn(serializedMoveList); err != nil {
				return err
			}

			// Print the board with the moved pawn.
			board.Print(os.Stdout)
			return nil
		},
	}

	// Add game state file flag without shorthand.
	chineseCheckersCommand.PersistentFlags().StringVarP(&gameStateFilePath, "state-file", "", "", "Game state file to read the board from.")
	// Add a required move flag without shorthand.
	chineseCheckersCommand.PersistentFlags().StringVarP(&serializedMoveList, "move", "m", "", "Move a pawn from a start position to an end position.")
	chineseCheckersCommand.MarkFlagRequired("move")

	// Execute the command and return the error.
	return chineseCheckersCommand.Execute()
}
