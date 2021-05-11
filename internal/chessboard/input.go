package chessboard

import (
	"fmt"
	"os"
	"strings"
)

func ValidateInput(parameters []string) error {
	if len(parameters) != 2 || parameters[0] == parameters[1] {
		return fmt.Errorf("usage: %s A1 B2\n", os.Args[0])
	}

	for _, parameter := range parameters {
		if err := validateParameter(parameter); err != nil {
			return err
		}
	}

	return nil
}

func validateParameter(parameter string) error {
	errorMsg := "invalid input: invalid parameter '%s'\n"

	if len(parameter) != 2 {
		return fmt.Errorf(errorMsg, parameter)
	}

	letter, number := extractSquareCoordinates(parameter)

	if letter < 'A' || letter > 'H' || number < '1' || number > '8' {
		return fmt.Errorf(errorMsg, parameter)
	}

	return nil
}

func extractSquareCoordinates(parameter string) (uint8, uint8) {
	parameter = strings.ToUpper(parameter)
	letter := parameter[0]
	number := parameter[1]

	return letter, number
}
