package chessboard

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestGetShortestKnightPath_HappyPath(t *testing.T) {
	testCases := []getShortestKnightPathTestCase{
		{
			input:          []string{"A1", "B2"},
			expectedLength: 4,
		},
		{
			input:          []string{"A8", "B7"},
			expectedLength: 4,
		},
		{
			input:          []string{"H8", "G7"},
			expectedLength: 4,
		},
		{
			input:          []string{"H1", "G2"},
			expectedLength: 4,
		},
		{
			input:          []string{"D5", "A1"},
			expectedLength: 3,
		},
		{
			input:          []string{"D5", "B7"},
			expectedLength: 4,
		},
		{
			input:          []string{"D5", "H8"},
			expectedLength: 3,
		},
		{
			input:          []string{"A1", "H8"},
			expectedLength: 6,
		},
	}

	for _, testCase := range testCases {
		input := testCase.input

		path := GetShortestKnightPath(input[0], input[1])

		if path == nil {
			t.Errorf("Should not return nil\n")
			continue
		}

		expected := testCase.expectedLength
		result := len(path)
		errorMsg := fmt.Sprintf("Result should be equals in length: expected[%d], result[%d]", expected, result)

		assertEqualsInt(expected, result, errorMsg, t)

		expectedLastSquare := testCase.input[1]
		resultLastSquare := path[len(path)-1]
		errorMsg = fmt.Sprintf("Result should be equals on last square: expected[%s], result[%s]", expectedLastSquare, resultLastSquare)

		assertEqualsInt(expected, result, errorMsg, t)
	}
}

func TestGetPossibleKnightMovesSquares_HappyPath(t *testing.T) {
	testCases := []getPossibleKnightMovesSquaresTestCase{
		{
			square:          "A1",
			expectedSquares: []string{"C2", "B3"},
		},
		{
			square:          "A8",
			expectedSquares: []string{"B6", "C7"},
		},
		{
			square:          "H1",
			expectedSquares: []string{"F2", "G3"},
		},
		{
			square:          "H8",
			expectedSquares: []string{"G6", "F7"},
		},
		{
			square:          "D1",
			expectedSquares: []string{"B2", "F2", "C3", "E3"},
		},
		{
			square:          "D8",
			expectedSquares: []string{"B7", "F7", "C6", "E6"},
		},
		{
			square:          "A5",
			expectedSquares: []string{"C4", "C6", "B7", "B3"},
		},
		{
			square:          "H5",
			expectedSquares: []string{"F4", "F6", "G3", "G7"},
		},
		{
			square:          "D5",
			expectedSquares: []string{"B4", "B6", "F4", "F6", "C3", "E3", "C7", "E7"},
		},
	}

	for _, testCase := range testCases {
		coordinates := newSquareCoordinates(extractSquareCoordinates(testCase.square))
		possibleCoordinates := getPossibleKnightMovesSquares(coordinates)

		if possibleCoordinates == nil {
			t.Errorf("Should not return nil\n")
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
		errorMsg := fmt.Sprintf("Result should be equals in length: expected[%s], result[%s]", expected, result)

		assertEquals(expected, result, errorMsg, t)
	}
}

type getShortestKnightPathTestCase struct {
	input          []string
	expectedLength int
}

type getPossibleKnightMovesSquaresTestCase struct {
	square          string
	expectedSquares []string
}
