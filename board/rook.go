package board

type Rook struct {
	core *piece
}

func NewRook(side Color) Piece {
	c := &piece{
		name: "Rook",
		side: side,
	}
	return &Rook{
		core: c,
	}
}

func (p *Rook) Name() string {
	return p.core.name
}

func (p *Rook) Ascii() rune {
	if p.core.side == White {
		return 'r'
	} else {
		return 'R'
	}
}

func (p *Rook) Unicode() rune {
	if p.core.side == White {
		return WhiteRook
	} else {
		return BlackRook
	}

}
