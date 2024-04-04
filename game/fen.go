package game

import (
	"fmt"
	"strconv"
	"strings"
)

const FEN_STARTING_POS = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func (g *Game) SetStartingPos() {
  g.FillFen(FEN_STARTING_POS)
}

func (g *Game) GetFen() string {
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
				res += piece.ToString()
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

func (g *Game) FillFen(fen string) error {
	if strings.Contains(fen, " ") {
		return parseFen(fen, g)
	} else {
		return parseFenPosition(fen, g)
	}
}

func parseFenRow(row string, i int, board *[]*Square) error {
	curr := 0
	split := strings.Split(row, "")

	for _, c := range split {
		gap, err := strconv.Atoi(c)
		if err != nil {
			piece, err := PieceFromString(c)
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

func parseFenPosition(fen string, d *Game) error {
	rows := strings.Split(fen, "/")

	if len(rows) != 8 {
		return fmt.Errorf("Invalid FEN format: expected 8 rows")
	}

	for i, row := range rows {
		err := parseFenRow(row, i, &d.Squares)
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

func parseEnPassant(passant string, g *Game) error {
	if passant == "-" {
		g.EnPassant = nil
		return nil
	}

	square, err := g.GetSquare(passant)
	if err != nil {
		return err
	}

	g.EnPassant = square
	return nil
}

func parseFen(fen string, g *Game) error {
	split := strings.Split(fen, " ")

	if len(split) != 6 {
		return fmt.Errorf("Invalid FEN string")
	}

	// position
	err := parseFenPosition(split[0], g)
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

	return nil
}
