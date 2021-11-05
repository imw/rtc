package board

type Piece interface {
	//	Moves()
	//	Stats()
	Name() string
	Ascii() string
	Unicode() string
	Side() Color
}

type piece struct {
	name string
	side Color
}

//Unicode glyphs
const (
	OutlineKing   = '\u2654'
	OutlineQueen  = '\u2655'
	OutlineRook   = '\u2656'
	OutlineBishop = '\u2657'
	OutlineKnight = '\u2658'
	OutlinePawn   = '\u2659'

	FillKing   = '\u265A'
	FillQueen  = '\u265B'
	FillRook   = '\u265C'
	FillBishop = '\u265D'
	FillKnight = '\u265E'
	FillPawn   = '\u265F'
)
