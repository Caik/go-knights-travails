package chessboard

type ChessPiece interface {
	getNextMoves(coordinates squareCoordinates) []squareCoordinates
}
