package chessboard

import (
	"fmt"
	"testing"
)

func TestValidateInput_HappyPath(t *testing.T) {
	if err := ValidateInput([]string{"A1", "B2"}); err != nil {
		t.Errorf("Should not return nil\n")
	}
}

func TestValidateInput_NilParameters(t *testing.T) {
	err := ValidateInput(nil)

	if err == nil {
		t.Errorf("Should return nil\n")
		return
	}

	expected := fmt.Sprintf("Usage: %s A1 B2", KnightsTravailsAppName)
	errorMsg := fmt.Sprintf("Error should be equals to %q\n", expected)

	assertEquals(expected, err.Error(), errorMsg, t)
}

func TestValidateInput_EmptyParameters(t *testing.T) {
	err := ValidateInput([]string{})

	if err == nil {
		t.Errorf("Should return nil\n")
		return
	}

	expected := fmt.Sprintf("Usage: %s A1 B2", KnightsTravailsAppName)
	errorMsg := fmt.Sprintf("Error should be equals to %q\n", expected)

	assertEquals(expected, err.Error(), errorMsg, t)
}

func TestValidateInput_SingleParameter(t *testing.T) {
	err := ValidateInput([]string{"A1"})

	if err == nil {
		t.Errorf("Should return nil\n")
		return
	}

	expected := fmt.Sprintf("Usage: %s A1 B2", KnightsTravailsAppName)
	errorMsg := fmt.Sprintf("Error should be equals to %q\n", expected)

	assertEquals(expected, err.Error(), errorMsg, t)
}

func TestValidateInput_MoreThanTwoParameters(t *testing.T) {
	err := ValidateInput([]string{"A1", "B2", "C3"})

	if err == nil {
		t.Errorf("Should return nil\n")
		return
	}

	expected := fmt.Sprintf("Usage: %s A1 B2", KnightsTravailsAppName)
	errorMsg := fmt.Sprintf("Error should be equals to %q\n", expected)

	assertEquals(expected, err.Error(), errorMsg, t)
}

func TestValidateInput_ParametersEquals(t *testing.T) {
	err := ValidateInput([]string{"A1", "A1"})

	if err == nil {
		t.Errorf("Should return nil\n")
		return
	}

	expected := fmt.Sprintf("Usage: %s A1 B2", KnightsTravailsAppName)
	errorMsg := fmt.Sprintf("Error should be equals to %q\n", expected)

	assertEquals(expected, err.Error(), errorMsg, t)
}

func TestValidateInput_FirstParameterInvalid(t *testing.T) {
	testCases := [][]string{
		{"", "B2"},
		{"1", "B2"},
		{"A", "B2"},
		{"11", "B2"},
		{"AA", "B2"},
		{"I1", "B2"},
		{"A9", "B2"},
		{"A0", "B2"},
	}

	for _, testCase := range testCases {
		err := ValidateInput(testCase)

		if err == nil {
			t.Errorf("Should return nil\n")
			continue
		}

		expected := fmt.Sprintf("Invalid input: invalid parameter '%s'", testCase[0])
		errorMsg := fmt.Sprintf("Error should be equals to %q\n", expected)

		assertEquals(expected, err.Error(), errorMsg, t)
	}
}

func TestValidateInput_SecondParameterInvalid(t *testing.T) {
	testCases := [][]string{
		{"A1", ""},
		{"A1", "2"},
		{"A1", "B"},
		{"A1", "22"},
		{"A1", "BB"},
		{"A1", "I2"},
		{"A1", "B9"},
		{"A1", "B0"},
	}

	for _, testCase := range testCases {
		err := ValidateInput(testCase)

		if err == nil {
			t.Errorf("Should return nil\n")
			continue
		}

		expected := fmt.Sprintf("Invalid input: invalid parameter '%s'", testCase[0])
		errorMsg := fmt.Sprintf("Error should be equals to %q\n", expected)

		assertEquals(expected, err.Error(), errorMsg, t)
	}
}
