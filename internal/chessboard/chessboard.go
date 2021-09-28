package chessboard

import (
	"fmt"
	"strings"
)

func GetShortestPath(start, end, piece string) ([]string, error) {
	chessPiece, err := GetChessPieceByName(piece)

	if err != nil {
		return nil, err
	}

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
			return strings.Split(current.path, " "), nil
		}

		next := chessPiece.getNextMoves(current.square)

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

	return []string{}, nil
}

type squarePath struct {
	square squareCoordinates
	path   string
}

func (s squarePath) String() string {
	return fmt.Sprintf("{%s [%s]}", s.square.String(), s.path)
}
