package board

type Queen struct {
	core *piece
}

func NewQueen(side Color) Piece {
	c := &piece{
		name: "Queen",
		side: side,
	}
	return &Queen{
		core: c,
	}
}

func (p *Queen) Name() string {
	return p.core.name
}

func (p *Queen) Ascii() string {
	if p.core.side == White {
		return string('q')
	} else {
		return string('Q')
	}
}

func (p *Queen) Unicode() string {
	return string(FillQueen)
}

func (p *Queen) Side() Color {
	return p.core.side
}

func (p *Queen) ValidMoves(board Board, loc Square) []Square {
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

//NOOP
func (p *Queen) Move() {}
