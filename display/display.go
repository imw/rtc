package display

import (
	"fmt"

	"0x539.lol/rtc/board"
	"github.com/gdamore/tcell/v2"
)

//Runes
const (
	WhiteKing = "\u2654"
	BlackKing = "\u265A"

	WhiteQueen = "\u2655"
	BlackQueen = "\u265B"

	WhiteRook = "\u2656"
	BlackRook = "\u265C"

	WhiteBishop = "\u2657"
	BlackBishop = "\u265D"

	WhiteKnight = "\u2658"
	BlackKnight = "\u265E"

	WhitePawn = "\u2659"
	BlackPawn = "\u265F"
)

const (
	upperMargin = 8
	leftMargin  = 8
)

const (
	squareHeight = 1
	squareWidth  = 2
)

/*
const boardBase = `
8| , . , . , . , .
7| . , . , . , . ,
6| , . , . , . , .
5| . , . , . , . ,
4| , . , . , . , .
3| . , . , . , . ,
2| , . , . , . , .
1| . , . , . , . ,
   A B C D E F G H
`

func (b *Board) render() string {
	var renderedBoard string
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b.squares[j][i].occupant != nil {
				renderedBoard = renderedBoard + " " + string(b.squares[j][i].occupant.Unicode())
			}
		}
		renderedBoard = renderedBoard + "\n"
	}
	return renderedBoard
}

func (b *Board) Print() {
	fmt.Printf("%s", b.render())
}
*/

func Greeting(s tcell.Screen) {
	w, h := s.Size()
	s.Clear()
	style := tcell.StyleDefault.Foreground(tcell.ColorCadetBlue.TrueColor()).Background(tcell.ColorWhite)
	drawText(s, w/2-8, h/2, w/2+8, h/2, style, "REAL TIME CHESS")
	drawText(s, w/2-9, h/2+1, w/2+9, h/2+1, tcell.StyleDefault, "Press ESC to exit.")
	drawText(s, w/2-12, h/2+2, w/2+12, h/2+2, tcell.StyleDefault, "Press any key to begin.")
	s.Show()
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	/*
			// Draw borders
			for col := x1; col <= x2; col++ {
				s.SetContent(col, y1, tcell.RuneHLine, nil, style)
				s.SetContent(col, y2, tcell.RuneHLine, nil, style)
			}
			for row := y1 + 1; row < y2; row++ {
				s.SetContent(x1, row, tcell.RuneVLine, nil, style)
				s.SetContent(x2, row, tcell.RuneVLine, nil, style)
			}

		// Only draw corners if necessary
		if y1 != y2 && x1 != x2 {
			s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
			s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
			s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
			s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
		}

	*/
	offset := len(text) / 2
	hcenter := ((x2 - x1) / 2) + x1
	vcenter := ((y2 - y1) / 2) + y1
	drawText(s, hcenter-offset, vcenter, hcenter+offset, vcenter, style, text)
}

/*
func (b *Board) render() string {
	var renderedBoard string
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b.squares[j][i].occupant != nil {
				renderedBoard = renderedBoard + " " + string(b.squares[j][i].occupant.Unicode())
			}
		}
		renderedBoard = renderedBoard + "\n"
	}
	return renderedBoard
}

func (b *Board) Print() {
	fmt.Printf("%s", b.render())
}
*/

//TODO: fix hardcode
func Render(b *board.Board, s tcell.Screen) {
	s.Clear()
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)
	w, h := s.Size()
	leftBound := w / leftMargin
	rightBound := leftBound + 64
	upperBound := h / upperMargin
	lowerBound := upperBound + 32
	sqWidth := (rightBound - leftBound) / 8
	sqHeight := (lowerBound - upperBound) / 8
	dims := fmt.Sprint(w, h, leftBound, upperBound, rightBound, lowerBound)
	drawBox(s, leftBound, upperBound, rightBound, lowerBound, defStyle, dims)
	style := tcell.StyleDefault
	styleWhite := tcell.StyleDefault.Background(tcell.ColorGhostWhite).Foreground(tcell.ColorBlack)
	styleBlack := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGhostWhite)
	sqs := b.Flatten()
	for _, sq := range sqs {
		x, y := sq.Indices()
		color := sq.Color()
		if color == board.White {
			style = styleWhite
		} else {
			style = styleBlack
		}
		s.SetStyle(style)
		left := leftBound + x*sqWidth
		upper := upperBound + y*sqHeight
		right := left + sqWidth
		lower := upper + sqHeight
		var symbol string
		if sq.Occupant() != nil {
			symbol = string(sq.Occupant().Unicode())
		}
		drawBox(s, left, upper, right, lower, style, symbol)
	}
	s.Sync()
}
