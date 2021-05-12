package chessboard

import (
	"fmt"
	"strings"
)

func GetShortestKnightPath(start, end string) []string {
	startSquare := newSquareCoordinates(extractSquareCoordinates(start))
	endSquare := newSquareCoordinates(extractSquareCoordinates(end))

	queue := []squarePath{{
		square: startSquare,
		path:   "",
	}}
	queueMap := map[string]bool{startSquare.String(): true}

	for i := 0; i < len(queue); i++ {
		current := queue[i]
		value := current.square.String()

		if value == endSquare.String() {
			return strings.Split(current.path, " ")
		}

		next := getPossibleKnightMovesSquares(current.square)

		for _, nextSquare := range next {
			nextValue := nextSquare.String()

			if _, ok := queueMap[nextValue]; ok {
				continue
			}

			path := strings.TrimSpace(fmt.Sprintf("%s %s", current.path, nextValue))

			queue = append(queue, squarePath{
				square: nextSquare,
				path:   path,
			})
			queueMap[nextValue] = true
		}
	}

	return []string{}
}

func getPossibleKnightMovesSquares(coordinates squareCoordinates) []squareCoordinates {
	possibleCoordinates := make([]squareCoordinates, 0)

	xIndex := coordinates.getXAsIndex()
	yIndex := coordinates.getYAsIndex()

	xPossibleCoord := []int{xIndex - 2, xIndex + 2}
	yPossibleCoord := []int{yIndex - 2, yIndex + 2}

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

type squarePath struct {
	square squareCoordinates
	path   string
}

func (s squarePath) String() string {
	return fmt.Sprintf("{%s [%s]}", s.square.String(), s.path)
}
