package main

import (
	"fmt"
	"github.com/Caik/go-knights-travails/internal/chessboard"
	"os"
	"strings"
)

func main() {
	parameters := os.Args[1:]

	// Validating input parameters
	if err := chessboard.ValidateInput(parameters); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Calculating shortest path
	path := chessboard.GetShortestKnightPath(parameters[0], parameters[1])

	// Showing shortest path
	fmt.Println(strings.Join(path, " "))
}
