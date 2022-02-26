package board

//Bishop implements the Piece interface for bishops
type Bishop struct {
	core *piece
}

//NewBishop returns a new Bishop
func NewBishop(side Color) Piece {
	c := &piece{
		name: "Bishop",
		side: side,
	}
	return &Bishop{
		core: c,
	}
}

//Name returns the name of this piece
func (p *Bishop) Name() string {
	return p.core.name
}

//ASCII returns the Ascii representation for this piece
func (p *Bishop) ASCII() string {
	if p.core.side == White {
		return string('b')
	}
	return string('B')
}

//Unicode returns the unicode rune for this piece
func (p *Bishop) Unicode() string {
	return string(FillBishop)
}

//Side return's this piece's side
func (p *Bishop) Side() Color {
	return p.core.side
}

//ValidMoves returns a slice of squares representing valid moves for this
//piece, given board and location
func (p *Bishop) ValidMoves(board Board, loc *Square) []Square {
	moves := []Square{}
	moves = append(moves, seekForwardL(board, loc, boardSize)...)
	moves = append(moves, seekForwardR(board, loc, boardSize)...)
	moves = append(moves, seekReverseL(board, loc, boardSize)...)
	moves = append(moves, seekReverseR(board, loc, boardSize)...)
	return moves
}

//Move is a noop for Bishops
func (p *Bishop) Move() {}
