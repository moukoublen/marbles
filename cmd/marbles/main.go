package main

import (
	"fmt"
	"github.com/moukoublen/marbles/internal/cli"
	"os"
)

func main() {
	commands, err := cli.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Marbles failed %v\n", err)
		os.Exit(1)
	}
	_ = commands
}
