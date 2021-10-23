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

func (p *Knight) Ascii() rune {
	if p.core.side == White {
		return 'n'
	} else {
		return 'N'
	}
}

func (p *Knight) Unicode() rune {
	if p.core.side == White {
		return WhiteKnight
	} else {
		return BlackKnight
	}

}
