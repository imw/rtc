package board

type Bishop struct {
	core *piece
}

func NewBishop(side Color) Piece {
	c := &piece{
		name: "Bishop",
		side: side,
	}
	return &Bishop{
		core: c,
	}
}

func (p *Bishop) Name() string {
	return p.core.name
}

func (p *Bishop) Ascii() rune {
	if p.core.side == White {
		return 'b'
	} else {
		return 'B'
	}
}

func (p *Bishop) Unicode() rune {
	if p.core.side == White {
		return WhiteBishop
	} else {
		return BlackBishop
	}

}
