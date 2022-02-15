package board

//Rook implements the Piece interface for rooks
type Rook struct {
	core  *piece
	moved bool
}

//NewRook returns a new Rook
func NewRook(side Color) Piece {
	c := &piece{
		name: "Rook",
		side: side,
	}
	return &Rook{
		core: c,
	}
}

//Name returns the name of this piece
func (p *Rook) Name() string {
	return p.core.name
}

//ASCII returns the Ascii representation for this piece
func (p *Rook) ASCII() string {
	if p.core.side == White {
		return string('r')
	}
	return string('R')
}

//Unicode returns the unicode rune for this piece
func (p *Rook) Unicode() string {
	return string(FillRook)
}

//Side return's this piece's side
func (p *Rook) Side() Color {
	return p.core.side
}

//ValidMoves returns a slice of squares representing valid moves for this
//piece, given board and location
//TODO: Castling
func (p *Rook) ValidMoves(board Board, loc *Square) []Square {
	moves := []Square{}
	moves = append(moves, seekForward(board, loc, boardSize)...)
	moves = append(moves, seekReverse(board, loc, boardSize)...)
	moves = append(moves, seekLeft(board, loc, boardSize)...)
	moves = append(moves, seekRight(board, loc, boardSize)...)
	return moves
}

//Move returns whether this piece has moved at least once
func (p *Rook) Move() {
	p.moved = true
}
