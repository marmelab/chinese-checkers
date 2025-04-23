package main

import (
	"os"
)

func main() {
	if err := RunCli(); err != nil {
		os.Exit(1)
	}
}
