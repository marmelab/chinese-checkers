package main

import (
	"fmt"
	"os"

	"github.com/go-color-term/go-color-term/coloring"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

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

				// Print the board with the moved pawn and the current score.
				board.Print(os.Stdout)
				board.PrintScore(os.Stdout)

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
	input := ""

	for {
		fmt.Print("\033[H\033[2J")

		println()

		// Print the current board.
		board.Print(os.Stdout)

		// Show the current score.
		board.PrintScore(os.Stdout)

		// Print the previous error if there is one.
		if len(errMsg) > 0 {
			println()

			// Print the previous input if there is one.
			if len(input) > 0 {
				println(coloring.Cyan("Tried to play \"" + input + "\""))
			}

			println(coloring.Red("Error: " + errMsg))
			errMsg = ""
		}

		println()

		// Prompt the current player for a new move list.
		fmt.Printf("%s to play, move a pawn (e.g. a2,a4): ", cases.Title(language.English).String(board.CurrentPlayer.Color()))
		input = ""
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
