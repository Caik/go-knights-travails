package chessboard

type CamelPiece struct {
}

func (c CamelPiece) getNextMoves(coordinates squareCoordinates) []squareCoordinates {
	possibleCoordinates := make([]squareCoordinates, 0)

	xIndex := coordinates.getXAsIndex()
	yIndex := coordinates.getYAsIndex()

	xPossibleCoord := []int{xIndex - 3, xIndex + 3}
	yPossibleCoord := []int{yIndex - 3, yIndex + 3}

	for _, x := range xPossibleCoord {
		if x < 0 || x > 7 {
			continue
		}

		yVariation := []int{yIndex - 1, yIndex + 1}

		for _, y := range yVariation {
			if y < 0 || y > 7 {
				continue
			}

			possibleCoordinates = append(possibleCoordinates, newSquareCoordinatesFromIndex(x, y))
		}
	}

	for _, y := range yPossibleCoord {
		if y < 0 || y > 7 {
			continue
		}

		xVariation := []int{xIndex - 1, xIndex + 1}

		for _, x := range xVariation {
			if x < 0 || x > 7 {
				continue
			}

			possibleCoordinates = append(possibleCoordinates, newSquareCoordinatesFromIndex(x, y))
		}
	}

	return possibleCoordinates
}

