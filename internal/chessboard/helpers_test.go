package chessboard

import "testing"

func assertEquals(expected, obtained, errorMsg string, t *testing.T) {
	if expected != obtained {
		t.Errorf(errorMsg)
	}
}

func assertEqualsInt(expected, obtained int, errorMsg string, t *testing.T) {
	if expected != obtained {
		t.Errorf(errorMsg)
	}
}
