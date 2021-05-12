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

func (s squareCoordinates) getXAsIndex() int {
	return int(s.x - 'A')
}

func (s squareCoordinates) getYAsIndex() int {
	return int(s.y - '1')
}

func (s squareCoordinates) String() string {
	return fmt.Sprintf("%s%s", string(s.x), string(s.y))
}

func newSquareCoordinates(x, y uint8) squareCoordinates {
	return squareCoordinates{
		x: x,
		y: y,
	}
}

func newSquareCoordinatesFromIndex(x, y int) squareCoordinates {
	return squareCoordinates{
		x: uint8(x + 'A'),
		y: uint8(y + '1'),
	}
}
