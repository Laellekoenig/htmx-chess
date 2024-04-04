package game

import (
	"fmt"
	"strings"
)

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

func (piece *Piece) ToString() string {
	res := ""
	switch piece.Type {
	case Pawn:
		res = "P"
	case Rook:
		res = "R"
	case Knight:
		res = "N"
	case Bishop:
		res = "B"
	case Queen:
		res = "Q"
	case King:
		res = "K"
	}

	if piece.IsBlack {
		res = strings.ToLower(res)
	}

	return res
}

func PieceFromString(s string) (*Piece, error) {
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
