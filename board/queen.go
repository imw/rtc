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

func (p *Queen) Ascii() rune {
	if p.core.side == White {
		return 'q'
	} else {
		return 'Q'
	}
}

func (p *Queen) Unicode() rune {
	if p.core.side == White {
		return WhiteQueen
	} else {
		return BlackQueen
	}

}
