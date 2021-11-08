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

//TODO: Implement
func (p *King) ValidMoves(board Board, loc Square) []Square {
	moves := []Square{}
	return moves
}

func (p *King) Move() {
	p.moved = true
}
