package board

//Queen implements the Piece interface for queens
type Queen struct {
	core *piece
}

//NewQueen returns a new queen
func NewQueen(side Color) Piece {
	c := &piece{
		name: "Queen",
		side: side,
	}
	return &Queen{
		core: c,
	}
}

//Name returns the name of this piece
func (p *Queen) Name() string {
	return p.core.name
}

//ASCII returns the Ascii representation for this piece
func (p *Queen) ASCII() string {
	if p.core.side == White {
		return string('q')
	}
	return string('Q')
}

//Unicode returns the unicode rune for this piece
func (p *Queen) Unicode() string {
	return string(FillQueen)
}

//Side return's this piece's side
func (p *Queen) Side() Color {
	return p.core.side
}

//ValidMoves returns a slice of squares representing valid moves for this
//piece, given board and location
func (p *Queen) ValidMoves(board Board, loc *Square) []Square {
	moves := []Square{}
	moves = append(moves, seekForward(board, loc, boardSize)...)
	moves = append(moves, seekReverse(board, loc, boardSize)...)
	moves = append(moves, seekLeft(board, loc, boardSize)...)
	moves = append(moves, seekRight(board, loc, boardSize)...)
	moves = append(moves, seekForwardL(board, loc, boardSize)...)
	moves = append(moves, seekForwardR(board, loc, boardSize)...)
	moves = append(moves, seekReverseL(board, loc, boardSize)...)
	moves = append(moves, seekReverseR(board, loc, boardSize)...)
	return moves
}

//Move is a noop for queens
func (p *Queen) Move() {}
