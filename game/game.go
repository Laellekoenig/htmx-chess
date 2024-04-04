package game

import (
	"fmt"
	"strings"
)

type Square struct {
	Num        int
	IsBlack    bool
	IsActive   bool
	Piece      *Piece
	Coordinate string
}

type Game struct {
	Squares             []*Square
	ActiveSquares       []*Square
	BlackToMove         bool
	WhiteMayCastleShort bool
	WhiteMayCastleLong  bool
	BlackMayCastleShort bool
	BlackMayCastleLong  bool
	EnPassant           *Square
	FiftyMoves          int
	Move                int
}

func NewGame() *Game {
	var squares []*Square

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			isBlack := true

			if i%2 == 0 && j%2 == 0 {
				isBlack = false
			} else if i%2 != 0 && j%2 != 0 {
				isBlack = false
			}

			coordinate := fmt.Sprintf("%c%d", 'a'+j, 8-i)
			newSquare := &Square{Num: i*8 + j,
				IsBlack:    isBlack,
				IsActive:   false,
				Coordinate: coordinate}

			squares = append(squares, newSquare)
		}
	}

	g := &Game{Squares: squares,
		ActiveSquares:       []*Square{},
		BlackToMove:         false,
		WhiteMayCastleShort: true,
		WhiteMayCastleLong:  true,
		BlackMayCastleShort: true,
		BlackMayCastleLong:  true,
		EnPassant:           nil,
		FiftyMoves:          0,
		Move:                1,
	}

	g.FillFen(FEN_STARTING_POS)
	return g
}

func (g *Game) ActivateSquare(num int) {
	square := g.Squares[num]
	square.IsActive = true
	g.ActiveSquares = append(g.ActiveSquares, square)
}

func (g *Game) ClearActiveSquares() {
	for _, square := range g.ActiveSquares {
		square.IsActive = false
	}
	g.ActiveSquares = []*Square{}
}

func (g *Game) ClearBoard() {
	for _, square := range g.Squares {
		square.Piece = nil
	}
}

func (g *Game) GetSquare(square string) (*Square, error) {
	square = strings.ToLower(square)
	runes := []rune(square)

	if len(runes) != 2 {
		return nil, fmt.Errorf("Invalid square")
	}

	col := int(runes[0] - 'a')
	if col < 0 || col > 7 {
		return nil, fmt.Errorf("Invalid column")
	}

	row := int(runes[1] - '1')
	if row < 0 || row > 7 {
		return nil, fmt.Errorf("Invalid row")
	}

	return g.Squares[row*8+col], nil
}
