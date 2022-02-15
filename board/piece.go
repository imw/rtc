package board

//Piece is an universal interface for difference chess pieces
type Piece interface {
	Name() string
	ASCII() string
	Unicode() string
	Side() Color
	ValidMoves(Board, *Square) []Square
	Move()
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
