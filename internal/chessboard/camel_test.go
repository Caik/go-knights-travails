package chessboard

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestGetShortestCamelPath_HappyPath(t *testing.T) {
	testCases := []getShortestCamelPathTestCase{
		{
			input:          []string{"A1", "B2"},
			expectedLength: 3,
		},
		{
			input:          []string{"A8", "B7"},
			expectedLength: 3,
		},
		{
			input:          []string{"H8", "G7"},
			expectedLength: 3,
		},
		{
			input:          []string{"H1", "G2"},
			expectedLength: 3,
		},
		{
			input:          []string{"D5", "A1"},
			expectedLength: 0,
		},
		{
			input:          []string{"D5", "B7"},
			expectedLength: 2,
		},
		{
			input:          []string{"D5", "H8"},
			expectedLength: 0,
		},
		{
			input:          []string{"A1", "H8"},
			expectedLength: 5,
		},
	}

	for _, testCase := range testCases {
		input := testCase.input

		path, err := GetShortestPath(input[0], input[1], "camel")

		if err != nil {
			t.Errorf("Should not return an error: %v", err)
			continue
		}

		if path == nil {
			t.Errorf("Should not return nil")
			continue
		}

		expected := testCase.expectedLength
		result := len(path)
		errorMsg := fmt.Sprintf("Result should be equals in length: expected[%d], result[%d]", expected, result)

		assertEqualsInt(expected, result, errorMsg, t)

		if expected == 0 || result == 0 {
			continue
		}

		expectedLastSquare := testCase.input[1]
		resultLastSquare := path[len(path)-1]
		errorMsg = fmt.Sprintf("Result should be equals on last square: expected[%s], result[%s]", expectedLastSquare, resultLastSquare)

		assertEquals(expectedLastSquare, resultLastSquare, errorMsg, t)
	}
}

func TestGetPossibleCamelMovesSquares_HappyPath(t *testing.T) {
	testCases := []getPossibleCamelMovesSquaresTestCase{
		{
			square:          "A1",
			expectedSquares: []string{"B4", "D2"},
		},
		{
			square:          "A8",
			expectedSquares: []string{"B5", "D7"},
		},
		{
			square:          "H1",
			expectedSquares: []string{"E2", "G4"},
		},
		{
			square:          "H8",
			expectedSquares: []string{"E7", "G5"},
		},
		{
			square:          "D1",
			expectedSquares: []string{"A2", "C4", "E4", "G2"},
		},
		{
			square:          "D8",
			expectedSquares: []string{"A7", "C5", "E5", "G7"},
		},
		{
			square:          "A5",
			expectedSquares: []string{"B2", "B8", "D4", "D6"},
		},
		{
			square:          "H5",
			expectedSquares: []string{"E4", "E6", "G2", "G8"},
		},
		{
			square:          "D5",
			expectedSquares: []string{"A4", "A6", "C2", "C8", "E2", "E8", "G4", "G6"},
		},
	}

	camel := CamelPiece{}

	for _, testCase := range testCases {
		coordinates := newSquareCoordinates(extractSquareCoordinates(testCase.square))
		possibleCoordinates := camel.getNextMoves(coordinates)

		if possibleCoordinates == nil {
			t.Errorf("Should not return nil")
			continue
		}

		possibleSquares := make([]string, 0)

		for _, coordinate := range possibleCoordinates {
			possibleSquares = append(possibleSquares, coordinate.String())
		}

		sort.Strings(testCase.expectedSquares)
		sort.Strings(possibleSquares)

		expected := strings.Join(testCase.expectedSquares, ", ")
		result := strings.Join(possibleSquares, ", ")
		errorMsg := fmt.Sprintf("Result should be equals: expected[%s], result[%s]", expected, result)

		assertEquals(expected, result, errorMsg, t)
	}
}

type getShortestCamelPathTestCase struct {
	input          []string
	expectedLength int
}

type getPossibleCamelMovesSquaresTestCase struct {
	square          string
	expectedSquares []string
}
