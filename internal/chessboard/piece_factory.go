package chessboard

import (
	"fmt"
	"strings"
	"sync"
)

var (
	once       sync.Once
	pieces     map[string]ChessPiece
	piecesName []string
)

func GetChessPieceByName(name string) (ChessPiece, error) {
	ensureInit()

	piece, ok := pieces[name]

	if !ok {
		return nil, fmt.Errorf("invalid chess piece: %s", name)
	}

	return piece, nil
}

func GetAvailablePiecesName() string {
	ensureInit()

	return strings.Join(piecesName, "|")
}

func ensureInit() {
	once.Do(func() {
		pieces = map[string]ChessPiece{
			"knight": &KnightPiece{},
			"camel":  &CamelPiece{},
		}

		piecesName = make([]string, 0, len(pieces))

		for k := range pieces {
			piecesName = append(piecesName, k)
		}
	})
}
