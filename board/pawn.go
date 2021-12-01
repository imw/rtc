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
	moves = append(moves, p.attackMoves(board, loc)...)
	if p.core.side == White {

		if p.moved == false {
			moves = append(moves, seekForward(board, loc, 2)...)
		}
		moves = append(moves, seekForward(board, loc, 1)...)

	} else {

		if p.moved == false {
			moves = append(moves, seekReverse(board, loc, 2)...)
		}
		moves = append(moves, seekReverse(board, loc, 1)...)

	}

	return moves

}

func (p *Pawn) attackMoves(board Board, loc Square) []Square {
	var yop binop
	if p.core.side == White {
		yop = sub
	} else {
		yop = add
	}
	moves := []Square{}
	for _, v := range search(sub, yop, board, loc, 1) {
		if v.Occupied() {
			if v.Occupant().Side() != p.core.side {
				moves = append(moves, v)
			}
		}
	}
	for _, v := range search(add, yop, board, loc, 1) {
		if v.Occupied() {
			if v.Occupant().Side() != p.core.side {
				moves = append(moves, v)
			}
		}
	}

	return moves
}

func (p *Pawn) Move() {
	p.moved = true
}
