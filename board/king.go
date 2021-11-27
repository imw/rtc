package board

type King struct {
	core  *piece
	moved bool
}

func NewKing(side Color) Piece {
	c := &piece{
		name: "King",
		side: side,
	}
	return &King{
		core: c,
	}
}

func (p *King) Name() string {
	return p.core.name
}

func (p *King) Ascii() string {
	if p.core.side == White {
		return string('k')
	} else {
		return string('K')
	}
}

func (p *King) Unicode() string {
	return string(FillKing)
}

func (p *King) Side() Color {
	return p.core.side
}

//TODO: Castling
func (p *King) ValidMoves(board Board, loc Square) []Square {
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

func (p *King) Move(from, to Square) {
	p.moved = true
	from.occupant = nil
	to.occupant = p
}
