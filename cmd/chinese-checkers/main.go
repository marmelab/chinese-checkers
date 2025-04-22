package main

import (
	"os"
)

func main() {
	if err := chineseCheckersCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
