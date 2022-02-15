package board

//King implements the Piece interface for kings
type King struct {
	core  *piece
	moved bool
}

//NewKing returns a new King
func NewKing(side Color) Piece {
	c := &piece{
		name: "King",
		side: side,
	}
	return &King{
		core: c,
	}
}

//Name returns the name of this piece
func (p *King) Name() string {
	return p.core.name
}

//ASCII returns the Ascii representation for this piece
func (p *King) ASCII() string {
	if p.core.side == White {
		return string('k')
	}
	return string('K')
}

//Unicode returns the unicode rune for this piece
func (p *King) Unicode() string {
	return string(FillKing)
}

//Side return's this piece's side
func (p *King) Side() Color {
	return p.core.side
}

//ValidMoves returns a slice of squares representing valid moves for this
//piece, given board and location
//TODO: Castling
func (p *King) ValidMoves(board Board, loc *Square) []Square {
	moves := []Square{}
	moves = append(moves, seekForward(board, loc, 1)...)
	moves = append(moves, seekReverse(board, loc, 1)...)
	moves = append(moves, seekLeft(board, loc, 1)...)
	moves = append(moves, seekRight(board, loc, 1)...)
	moves = append(moves, seekForwardL(board, loc, 1)...)
	moves = append(moves, seekForwardR(board, loc, 1)...)
	moves = append(moves, seekReverseL(board, loc, 1)...)
	moves = append(moves, seekReverseR(board, loc, 1)...)
	return moves
}

//Move marks piece as having moved at least once
func (p *King) Move() {
	p.moved = true
}
