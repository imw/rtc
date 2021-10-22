package board

import "0x539.lol/rtc/piece"

type Square struct {
	x        int
	y        int
	name     string
	occupant *piece.Piece
}
