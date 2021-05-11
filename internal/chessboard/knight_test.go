package chessboard

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetShortestKnightPath_HappyPath(t *testing.T) {
	testCases := []getShortestKnightPathTestCase{
		{
			input:    []string{"A1", "B2"},
			expected: []string{"B3", "A5", "C4", "B2"},
		},
		{
			input:    []string{"A8", "B7"},
			expected: []string{"B6", "C4", "D6", "B7"},
		},
		{
			input:    []string{"H8", "G7"},
			expected: []string{"G6", "E7", "F5", "G7"},
		},
		{
			input:    []string{"H1", "G2"},
			expected: []string{"F2", "D1", "E3", "G2"},
		},
		{
			input:    []string{"D5", "A1"},
			expected: []string{"B4", "C2", "A1"},
		},
		{
			input:    []string{"D5", "B7"},
			expected: []string{"B4", "D3", "C5", "B7"},
		},
		{
			input:    []string{"D5", "H8"},
			expected: []string{"F4", "G6", "H8"},
		},
		{
			input:    []string{"A1", "H8"},
			expected: []string{"B3", "C5", "E4", "G5", "F7", "H8"},
		},
	}

	for _, testCase := range testCases {
		input := testCase.input

		path := GetShortestKnightPath(input[0], input[1])

		if path == nil {
			t.Errorf("Should not return nil\n")
			continue
		}

		expected := strings.Join(testCase.expected, ", ")
		result := strings.Join(path, ", ")
		errorMsg := fmt.Sprintf("Result should be equals to expected: expected[%s], result[%s]", expected, result)

		assertEquals(expected, result, errorMsg, t)
	}
}

type getShortestKnightPathTestCase struct {
	input    []string
	expected []string
}
