package board

import (
	"fmt"
	"strconv"

	"0x539.lol/rtc/util"
)

const boardSize = 8

//Board represents a game board
type Board struct {
	size        int
	squares     [boardSize][boardSize]Square
	whiteCursor Cursor
	blackCursor Cursor
	toMove      Color
}

//Color enumerates the opposing sides
type Color int

//possible colors are White and Black
const (
	White Color = iota
	Black
)

func invertColor(c Color) Color {
	if c == White {
		return Black
	}
	return White
}

//New returns a new, initialized game board
func New(c Color) *Board {
	b := new(Board)
	b.size = boardSize
	b.squares = [boardSize][boardSize]Square{}
	for i := 0; i < boardSize; i++ {
		color := Color(i % 2)
		for j := 0; j < boardSize; j++ {
			name := string('A'+i) + strconv.Itoa(boardSize-j)
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
	b.whiteCursor = Cursor{
		loc:   &b.squares[4][6],
		mode:  Select,
		color: White,
	}

	b.blackCursor = Cursor{
		loc:   &b.squares[4][1],
		mode:  Select,
		color: Black,
	}
	b.toMove = c
	return b
}

func (b *Board) activeCursor() *Cursor {
	if b.toMove == White {
		return &b.whiteCursor
	}
	return &b.blackCursor
}

func (b *Board) switchCursor() {
	if b.toMove == White {
		b.toMove = Black
	} else {
		b.toMove = White
	}
	b.resetCursor()
}

func (b *Board) resetCursor() {
	util.Write("resetting cursor")
	loc := b.Loc()
	util.Write(fmt.Sprintf("cursor loc: %v", loc))
	if !loc.Occupied() || loc.Occupant().Side() != b.toMove {
		util.Write("needs adjustment")
		for _, sq := range b.Flatten() {
			if sq.Occupied() {
				if sq.Occupant().Side() == b.activeCursor().color {
					b.activeCursor().loc = sq
					break
				}
			}
		}
	}
}

//Moves returns the valid next moves for the currently active cursor
func (b *Board) Moves() []Square {
	c := b.activeCursor()
	return c.choices(*b)
}

//Target returns the target of the currently active cursor
func (b *Board) Target() *Square {
	return b.activeCursor().target
}

//Loc returns the location of the currently active cursor
func (b *Board) Loc() *Square {
	return b.activeCursor().loc
}

//SetLoc moves the currently active cursor's location to s
func (b *Board) SetLoc(s *Square) {
	b.activeCursor().loc = s
}

//Side returns which side's cursor is active
func (b *Board) Side() Color {
	return b.activeCursor().color
}

//Flatten returns a sorted 1-d array of pointers to board squares
func (b *Board) Flatten() [boardSize * boardSize]*Square {
	sqs := [boardSize * boardSize]*Square{}
	k := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			sqs[k] = &b.squares[i][j]
			k++
		}
	}
	return sqs
}

//Position returns a pointer to a square given the name of that square
func (b *Board) Position(name string) *Square {
	file := name[0]
	fileidx := file - 'A'
	ranktable := [8]int{7, 6, 5, 4, 3, 2, 1, 0}
	rank, _ := strconv.Atoi(string(name[1]))
	rankidx := ranktable[rank-1]
	return &b.squares[fileidx][rankidx]
}

func (b *Board) setup() {
	b.Position("A1").occupant = NewRook(White)
	b.Position("B1").occupant = NewKnight(White)
	b.Position("C1").occupant = NewBishop(White)
	b.Position("D1").occupant = NewQueen(White)
	b.Position("E1").occupant = NewKing(White)
	b.Position("F1").occupant = NewBishop(White)
	b.Position("G1").occupant = NewKnight(White)
	b.Position("H1").occupant = NewRook(White)
	files := "ABCDEFGH"
	for _, k := range files {
		b.Position(fmt.Sprintf("%s%s", string(k), "2")).occupant = NewPawn(White)
	}

	for _, k := range files {
		b.Position(fmt.Sprintf("%s%s", string(k), "7")).occupant = NewPawn(Black)
	}
	b.Position("A8").occupant = NewRook(Black)
	b.Position("B8").occupant = NewKnight(Black)
	b.Position("C8").occupant = NewBishop(Black)
	b.Position("D8").occupant = NewKing(Black)
	b.Position("E8").occupant = NewQueen(Black)
	b.Position("F8").occupant = NewBishop(Black)
	b.Position("G8").occupant = NewKnight(Black)
	b.Position("H8").occupant = NewRook(Black)
}

//Size returns size of board
func (b *Board) Size() int {
	return b.size
}
