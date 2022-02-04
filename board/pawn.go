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

func (p *Pawn) ValidMoves(board Board, loc *Square) []Square {
	moves := []Square{}
	moves = append(moves, *loc)
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

	for i, v := range moves {
		if p.opponentOccupied(v) {
			moves = removeSquare(moves, i)
		}
	}

	moves = append(moves, p.attackMoves(board, loc)...)
	return moves

}

func (p *Pawn) attackMoves(board Board, loc *Square) []Square {
	var yop binop
	if p.core.side == White {
		yop = sub
	} else {
		yop = add
	}
	moves := []Square{}
	for _, v := range search(sub, yop, board, loc, 1) {
		if p.opponentOccupied(v) {
			moves = append(moves, v)
		}
	}
	for _, v := range search(add, yop, board, loc, 1) {
		if p.opponentOccupied(v) {
			moves = append(moves, v)
		}
	}

	return moves
}

func (p *Pawn) opponentOccupied(s Square) bool {
	if s.Occupied() {
		if s.Occupant().Side() != p.core.side {
			return true
		}
	}
	return false
}

func removeSquare(sqs []Square, index int) []Square {
	return append(sqs[:index], sqs[index+1:]...)
}

func (p *Pawn) Move() {
	p.moved = true
}
