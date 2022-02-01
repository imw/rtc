package display

import (
	"0x539.lol/rtc/board"
	"0x539.lol/rtc/util"
	"github.com/gdamore/tcell/v2"
)

const (
	upperMargin = 8
	leftMargin  = 8
)

const (
	squareHeight = 1
	squareWidth  = 2
)

func Loading(s tcell.Screen) {
	w, h := s.Size()
	s.Clear()
	style := tcell.StyleDefault.Foreground(tcell.ColorCadetBlue.TrueColor()).Background(tcell.ColorWhite)
	drawText(s, w/2-8, h/2, w/2+8, h/2, style, "REAL TIME CHESS")
	drawText(s, w/2-10, h/2+1, w/2+10, h/2+1, tcell.StyleDefault, "Waiting for peer...")
	drawText(s, w/2-9, h/2+2, w/2+9, h/2+2, tcell.StyleDefault, "Press ESC to exit.")
	s.Show()
}

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

func drawSquare(s tcell.Screen, x1, y1, x2, y2 int, boardStyle tcell.Style, pieceStyle tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, boardStyle)
		}
	}

	offset := len(text) / 2
	hcenter := ((x2 - x1) / 2) + x1
	vcenter := ((y2 - y1) / 2) + y1
	drawText(s, hcenter-offset, vcenter, hcenter+offset, vcenter, pieceStyle, text)
}

//TODO: fix hardcoded board size?
func Render(b *board.Board, s tcell.Screen) {
	s.Clear()
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)
	w, h := s.Size()
	leftBound := w / leftMargin
	rightBound := leftBound + 64
	upperBound := h / upperMargin
	lowerBound := upperBound + 32
	sqWidth := (rightBound - leftBound) / b.Size()
	sqHeight := (lowerBound - upperBound) / b.Size()
	style := tcell.StyleDefault
	sqs := b.Flatten()
	loc := b.Loc()
	moves := b.Moves()
	target := b.Target()
	for _, sq := range sqs {
		piece := sq.Occupant()
		var symbol string
		if piece != nil {
			symbol = piece.Unicode()
			if piece.Side() == board.White {
				style = style.Foreground(tcell.ColorWhite)
			} else {
				style = style.Foreground(tcell.ColorLightSlateGray)
			}
		}

		if sq.Color() == board.White {
			style = style.Background(tcell.ColorDarkGray)
		} else {
			style = style.Background(tcell.ColorBlack)
		}

		if util.ItemExists(moves, *sq) {
			style = style.Background(tcell.ColorLightYellow)
		}

		if sq == loc {
			style = style.Background(tcell.ColorLightBlue)
		}

		if sq == target {
			style = style.Background(tcell.ColorLightPink)
		}

		x, y := sq.Indices()
		left := leftBound + x*sqWidth
		upper := upperBound + y*sqHeight
		right := left + sqWidth
		lower := upper + sqHeight

		drawSquare(s, left, upper, right, lower, style, style, symbol)
	}
	s.Sync()
}
