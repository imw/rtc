package board

import (
	"fmt"
	"strconv"
)

const boardSize = 8

type Board struct {
	squares [boardSize][boardSize]Square
}

type Color int

const (
	White Color = iota
	Black
)

func invertColor(c Color) Color {
	if c == White {
		return Black
	} else {
		return White
	}
}

//create and color squares
func New() *Board {
	b := new(Board)
	b.squares = [boardSize][boardSize]Square{}
	for i := 0; i < boardSize; i++ {
		color := Color(i % 2)
		for j := 0; j < boardSize; j++ {
			name := string('A'+i) + strconv.Itoa(j+1)
			color = invertColor(color)
			s := Square{
				x:        i,
				y:        j,
				name:     name,
				color:    color,
				occupant: nil,
			}
			b.squares[i][j] = s
		}
	}
	b.setup()
	return b
}

func (b *Board) Flatten() [boardSize * boardSize]Square {
	sqs := [boardSize * boardSize]Square{}
	k := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			sqs[k] = b.squares[i][j]
			k++
		}
	}
	return sqs
}

//decode file rune to i, convert rank to j and invert
func (b *Board) position(name string) *Square {
	file := name[0]
	fileidx := file - 'A'
	ranktable := [8]int{7, 6, 5, 4, 3, 2, 1, 0}
	rank, _ := strconv.Atoi(string(name[1]))
	rankidx := ranktable[rank-1]
	return &b.squares[fileidx][rankidx]
}

func (b *Board) setup() {
	b.position("A1").occupant = NewRook(White)
	b.position("B1").occupant = NewKnight(White)
	b.position("C1").occupant = NewBishop(White)
	b.position("D1").occupant = NewQueen(White)
	b.position("E1").occupant = NewKing(White)
	b.position("F1").occupant = NewBishop(White)
	b.position("G1").occupant = NewKnight(White)
	b.position("H1").occupant = NewRook(White)
	files := "ABCDEFGH"
	for _, k := range files {
		b.position(fmt.Sprintf("%s%s", string(k), "2")).occupant = NewPawn(White)
	}

	for _, k := range files {
		b.position(fmt.Sprintf("%s%s", string(k), "7")).occupant = NewPawn(Black)
	}
	b.position("A8").occupant = NewRook(Black)
	b.position("B8").occupant = NewKnight(Black)
	b.position("C8").occupant = NewBishop(Black)
	b.position("D8").occupant = NewKing(Black)
	b.position("E8").occupant = NewQueen(Black)
	b.position("F8").occupant = NewBishop(Black)
	b.position("G8").occupant = NewKnight(Black)
	b.position("H8").occupant = NewRook(Black)
}
