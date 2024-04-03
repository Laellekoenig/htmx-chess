package game

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
	Fen                 string
}

func CreateGame() *Game {
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
			squares = append(squares, &Square{Num: i*8 + j, IsBlack: isBlack, IsActive: false, Coordinate: coordinate})
		}
	}

	squares[0].Piece = &Piece{Type: Knight, IsBlack: true}

  g := &Game{Squares: squares,
		ActiveSquares:       []*Square{},
		BlackToMove:         false,
		WhiteMayCastleShort: true,
		WhiteMayCastleLong:  true,
		BlackMayCastleLong:  true,
		BlackMayCastleShort: true,
		EnPassant:           nil,
		FiftyMoves:          0,
		Move:                1,
		Fen:                 ""}

  ParseFEN(FEN_START, g)
  return g
}

func SetActive(data *Game, num int) {
	square := data.Squares[num]
	square.IsActive = true
	data.ActiveSquares = append(data.ActiveSquares, square)
}

func ClearAllActiveSquares(data *Game) {
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

func ClearBoard(d *Game) {
	for _, square := range d.Squares {
		square.Piece = nil
	}
}

func parseFENRow(row string, i int, board *[]*Square) error {
	curr := 0
	split := strings.Split(row, "")

	for _, c := range split {
		gap, err := strconv.Atoi(c)
		if err != nil {
			piece, err := getPiece(c)
			if err != nil {
				return err
			} else {
				(*board)[i*8+curr].Piece = piece
				curr += 1
			}
		} else {
			curr += gap
		}
	}

	if curr != 8 {
		return fmt.Errorf("Invalid FEN position")
	}

	return nil
}

func ParseFENPosition(fen string, d *Game) error {
	rows := strings.Split(fen, "/")

	if len(rows) != 8 {
		return fmt.Errorf("Invalid FEN format: expected 8 rows")
	}

	for i, row := range rows {
		err := parseFENRow(row, i, &d.Squares)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseCasteling(castle string, g *Game) error {
	wShort := strings.Contains(castle, "K")
	bShort := strings.Contains(castle, "k")
	wLong := strings.Contains(castle, "Q")
	bLong := strings.Contains(castle, "q")

	if castle == "-" {
		g.WhiteMayCastleLong = false
		g.WhiteMayCastleShort = false
		g.BlackMayCastleLong = false
		g.BlackMayCastleShort = false
		return nil
	}

	if !wShort && !wLong && !bShort && !bLong {
		return fmt.Errorf("Invalid casteling string")
	}

	if wShort {
		g.WhiteMayCastleShort = true
	}

	if bShort {
		g.BlackMayCastleShort = true
	}

	if wLong {
		g.WhiteMayCastleLong = true
	}

	if bLong {
		g.BlackMayCastleLong = true
	}

	return nil
}

func getSquare(square string, g *Game) (*Square, error) {
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

func parseEnPassant(passant string, g *Game) error {
  if passant == "-" {
    g.EnPassant = nil
    return nil
  }

	square, err := getSquare(passant, g)
	if err != nil {
		return err
	}

	g.EnPassant = square
	return nil
}

func ParseFEN(fen string, g *Game) error {
	split := strings.Split(fen, " ")

	if len(split) != 6 {
		return fmt.Errorf("Invalid FEN string")
	}

	// position
	err := ParseFENPosition(split[0], g)
	if err != nil {
		return err
	}

	// player to move
	player := split[1]
	if player == "w" {
		g.BlackToMove = false
	} else if player == "b" {
		g.BlackToMove = true
	} else {
		return fmt.Errorf("Invalid next player, expected w or b")
	}

	// casteling
	err = parseCasteling(split[2], g)
	if err != nil {
		return err
	}

	// en passant
	err = parseEnPassant(split[3], g)
	if err != nil {
		return err
	}

	// 50-rule
	fiftyMoves, err := strconv.Atoi(split[4])
	if err != nil || fiftyMoves < 0 {
		return fmt.Errorf("Invalid number of 50-rule moves")
	}
	g.FiftyMoves = fiftyMoves

	// move number
	move, err := strconv.Atoi(split[5])
	if err != nil || move < 0 {
		return fmt.Errorf("Invalid move number")
	}
	g.Move = move

  g.Fen = ToFEN(g)

	return nil
}

func pieceToString(piece *Piece) string {
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

func ToFEN(g *Game) string {
	res := ""

	board := g.Squares
	for i := 0; i < 8; i++ {
		curr := 0

		for j := 0; j < 8; j++ {
			piece := board[i*8+j].Piece

			if piece == nil {
				curr += 1
			} else {
        if curr != 0 {
          res += fmt.Sprintf("%d", curr)
          curr = 0
        }
				res += pieceToString(piece)
			}
		}

    if curr != 0 {
      res += fmt.Sprintf("%d", curr)
    }

    if i != 7 {
      res += "/"
    }
	}

	return res
}
