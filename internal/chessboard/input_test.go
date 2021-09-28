package chessboard

import (
	"fmt"
	"os"
	"testing"
)

func TestValidateInput_HappyPath(t *testing.T) {
	if err := ValidateInput([]string{"A1", "B2", "knight"}); err != nil {
		t.Errorf("should not return nil")
	}

	if err := ValidateInput([]string{"A1", "B2", "camel"}); err != nil {
		t.Errorf("should not return nil")
	}
}

func TestValidateInput_NilParameters(t *testing.T) {
	err := ValidateInput(nil)

	if err == nil {
		t.Errorf("should return nil")
		return
	}

	expected := fmt.Sprintf("usage: %s A1 B2 %s", os.Args[0], GetAvailablePiecesName())
	errorMsg := fmt.Sprintf("error should be equals to %q", expected)

	assertEquals(expected, err.Error(), errorMsg, t)
}

func TestValidateInput_EmptyParameters(t *testing.T) {
	err := ValidateInput([]string{})

	if err == nil {
		t.Errorf("should return nil")
		return
	}

	expected := fmt.Sprintf("usage: %s A1 B2 %s", os.Args[0], GetAvailablePiecesName())
	errorMsg := fmt.Sprintf("error should be equals to %q", expected)

	assertEquals(expected, err.Error(), errorMsg, t)
}

func TestValidateInput_SingleParameter(t *testing.T) {
	err := ValidateInput([]string{"A1"})

	if err == nil {
		t.Errorf("should return nil")
		return
	}

	expected := fmt.Sprintf("usage: %s A1 B2 %s", os.Args[0], GetAvailablePiecesName())
	errorMsg := fmt.Sprintf("error should be equals to %q", expected)

	assertEquals(expected, err.Error(), errorMsg, t)
}

func TestValidateInput_MoreThanThreeParameters(t *testing.T) {
	err := ValidateInput([]string{"A1", "B2", "camel", "C3"})

	if err == nil {
		t.Errorf("should return nil")
		return
	}

	expected := fmt.Sprintf("usage: %s A1 B2 %s", os.Args[0], GetAvailablePiecesName())
	errorMsg := fmt.Sprintf("error should be equals to %q", expected)

	assertEquals(expected, err.Error(), errorMsg, t)
}

func TestValidateInput_ParametersEquals(t *testing.T) {
	err := ValidateInput([]string{"A1", "A1"})

	if err == nil {
		t.Errorf("should return nil")
		return
	}

	expected := fmt.Sprintf("usage: %s A1 B2 %s", os.Args[0], GetAvailablePiecesName())
	errorMsg := fmt.Sprintf("error should be equals to %q", expected)

	assertEquals(expected, err.Error(), errorMsg, t)
}

func TestValidateInput_FirstParameterInvalid(t *testing.T) {
	testCases := [][]string{
		{"", "B2", "knight"},
		{"1", "B2", "knight"},
		{"A", "B2", "knight"},
		{"11", "B2", "knight"},
		{"AA", "B2", "knight"},
		{"I1", "B2", "knight"},
		{"A9", "B2", "knight"},
		{"A0", "B2", "knight"},
		{"", "B2", "camel"},
		{"1", "B2", "camel"},
		{"A", "B2", "camel"},
		{"11", "B2", "camel"},
		{"AA", "B2", "camel"},
		{"I1", "B2", "camel"},
		{"A9", "B2", "camel"},
		{"A0", "B2", "camel"},
	}

	for _, testCase := range testCases {
		err := ValidateInput(testCase)

		if err == nil {
			t.Errorf("should return nil")
			continue
		}

		expected := fmt.Sprintf("invalid input: invalid parameter '%s'", testCase[0])
		errorMsg := fmt.Sprintf("error should be equals to %q", expected)

		assertEquals(expected, err.Error(), errorMsg, t)
	}
}

func TestValidateInput_SecondParameterInvalid(t *testing.T) {
	testCases := [][]string{
		{"A1", "", "knight"},
		{"A1", "2", "knight"},
		{"A1", "B", "knight"},
		{"A1", "22", "knight"},
		{"A1", "BB", "knight"},
		{"A1", "I2", "knight"},
		{"A1", "B9", "knight"},
		{"A1", "B0", "knight"},
		{"A1", "", "camel"},
		{"A1", "2", "camel"},
		{"A1", "B", "camel"},
		{"A1", "22", "camel"},
		{"A1", "BB", "camel"},
		{"A1", "I2", "camel"},
		{"A1", "B9", "camel"},
		{"A1", "B0", "camel"},
	}

	for _, testCase := range testCases {
		err := ValidateInput(testCase)

		if err == nil {
			t.Errorf("should return nil")
			continue
		}

		expected := fmt.Sprintf("invalid input: invalid parameter '%s'", testCase[1])
		errorMsg := fmt.Sprintf("error should be equals to %q", expected)

		assertEquals(expected, err.Error(), errorMsg, t)
	}
}
