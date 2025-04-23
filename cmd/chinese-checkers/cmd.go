package main

import (
	"github.com/spf13/cobra"

	"github.com/marmelab/chinese-checkers/internal/game"
)

var gameStateFilePath string

func GetGameStateFilePath() *string {
	if len(gameStateFilePath) > 0 {
		return &gameStateFilePath
	} else {
		return nil
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
			board, err := game.InitBoard(GetGameStateFilePath())
			if err != nil {
				return err
			}
			game.PrintBoard(board)
			return nil
		},
	}

	// Add game state file flag without shorthand.
	chineseCheckersCommand.PersistentFlags().StringVarP(&gameStateFilePath, "state-file", "", "", "Game state file to read the board from.")

	// Execute the command and return the error.
	return chineseCheckersCommand.Execute()
}
