package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// The main chinese checkers command instance.
var chineseCheckersCommand = &cobra.Command{
	Use:   "chinese-checkers",
	Short: "Fun chinese checkers game implementation",
	Long:  "A fun chinese checkers game implementation, to play with your sysadmin friends in a stealthy manner.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to chinese checkers!")
	},
}
