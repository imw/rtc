package board

type Knight struct {
	core *piece
}

func NewKnight(side Color) Piece {
	c := &piece{
		name: "Knight",
		side: side,
	}
	return &Knight{
		core: c,
	}
}

func (p *Knight) Name() string {
	return p.core.name
}

func (p *Knight) Ascii() string {
	if p.core.side == White {
		return string('n')
	} else {
		return string('N')
	}
}

func (p *Knight) Unicode() string {
	return string(FillKnight)
}

func (p *Knight) Side() Color {
	return p.core.side
}

//TODO: Implement
func (p *Knight) GetMoves(board Board, loc Square) []Square {
	moves := []Square{}
	return moves
}

//NOOP
func (p *Knight) Move() {}
