package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Caik/go-knights-travails/internal/chessboard"
)

func main() {
	parameters := os.Args[1:]

	// Validating input parameters
	if err := chessboard.ValidateInput(parameters); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Calculating shortest path
	path, err := chessboard.GetShortestPath(parameters[0], parameters[1], parameters[2])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Showing shortest path
	fmt.Println(strings.Join(path, " "))
}
