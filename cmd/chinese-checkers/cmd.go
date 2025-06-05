package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/go-color-term/go-color-term/coloring"
	"github.com/spf13/cobra"

	"github.com/marmelab/chinese-checkers/internal/game"
)

var gameStateFilePath string
var serializedMoveList string
var boardSizeIdentifier string
var botMode bool

// Initialize a game state by using the provided state file path if there is one.
func InitGameState() (*game.BoardState, error) {
	if len(gameStateFilePath) > 0 {
		return game.NewBoardFromStateFile(gameStateFilePath)
	} else {
		if len(boardSizeIdentifier) > 0 {
			// A board size identifier has been provided, create a board of the provided size.
			switch boardSizeIdentifier {
			case "5x5":
				return game.NewDefaultBoard5(), nil
			case "7x7":
				return game.NewDefaultBoard7(), nil
			default:
				return nil, errors.New("invalid board size. please choose a valid board size (--board-size 5x5 (default) or --board-size 7x7)")
			}
		} else {
			// By default, use 5x5.
			return game.NewDefaultBoard5(), nil
		}
	}
}

// Show the error message.
func ShowError(errMsg string) {
	println(coloring.Red("Error: " + errMsg))
	println()
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
				ShowError(err.Error())
				return err
			}

			if len(serializedMoveList) > 0 {
				// Move a pawn on the board and save if a state file has been provided.
				if err = board.MovePawnAndSave(serializedMoveList); err != nil {
					ShowError(err.Error())
					return err
				}

				// Print the board with the moved pawn and the current score.
				board.Print(os.Stdout)
				board.PrintScore(os.Stdout)
				board.PrintLastMove(os.Stdout)

				return nil
			} else {
				// Run the game loop.
				runGameLoop(board)

				return nil
			}
		},
		SilenceErrors: true,
	}

	// Add game state file flag without shorthand.
	chineseCheckersCommand.PersistentFlags().StringVarP(&gameStateFilePath, "state-file", "", "", "Game state file to read the board from.")
	// Add a move flag.
	chineseCheckersCommand.PersistentFlags().StringVarP(&serializedMoveList, "move", "m", "", "Move a pawn from a start position to an end position.")
	// Add a board size flag.
	chineseCheckersCommand.PersistentFlags().StringVarP(&boardSizeIdentifier, "board-size", "", "", "Set the board size to start with.")
	// Add a flag for watching a game played by robots.
	chineseCheckersCommand.PersistentFlags().BoolVarP(&botMode, "bots", "b", false, "The robot plays the game.")

	// Execute the command and return the error.
	return chineseCheckersCommand.Execute()
}

// Run the CLI infinite loop to interact with the game.
func runGameLoop(board *game.BoardState) {
	errMsg := ""
	input := ""

	for {
		// Clear the screen.
		fmt.Print("\033[H\033[2J")

		println()

		// Print the current board.
		board.Print(os.Stdout)

		// Show the current score.
		board.PrintScore(os.Stdout)

		// Print the last move.
		board.PrintLastMove(os.Stdout)

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

		// If there is a winner, show them and exit.
		if winner := board.GetWinner(); winner != game.None {
			fmt.Printf("%s won the game!\n", winner.ColoredName())

			println()

			print("Press enter to exit the game...")
			// Do not display entered characters on the screen.
			exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
			fmt.Scanln(&input)
			println()
			break
		} else {
			if botMode {
				input = game.MoveToString(board.FindBestMoveIn30s())
			} else {
				println()

				// Prompt the current player for a new move list.
				fmt.Printf("%s to play, move a pawn (e.g. a2,a4): ", board.CurrentPlayer.ColoredName())
				input = ""
				fmt.Scanln(&input)
				println()

				// The provided move is empty, exit the game.
				if len(input) == 0 {
					println("Bye bye!")
					break
				}
			}

			// Try to move the pawn and store the error if there is one.
			if err := board.MovePawnAndSave(input); err != nil {
				errMsg = err.Error()
			}
		}
	}
}
