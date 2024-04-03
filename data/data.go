package data

import (
	"fmt"
	"strconv"
	"strings"
)

const FEN_START = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

type PieceType = int

const (
	Pawn PieceType = iota
	Rook
	Knight
	Bishop
	Queen
	King
)

type Piece struct {
	Type    PieceType
	IsBlack bool
}

type Square struct {
	Num      int
	IsBlack  bool
	IsActive bool
	Piece    *Piece
}

type Data struct {
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

func CreateData() *Data {
	var squares []*Square

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			isBlack := true

			if i%2 == 0 && j%2 == 0 {
				isBlack = false
			} else if i%2 != 0 && j%2 != 0 {
				isBlack = false
			}

			squares = append(squares, &Square{Num: i*8 + j, IsBlack: isBlack, IsActive: false})
		}
	}

	squares[0].Piece = &Piece{Type: Knight, IsBlack: true}

	return &Data{Squares: squares, ActiveSquares: []*Square{}}
}

func SetActive(data *Data, num int) {
	square := data.Squares[num]
	square.IsActive = true
	data.ActiveSquares = append(data.ActiveSquares, square)
}

func ClearAllActiveSquares(data *Data) {
	for _, square := range (*data).ActiveSquares {
		(*square).IsActive = false
	}
	data.ActiveSquares = []*Square{}
}

func getPiece(s string) (*Piece, error) {
	switch s {
	case "k":
		return &Piece{Type: King, IsBlack: true}, nil
	case "K":
		return &Piece{Type: King, IsBlack: false}, nil
	case "q":
		return &Piece{Type: Queen, IsBlack: true}, nil
	case "Q":
		return &Piece{Type: Queen, IsBlack: false}, nil
	case "r":
		return &Piece{Type: Rook, IsBlack: true}, nil
	case "R":
		return &Piece{Type: Rook, IsBlack: false}, nil
	case "n":
		return &Piece{Type: Knight, IsBlack: true}, nil
	case "N":
		return &Piece{Type: Knight, IsBlack: false}, nil
	case "b":
		return &Piece{Type: Bishop, IsBlack: true}, nil
	case "B":
		return &Piece{Type: Bishop, IsBlack: false}, nil
	case "p":
		return &Piece{Type: Pawn, IsBlack: true}, nil
	case "P":
		return &Piece{Type: Pawn, IsBlack: false}, nil
	}

	return nil, fmt.Errorf("Invalid piece")
}

func ClearBoard(d *Data) {
	for _, square := range d.Squares {
		square.Piece = nil
	}
}

func setRow(row string, i int, board *[]*Square) {
	curr := 0
	split := strings.Split(row, "")
	for _, c := range split {
		gap, err := strconv.Atoi(c)
		if err != nil {
			piece, err := getPiece(c)
			if err != nil {
				fmt.Printf("%v\n", err)
			} else {
				(*board)[i*8+curr].Piece = piece
				curr++
			}
		} else {
			curr += gap
		}
	}
}

func ParseFENPosition(fen string, d *Data) {
	rows := strings.Split(fen, "/")
	for i, row := range rows {
		setRow(row, i, &d.Squares)
	}
}

func ParseFEN(fen string, d *Data) {
	split := strings.Split(fen, " ")
	ParseFENPosition(split[0], d)
}
