package board

type Pawn struct {
	core  *piece
	moved bool
}

func NewPawn(side Color) Piece {
	c := &piece{
		name: "Pawn",
		side: side,
	}
	return &Pawn{
		core: c,
	}
}

func (p *Pawn) Name() string {
	return p.core.name
}

func (p *Pawn) Ascii() string {
	if p.core.side == White {
		return string('P')
	} else {
		return string('p')
	}
}

func (p *Pawn) Unicode() string {
	return string(FillPawn)
}

func (p *Pawn) Side() Color {
	return p.core.side
}

func (p *Pawn) ValidMoves(board Board, loc Square) []Square {
	moves := []Square{}
	moves = append(moves, loc)
	if p.core.side == White {

		if p.moved == false {
			moves = append(moves, seekForward(board, loc, 2)...)
		} else {
			moves = append(moves, seekForward(board, loc, 1)...)
		}

		moves = append(moves, seekForwardL(board, loc, 1)...)
		moves = append(moves, seekForwardR(board, loc, 1)...)

	} else {

		if p.moved == false {
			moves = append(moves, seekReverse(board, loc, 2)...)
		} else {
			moves = append(moves, seekReverse(board, loc, 1)...)
		}

		moves = append(moves, seekReverseL(board, loc, 1)...)
		moves = append(moves, seekReverseR(board, loc, 1)...)
	}

	return moves

}

func (p *Pawn) Move(from, to Square) {
	p.moved = true
	to.occupant = p
	from.occupant = nil
}
