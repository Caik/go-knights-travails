package chessboard

import "fmt"

type squareCoordinates struct {
	x uint8
	y uint8
}

func (s squareCoordinates) getX() uint8 {
	return s.x
}

func (s squareCoordinates) getY() uint8 {
	return s.y
}

func (s squareCoordinates) getXAsIndex() uint8 {
	return s.x - 'A'
}

func (s squareCoordinates) getYAsIndex() uint8 {
	return s.y - '1'
}

func (s squareCoordinates) getAsString() string {
	return fmt.Sprintf("%s%s\n", string(s.x), string(s.y))
}
