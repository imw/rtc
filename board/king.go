package board

type King struct {
	core *piece
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

func (p *King) Ascii() rune {
	if p.core.side == White {
		return 'k'
	} else {
		return 'K'
	}
}

func (p *King) Unicode() rune {
	if p.core.side == White {
		return WhiteKing
	} else {
		return BlackKing
	}

}
